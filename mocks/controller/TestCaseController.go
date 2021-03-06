// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// TestCaseController is an autogenerated mock type for the TestCaseController type
type TestCaseController struct {
	mock.Mock
}

// HandleDeleteTestCase provides a mock function with given fields: w, r
func (_m *TestCaseController) HandleDeleteTestCase(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// HandleGetTestCase provides a mock function with given fields: w, r
func (_m *TestCaseController) HandleGetTestCase(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// HandleGetTestCases provides a mock function with given fields: w, r
func (_m *TestCaseController) HandleGetTestCases(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// HandlePostTestCase provides a mock function with given fields: w, r
func (_m *TestCaseController) HandlePostTestCase(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// HandlePutTestCase provides a mock function with given fields: w, r
func (_m *TestCaseController) HandlePutTestCase(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
