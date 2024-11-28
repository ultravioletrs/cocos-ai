// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	json "encoding/json"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// SendEvent provides a mock function with given fields: event, status, details
func (_m *Service) SendEvent(event string, status string, details json.RawMessage) error {
	ret := _m.Called(event, status, details)

	if len(ret) == 0 {
		panic("no return value specified for SendEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, json.RawMessage) error); ok {
		r0 = rf(event, status, details)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_SendEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEvent'
type Service_SendEvent_Call struct {
	*mock.Call
}

// SendEvent is a helper method to define mock.On call
//   - event string
//   - status string
//   - details json.RawMessage
func (_e *Service_Expecter) SendEvent(event interface{}, status interface{}, details interface{}) *Service_SendEvent_Call {
	return &Service_SendEvent_Call{Call: _e.mock.On("SendEvent", event, status, details)}
}

func (_c *Service_SendEvent_Call) Run(run func(event string, status string, details json.RawMessage)) *Service_SendEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(json.RawMessage))
	})
	return _c
}

func (_c *Service_SendEvent_Call) Return(_a0 error) *Service_SendEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_SendEvent_Call) RunAndReturn(run func(string, string, json.RawMessage) error) *Service_SendEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
