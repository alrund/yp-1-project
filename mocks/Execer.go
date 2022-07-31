// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Execer is an autogenerated mock type for the Execer type
type Execer struct {
	mock.Mock
}

// ExecContext provides a mock function with given fields: ctx, db, query, args
func (_m *Execer) ExecContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sql.DB, string, ...interface{}) error); ok {
		r1 = rf(ctx, db, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExecer interface {
	mock.TestingT
	Cleanup(func())
}

// NewExecer creates a new instance of Execer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExecer(t mockConstructorTestingTNewExecer) *Execer {
	mock := &Execer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
