// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	txanalizer "github.com/manicar2093/stori-challenge/internal/txanalizer"
	mock "github.com/stretchr/testify/mock"
)

// Notificator is an autogenerated mock type for the Notificator type
type Notificator struct {
	mock.Mock
}

type Notificator_Expecter struct {
	mock *mock.Mock
}

func (_m *Notificator) EXPECT() *Notificator_Expecter {
	return &Notificator_Expecter{mock: &_m.Mock}
}

// SendAccountDetailsEmail provides a mock function with given fields: input
func (_m *Notificator) SendAccountDetailsEmail(input txanalizer.SendAccountDetailsEmailInput) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(txanalizer.SendAccountDetailsEmailInput) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notificator_SendAccountDetailsEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendAccountDetailsEmail'
type Notificator_SendAccountDetailsEmail_Call struct {
	*mock.Call
}

// SendAccountDetailsEmail is a helper method to define mock.On call
//   - input txanalizer.SendAccountDetailsEmailInput
func (_e *Notificator_Expecter) SendAccountDetailsEmail(input interface{}) *Notificator_SendAccountDetailsEmail_Call {
	return &Notificator_SendAccountDetailsEmail_Call{Call: _e.mock.On("SendAccountDetailsEmail", input)}
}

func (_c *Notificator_SendAccountDetailsEmail_Call) Run(run func(input txanalizer.SendAccountDetailsEmailInput)) *Notificator_SendAccountDetailsEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(txanalizer.SendAccountDetailsEmailInput))
	})
	return _c
}

func (_c *Notificator_SendAccountDetailsEmail_Call) Return(_a0 error) *Notificator_SendAccountDetailsEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Notificator_SendAccountDetailsEmail_Call) RunAndReturn(run func(txanalizer.SendAccountDetailsEmailInput) error) *Notificator_SendAccountDetailsEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotificator creates a new instance of Notificator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Notificator {
	mock := &Notificator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
