// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddInterval provides a mock function with given fields: interval
func (_m *DBClient) AddInterval(interval models.Interval) (models.Interval, errors.EdgeX) {
	ret := _m.Called(interval)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(models.Interval) models.Interval); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.Interval) errors.EdgeX); ok {
		r1 = rf(interval)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddIntervalAction provides a mock function with given fields: e
func (_m *DBClient) AddIntervalAction(e models.IntervalAction) (models.IntervalAction, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.IntervalAction
	if rf, ok := ret.Get(0).(func(models.IntervalAction) models.IntervalAction); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.IntervalAction)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(models.IntervalAction) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllIntervalActions provides a mock function with given fields: offset, limit
func (_m *DBClient) AllIntervalActions(offset int, limit int) ([]models.IntervalAction, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(int, int) []models.IntervalAction); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllIntervals provides a mock function with given fields: offset, limit
func (_m *DBClient) AllIntervals(offset int, limit int) ([]models.Interval, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	var r0 []models.Interval
	if rf, ok := ret.Get(0).(func(int, int) []models.Interval); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Interval)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteIntervalActionByName provides a mock function with given fields: name
func (_m *DBClient) DeleteIntervalActionByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteIntervalByName provides a mock function with given fields: name
func (_m *DBClient) DeleteIntervalByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// IntervalActionById provides a mock function with given fields: id
func (_m *DBClient) IntervalActionById(id string) (models.IntervalAction, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) models.IntervalAction); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.IntervalAction)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalActionByName provides a mock function with given fields: name
func (_m *DBClient) IntervalActionByName(name string) (models.IntervalAction, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) models.IntervalAction); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.IntervalAction)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalActionTotalCount provides a mock function with given fields:
func (_m *DBClient) IntervalActionTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalActionsByIntervalName provides a mock function with given fields: offset, limit, IntervalName
func (_m *DBClient) IntervalActionsByIntervalName(offset int, limit int, IntervalName string) ([]models.IntervalAction, errors.EdgeX) {
	ret := _m.Called(offset, limit, IntervalName)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(int, int, string) []models.IntervalAction); ok {
		r0 = rf(offset, limit, IntervalName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, IntervalName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalById provides a mock function with given fields: id
func (_m *DBClient) IntervalById(id string) (models.Interval, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalByName provides a mock function with given fields: name
func (_m *DBClient) IntervalByName(name string) (models.Interval, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalTotalCount provides a mock function with given fields:
func (_m *DBClient) IntervalTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateInterval provides a mock function with given fields: interval
func (_m *DBClient) UpdateInterval(interval models.Interval) errors.EdgeX {
	ret := _m.Called(interval)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Interval) errors.EdgeX); ok {
		r0 = rf(interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// UpdateIntervalAction provides a mock function with given fields: action
func (_m *DBClient) UpdateIntervalAction(action models.IntervalAction) errors.EdgeX {
	ret := _m.Called(action)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.IntervalAction) errors.EdgeX); ok {
		r0 = rf(action)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}
