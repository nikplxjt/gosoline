// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	ladon "github.com/ory/ladon"
	mock "github.com/stretchr/testify/mock"
)

// Guard is an autogenerated mock type for the Guard type
type Guard struct {
	mock.Mock
}

// CreatePolicy provides a mock function with given fields: pol
func (_m *Guard) CreatePolicy(pol ladon.Policy) error {
	ret := _m.Called(pol)

	var r0 error
	if rf, ok := ret.Get(0).(func(ladon.Policy) error); ok {
		r0 = rf(pol)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePolicy provides a mock function with given fields: pol
func (_m *Guard) DeletePolicy(pol ladon.Policy) error {
	ret := _m.Called(pol)

	var r0 error
	if rf, ok := ret.Get(0).(func(ladon.Policy) error); ok {
		r0 = rf(pol)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPolicies provides a mock function with given fields:
func (_m *Guard) GetPolicies() (ladon.Policies, error) {
	ret := _m.Called()

	var r0 ladon.Policies
	if rf, ok := ret.Get(0).(func() ladon.Policies); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ladon.Policies)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoliciesBySubject provides a mock function with given fields: subject
func (_m *Guard) GetPoliciesBySubject(subject string) (ladon.Policies, error) {
	ret := _m.Called(subject)

	var r0 ladon.Policies
	if rf, ok := ret.Get(0).(func(string) ladon.Policies); ok {
		r0 = rf(subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ladon.Policies)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAllowed provides a mock function with given fields: request
func (_m *Guard) IsAllowed(request *ladon.Request) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ladon.Request) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePolicy provides a mock function with given fields: pol
func (_m *Guard) UpdatePolicy(pol ladon.Policy) error {
	ret := _m.Called(pol)

	var r0 error
	if rf, ok := ret.Get(0).(func(ladon.Policy) error); ok {
		r0 = rf(pol)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
