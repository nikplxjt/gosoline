// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx
func (_m *Reader) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Reader) Stop() {
	_m.Called()
}