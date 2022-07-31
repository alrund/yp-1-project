// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cooker is an autogenerated mock type for the Cooker type
type Cooker struct {
	mock.Mock
}

// AddCookie provides a mock function with given fields: name, value, expireTime, w
func (_m *Cooker) AddCookie(name string, value string, expireTime time.Time, w http.ResponseWriter) {
	_m.Called(name, value, expireTime, w)
}

// AddCookieWithDuration provides a mock function with given fields: name, value, duration, w
func (_m *Cooker) AddCookieWithDuration(name string, value string, duration string, w http.ResponseWriter) error {
	ret := _m.Called(name, value, duration, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, http.ResponseWriter) error); ok {
		r0 = rf(name, value, duration, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearCookie provides a mock function with given fields: name, w
func (_m *Cooker) ClearCookie(name string, w http.ResponseWriter) {
	_m.Called(name, w)
}

type mockConstructorTestingTNewCooker interface {
	mock.TestingT
	Cleanup(func())
}

// NewCooker creates a new instance of Cooker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCooker(t mockConstructorTestingTNewCooker) *Cooker {
	mock := &Cooker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
