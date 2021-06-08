// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UpdaterService is an autogenerated mock type for the UpdaterService type
type UpdaterService struct {
	mock.Mock
}

// EnsureHistoricalExchangeRates provides a mock function with given fields: ctx
func (_m *UpdaterService) EnsureHistoricalExchangeRates(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureRecentExchangeRates provides a mock function with given fields: ctx
func (_m *UpdaterService) EnsureRecentExchangeRates(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
