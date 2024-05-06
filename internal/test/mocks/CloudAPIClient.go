// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bundle "github.com/cerbos/cloud-api/bundle"

	mock "github.com/stretchr/testify/mock"
)

// CloudAPIClient is an autogenerated mock type for the CloudAPIClient type
type CloudAPIClient struct {
	mock.Mock
}

type CloudAPIClient_Expecter struct {
	mock *mock.Mock
}

func (_m *CloudAPIClient) EXPECT() *CloudAPIClient_Expecter {
	return &CloudAPIClient_Expecter{mock: &_m.Mock}
}

// BootstrapBundle provides a mock function with given fields: _a0, _a1
func (_m *CloudAPIClient) BootstrapBundle(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for BootstrapBundle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudAPIClient_BootstrapBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BootstrapBundle'
type CloudAPIClient_BootstrapBundle_Call struct {
	*mock.Call
}

// BootstrapBundle is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *CloudAPIClient_Expecter) BootstrapBundle(_a0 interface{}, _a1 interface{}) *CloudAPIClient_BootstrapBundle_Call {
	return &CloudAPIClient_BootstrapBundle_Call{Call: _e.mock.On("BootstrapBundle", _a0, _a1)}
}

func (_c *CloudAPIClient_BootstrapBundle_Call) Run(run func(_a0 context.Context, _a1 string)) *CloudAPIClient_BootstrapBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CloudAPIClient_BootstrapBundle_Call) Return(_a0 string, _a1 error) *CloudAPIClient_BootstrapBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudAPIClient_BootstrapBundle_Call) RunAndReturn(run func(context.Context, string) (string, error)) *CloudAPIClient_BootstrapBundle_Call {
	_c.Call.Return(run)
	return _c
}

// GetBundle provides a mock function with given fields: _a0, _a1
func (_m *CloudAPIClient) GetBundle(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetBundle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudAPIClient_GetBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBundle'
type CloudAPIClient_GetBundle_Call struct {
	*mock.Call
}

// GetBundle is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *CloudAPIClient_Expecter) GetBundle(_a0 interface{}, _a1 interface{}) *CloudAPIClient_GetBundle_Call {
	return &CloudAPIClient_GetBundle_Call{Call: _e.mock.On("GetBundle", _a0, _a1)}
}

func (_c *CloudAPIClient_GetBundle_Call) Run(run func(_a0 context.Context, _a1 string)) *CloudAPIClient_GetBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CloudAPIClient_GetBundle_Call) Return(_a0 string, _a1 error) *CloudAPIClient_GetBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudAPIClient_GetBundle_Call) RunAndReturn(run func(context.Context, string) (string, error)) *CloudAPIClient_GetBundle_Call {
	_c.Call.Return(run)
	return _c
}

// GetCachedBundle provides a mock function with given fields: _a0
func (_m *CloudAPIClient) GetCachedBundle(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCachedBundle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudAPIClient_GetCachedBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCachedBundle'
type CloudAPIClient_GetCachedBundle_Call struct {
	*mock.Call
}

// GetCachedBundle is a helper method to define mock.On call
//   - _a0 string
func (_e *CloudAPIClient_Expecter) GetCachedBundle(_a0 interface{}) *CloudAPIClient_GetCachedBundle_Call {
	return &CloudAPIClient_GetCachedBundle_Call{Call: _e.mock.On("GetCachedBundle", _a0)}
}

func (_c *CloudAPIClient_GetCachedBundle_Call) Run(run func(_a0 string)) *CloudAPIClient_GetCachedBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CloudAPIClient_GetCachedBundle_Call) Return(_a0 string, _a1 error) *CloudAPIClient_GetCachedBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudAPIClient_GetCachedBundle_Call) RunAndReturn(run func(string) (string, error)) *CloudAPIClient_GetCachedBundle_Call {
	_c.Call.Return(run)
	return _c
}

// WatchBundle provides a mock function with given fields: _a0, _a1
func (_m *CloudAPIClient) WatchBundle(_a0 context.Context, _a1 string) (bundle.WatchHandle, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for WatchBundle")
	}

	var r0 bundle.WatchHandle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bundle.WatchHandle, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bundle.WatchHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bundle.WatchHandle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudAPIClient_WatchBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchBundle'
type CloudAPIClient_WatchBundle_Call struct {
	*mock.Call
}

// WatchBundle is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *CloudAPIClient_Expecter) WatchBundle(_a0 interface{}, _a1 interface{}) *CloudAPIClient_WatchBundle_Call {
	return &CloudAPIClient_WatchBundle_Call{Call: _e.mock.On("WatchBundle", _a0, _a1)}
}

func (_c *CloudAPIClient_WatchBundle_Call) Run(run func(_a0 context.Context, _a1 string)) *CloudAPIClient_WatchBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CloudAPIClient_WatchBundle_Call) Return(_a0 bundle.WatchHandle, _a1 error) *CloudAPIClient_WatchBundle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudAPIClient_WatchBundle_Call) RunAndReturn(run func(context.Context, string) (bundle.WatchHandle, error)) *CloudAPIClient_WatchBundle_Call {
	_c.Call.Return(run)
	return _c
}

// NewCloudAPIClient creates a new instance of CloudAPIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudAPIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudAPIClient {
	mock := &CloudAPIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
