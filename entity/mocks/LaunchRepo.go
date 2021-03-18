// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "github.com/soyantonio-w/academy-go-q12021/entity"
	mock "github.com/stretchr/testify/mock"
)

// LaunchRepo is an autogenerated mock type for the LaunchRepo type
type LaunchRepo struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *LaunchRepo) Get(id entity.LaunchId) (entity.Launch, error) {
	ret := _m.Called(id)

	var r0 entity.Launch
	if rf, ok := ret.Get(0).(func(entity.LaunchId) entity.Launch); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Launch)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.LaunchId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLaunches provides a mock function with given fields:
func (_m *LaunchRepo) GetLaunches() ([]entity.Launch, error) {
	ret := _m.Called()

	var r0 []entity.Launch
	if rf, ok := ret.Get(0).(func() []entity.Launch); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Launch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncAll provides a mock function with given fields: launches
func (_m *LaunchRepo) SyncAll(launches []entity.Launch) error {
	ret := _m.Called(launches)

	var r0 error
	if rf, ok := ret.Get(0).(func([]entity.Launch) error); ok {
		r0 = rf(launches)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
