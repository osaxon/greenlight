package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"greenlight.webjenga.com/internal/assert"
)

func TestHealthcheck(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	app := &application{}

	app.healthcheck(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	var response map[string]map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, response["data"]["status"], "UP")
}
