// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	filestores "github.com/manicar2093/filestores"
	mock "github.com/stretchr/testify/mock"
)

// FileStore is an autogenerated mock type for the FileStore type
type FileStore struct {
	mock.Mock
}

type FileStore_Expecter struct {
	mock *mock.Mock
}

func (_m *FileStore) EXPECT() *FileStore_Expecter {
	return &FileStore_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: filepath
func (_m *FileStore) Get(filepath string) (filestores.ObjectInfo, error) {
	ret := _m.Called(filepath)

	var r0 filestores.ObjectInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (filestores.ObjectInfo, error)); ok {
		return rf(filepath)
	}
	if rf, ok := ret.Get(0).(func(string) filestores.ObjectInfo); ok {
		r0 = rf(filepath)
	} else {
		r0 = ret.Get(0).(filestores.ObjectInfo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filepath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type FileStore_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - filepath string
func (_e *FileStore_Expecter) Get(filepath interface{}) *FileStore_Get_Call {
	return &FileStore_Get_Call{Call: _e.mock.On("Get", filepath)}
}

func (_c *FileStore_Get_Call) Run(run func(filepath string)) *FileStore_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FileStore_Get_Call) Return(_a0 filestores.ObjectInfo, _a1 error) *FileStore_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileStore_Get_Call) RunAndReturn(run func(string) (filestores.ObjectInfo, error)) *FileStore_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFileStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewFileStore creates a new instance of FileStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFileStore(t mockConstructorTestingTNewFileStore) *FileStore {
	mock := &FileStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
