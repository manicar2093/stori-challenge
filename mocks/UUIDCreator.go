// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UUIDCreator is an autogenerated mock type for the UUIDCreator type
type UUIDCreator struct {
	mock.Mock
}

type UUIDCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *UUIDCreator) EXPECT() *UUIDCreator_Expecter {
	return &UUIDCreator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *UUIDCreator) Execute() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// UUIDCreator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type UUIDCreator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *UUIDCreator_Expecter) Execute() *UUIDCreator_Execute_Call {
	return &UUIDCreator_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *UUIDCreator_Execute_Call) Run(run func()) *UUIDCreator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UUIDCreator_Execute_Call) Return(_a0 uuid.UUID) *UUIDCreator_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UUIDCreator_Execute_Call) RunAndReturn(run func() uuid.UUID) *UUIDCreator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewUUIDCreator creates a new instance of UUIDCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUUIDCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *UUIDCreator {
	mock := &UUIDCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
