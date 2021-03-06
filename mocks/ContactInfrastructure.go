// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ContactInfrastructure is an autogenerated mock type for the ContactInfrastructure type
type ContactInfrastructure struct {
	mock.Mock
}

// Create provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) Create(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Delete provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) Delete(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Health provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) Health(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// ReadMany provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) ReadMany(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// ReadOne provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) ReadOne(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Update provides a mock function with given fields: w, r
func (_m *ContactInfrastructure) Update(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewContactInfrastructure creates a new instance of ContactInfrastructure. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewContactInfrastructure(t testing.TB) *ContactInfrastructure {
	mock := &ContactInfrastructure{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
