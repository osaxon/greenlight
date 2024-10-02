package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"sync"
	"time"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a deferred function (always runs in the event of panic
		// as Go unwinds the stack).
		//
		defer func() {
			// use builtin recover function to check if there has been a panic or not
			if err := recover(); err != nil {
				// if there was a panic, set connection header to close
				// on the response. This makes Go's HTTP server close the connection
				w.Header().Set("Connection", "close")

				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// background goroutine to remove old entries from the client map once every minute
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if app.config.limiter.enabled {
			// extract the client's IP addr from the req
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			// lock the mutex to prevent code running concurrently
			mu.Lock()

			// check if the IP already exists in the map
			// if not, initialise a new rate limiter and add the IP and limiter to the map
			if _, found := clients[ip]; !found {
				clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst)}
			}
			clients[ip].lastSeen = time.Now()

			// call the Allow() method on the limiter for the current IP
			// if the request isn't allowed, unlock the mutex and return an error response
			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceededResponse(w, r)
				return
			}

			// unlock the mutex before calling the next handler in the chain
			// IMPORTANT: don't defer otherwise all handlers in chain would need to complete
			// before the mutex is unlocked
			mu.Unlock()
		}

		next.ServeHTTP(w, r)
	})
}
