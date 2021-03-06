// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fikrirnurhidayat/codeotd/app/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TestCaseRepository is an autogenerated mock type for the TestCaseRepository type
type TestCaseRepository struct {
	mock.Mock
}

// CreateTestCase provides a mock function with given fields: _a0, _a1
func (_m *TestCaseRepository) CreateTestCase(_a0 context.Context, _a1 *entity.TestCase) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TestCase) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTestCase provides a mock function with given fields: _a0, _a1
func (_m *TestCaseRepository) DeleteTestCase(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTestCase provides a mock function with given fields: _a0, _a1
func (_m *TestCaseRepository) GetTestCase(_a0 context.Context, _a1 uuid.UUID) (entity.TestCase, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.TestCase
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.TestCase); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.TestCase)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTestCases provides a mock function with given fields: ctx, challengeId, offset
func (_m *TestCaseRepository) ListTestCases(ctx context.Context, challengeId uuid.UUID, offset entity.Offset) ([]entity.TestCase, int32, error) {
	ret := _m.Called(ctx, challengeId, offset)

	var r0 []entity.TestCase
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, entity.Offset) []entity.TestCase); ok {
		r0 = rf(ctx, challengeId, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.TestCase)
		}
	}

	var r1 int32
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, entity.Offset) int32); ok {
		r1 = rf(ctx, challengeId, offset)
	} else {
		r1 = ret.Get(1).(int32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, entity.Offset) error); ok {
		r2 = rf(ctx, challengeId, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTestCase provides a mock function with given fields: _a0, _a1, _a2
func (_m *TestCaseRepository) UpdateTestCase(_a0 context.Context, _a1 uuid.UUID, _a2 *entity.TestCase) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *entity.TestCase) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
