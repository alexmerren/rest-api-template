// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// Configurator is an autogenerated mock type for the Configurator type
type Configurator struct {
	mock.Mock
}

// GetInt provides a mock function with given fields: name
func (_m *Configurator) GetInt(name string) (int, error) {
	ret := _m.Called(name)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetString provides a mock function with given fields: name
func (_m *Configurator) GetString(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewConfigurator creates a new instance of Configurator. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigurator(t testing.TB) *Configurator {
	mock := &Configurator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}