// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	port "github.com/alrund/yp-1-project/internal/domain/port"
	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

type Router_Expecter struct {
	mock *mock.Mock
}

func (_m *Router) EXPECT() *Router_Expecter {
	return &Router_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: path, handler
func (_m *Router) Get(path string, handler http.Handler) {
	_m.Called(path, handler)
}

// Router_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Router_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - path string
//  - handler http.Handler
func (_e *Router_Expecter) Get(path interface{}, handler interface{}) *Router_Get_Call {
	return &Router_Get_Call{Call: _e.mock.On("Get", path, handler)}
}

func (_c *Router_Get_Call) Run(run func(path string, handler http.Handler)) *Router_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.Handler))
	})
	return _c
}

func (_c *Router_Get_Call) Return() *Router_Get_Call {
	_c.Call.Return()
	return _c
}

// Post provides a mock function with given fields: path, handler
func (_m *Router) Post(path string, handler http.Handler) {
	_m.Called(path, handler)
}

// Router_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type Router_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//  - path string
//  - handler http.Handler
func (_e *Router_Expecter) Post(path interface{}, handler interface{}) *Router_Post_Call {
	return &Router_Post_Call{Call: _e.mock.On("Post", path, handler)}
}

func (_c *Router_Post_Call) Run(run func(path string, handler http.Handler)) *Router_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.Handler))
	})
	return _c
}

func (_c *Router_Post_Call) Return() *Router_Post_Call {
	_c.Call.Return()
	return _c
}

// ServeHTTP provides a mock function with given fields: wrt, req
func (_m *Router) ServeHTTP(wrt http.ResponseWriter, req *http.Request) {
	_m.Called(wrt, req)
}

// Router_ServeHTTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServeHTTP'
type Router_ServeHTTP_Call struct {
	*mock.Call
}

// ServeHTTP is a helper method to define mock.On call
//  - wrt http.ResponseWriter
//  - req *http.Request
func (_e *Router_Expecter) ServeHTTP(wrt interface{}, req interface{}) *Router_ServeHTTP_Call {
	return &Router_ServeHTTP_Call{Call: _e.mock.On("ServeHTTP", wrt, req)}
}

func (_c *Router_ServeHTTP_Call) Run(run func(wrt http.ResponseWriter, req *http.Request)) *Router_ServeHTTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Router_ServeHTTP_Call) Return() *Router_ServeHTTP_Call {
	_c.Call.Return()
	return _c
}

// Use provides a mock function with given fields: mwf
func (_m *Router) Use(mwf ...func(http.Handler) http.Handler) {
	_va := make([]interface{}, len(mwf))
	for _i := range mwf {
		_va[_i] = mwf[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Router_Use_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Use'
type Router_Use_Call struct {
	*mock.Call
}

// Use is a helper method to define mock.On call
//  - mwf ...func(http.Handler) http.Handler
func (_e *Router_Expecter) Use(mwf ...interface{}) *Router_Use_Call {
	return &Router_Use_Call{Call: _e.mock.On("Use",
		append([]interface{}{}, mwf...)...)}
}

func (_c *Router_Use_Call) Run(run func(mwf ...func(http.Handler) http.Handler)) *Router_Use_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(http.Handler) http.Handler, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(func(http.Handler) http.Handler)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Router_Use_Call) Return() *Router_Use_Call {
	_c.Call.Return()
	return _c
}

// WithPrefix provides a mock function with given fields: prefix
func (_m *Router) WithPrefix(prefix string) port.Router {
	ret := _m.Called(prefix)

	var r0 port.Router
	if rf, ok := ret.Get(0).(func(string) port.Router); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(port.Router)
		}
	}

	return r0
}

// Router_WithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithPrefix'
type Router_WithPrefix_Call struct {
	*mock.Call
}

// WithPrefix is a helper method to define mock.On call
//  - prefix string
func (_e *Router_Expecter) WithPrefix(prefix interface{}) *Router_WithPrefix_Call {
	return &Router_WithPrefix_Call{Call: _e.mock.On("WithPrefix", prefix)}
}

func (_c *Router_WithPrefix_Call) Run(run func(prefix string)) *Router_WithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Router_WithPrefix_Call) Return(_a0 port.Router) *Router_WithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewRouter interface {
	mock.TestingT
	Cleanup(func())
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRouter(t mockConstructorTestingTNewRouter) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
