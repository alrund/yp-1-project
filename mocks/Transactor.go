// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Transactor is an autogenerated mock type for the Transactor type
type Transactor struct {
	mock.Mock
}

// ExecContext provides a mock function with given fields: ctx, db, query, args
func (_m *Transactor) ExecContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
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

// QueryContext provides a mock function with given fields: ctx, db, query, args
func (_m *Transactor) QueryContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
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

// QueryRowContext provides a mock function with given fields: ctx, db, query, args
func (_m *Transactor) QueryRowContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, db, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, *sql.DB, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, db, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// WithinTransaction provides a mock function with given fields: ctx, tFunc
func (_m *Transactor) WithinTransaction(ctx context.Context, tFunc func(context.Context) error) error {
	ret := _m.Called(ctx, tFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, tFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactor interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactor creates a new instance of Transactor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactor(t mockConstructorTestingTNewTransactor) *Transactor {
	mock := &Transactor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
