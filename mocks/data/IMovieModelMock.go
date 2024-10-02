// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	data "greenlight.webjenga.com/internal/data"
)

// IMovieModelMock is an autogenerated mock type for the IMovieModel type
type IMovieModelMock struct {
	mock.Mock
}

type IMovieModelMock_Expecter struct {
	mock *mock.Mock
}

func (_m *IMovieModelMock) EXPECT() *IMovieModelMock_Expecter {
	return &IMovieModelMock_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: id
func (_m *IMovieModelMock) Delete(id int64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IMovieModelMock_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IMovieModelMock_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int64
func (_e *IMovieModelMock_Expecter) Delete(id interface{}) *IMovieModelMock_Delete_Call {
	return &IMovieModelMock_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *IMovieModelMock_Delete_Call) Run(run func(id int64)) *IMovieModelMock_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *IMovieModelMock_Delete_Call) Return(_a0 error) *IMovieModelMock_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMovieModelMock_Delete_Call) RunAndReturn(run func(int64) error) *IMovieModelMock_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: title, genres, filters
func (_m *IMovieModelMock) FindAll(title string, genres []string, filters data.Filters) ([]*data.Movie, data.Metadata, error) {
	ret := _m.Called(title, genres, filters)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*data.Movie
	var r1 data.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(string, []string, data.Filters) ([]*data.Movie, data.Metadata, error)); ok {
		return rf(title, genres, filters)
	}
	if rf, ok := ret.Get(0).(func(string, []string, data.Filters) []*data.Movie); ok {
		r0 = rf(title, genres, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.Movie)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, data.Filters) data.Metadata); ok {
		r1 = rf(title, genres, filters)
	} else {
		r1 = ret.Get(1).(data.Metadata)
	}

	if rf, ok := ret.Get(2).(func(string, []string, data.Filters) error); ok {
		r2 = rf(title, genres, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// IMovieModelMock_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type IMovieModelMock_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - title string
//   - genres []string
//   - filters data.Filters
func (_e *IMovieModelMock_Expecter) FindAll(title interface{}, genres interface{}, filters interface{}) *IMovieModelMock_FindAll_Call {
	return &IMovieModelMock_FindAll_Call{Call: _e.mock.On("FindAll", title, genres, filters)}
}

func (_c *IMovieModelMock_FindAll_Call) Run(run func(title string, genres []string, filters data.Filters)) *IMovieModelMock_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(data.Filters))
	})
	return _c
}

func (_c *IMovieModelMock_FindAll_Call) Return(_a0 []*data.Movie, _a1 data.Metadata, _a2 error) *IMovieModelMock_FindAll_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *IMovieModelMock_FindAll_Call) RunAndReturn(run func(string, []string, data.Filters) ([]*data.Movie, data.Metadata, error)) *IMovieModelMock_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: id
func (_m *IMovieModelMock) FindOne(id int64) (*data.Movie, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 *data.Movie
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*data.Movie, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *data.Movie); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Movie)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IMovieModelMock_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type IMovieModelMock_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - id int64
func (_e *IMovieModelMock_Expecter) FindOne(id interface{}) *IMovieModelMock_FindOne_Call {
	return &IMovieModelMock_FindOne_Call{Call: _e.mock.On("FindOne", id)}
}

func (_c *IMovieModelMock_FindOne_Call) Run(run func(id int64)) *IMovieModelMock_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *IMovieModelMock_FindOne_Call) Return(_a0 *data.Movie, _a1 error) *IMovieModelMock_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IMovieModelMock_FindOne_Call) RunAndReturn(run func(int64) (*data.Movie, error)) *IMovieModelMock_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: movie
func (_m *IMovieModelMock) Insert(movie *data.Movie) error {
	ret := _m.Called(movie)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Movie) error); ok {
		r0 = rf(movie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IMovieModelMock_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type IMovieModelMock_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - movie *data.Movie
func (_e *IMovieModelMock_Expecter) Insert(movie interface{}) *IMovieModelMock_Insert_Call {
	return &IMovieModelMock_Insert_Call{Call: _e.mock.On("Insert", movie)}
}

func (_c *IMovieModelMock_Insert_Call) Run(run func(movie *data.Movie)) *IMovieModelMock_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*data.Movie))
	})
	return _c
}

func (_c *IMovieModelMock_Insert_Call) Return(_a0 error) *IMovieModelMock_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMovieModelMock_Insert_Call) RunAndReturn(run func(*data.Movie) error) *IMovieModelMock_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: movie
func (_m *IMovieModelMock) Update(movie *data.Movie) error {
	ret := _m.Called(movie)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Movie) error); ok {
		r0 = rf(movie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IMovieModelMock_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type IMovieModelMock_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - movie *data.Movie
func (_e *IMovieModelMock_Expecter) Update(movie interface{}) *IMovieModelMock_Update_Call {
	return &IMovieModelMock_Update_Call{Call: _e.mock.On("Update", movie)}
}

func (_c *IMovieModelMock_Update_Call) Run(run func(movie *data.Movie)) *IMovieModelMock_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*data.Movie))
	})
	return _c
}

func (_c *IMovieModelMock_Update_Call) Return(_a0 error) *IMovieModelMock_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMovieModelMock_Update_Call) RunAndReturn(run func(*data.Movie) error) *IMovieModelMock_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMovieModelMock creates a new instance of IMovieModelMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMovieModelMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMovieModelMock {
	mock := &IMovieModelMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}