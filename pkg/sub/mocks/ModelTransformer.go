// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import sub "github.com/applike/gosoline/pkg/sub"

// ModelTransformer is an autogenerated mock type for the ModelTransformer type
type ModelTransformer struct {
	mock.Mock
}

// GetInput provides a mock function with given fields:
func (_m *ModelTransformer) GetInput() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Transform provides a mock function with given fields: ctx, inp
func (_m *ModelTransformer) Transform(ctx context.Context, inp interface{}) (sub.Model, error) {
	ret := _m.Called(ctx, inp)

	var r0 sub.Model
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) sub.Model); ok {
		r0 = rf(ctx, inp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sub.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, inp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}