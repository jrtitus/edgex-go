// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddNotification provides a mock function with given fields: n
func (_m *DBClient) AddNotification(n models.Notification) (string, error) {
	ret := _m.Called(n)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Notification) string); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Notification) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSubscription provides a mock function with given fields: s
func (_m *DBClient) AddSubscription(s models.Subscription) (string, error) {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Subscription) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Subscription) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddTransmission provides a mock function with given fields: t
func (_m *DBClient) AddTransmission(t models.Transmission) (string, error) {
	ret := _m.Called(t)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Transmission) string); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Transmission) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cleanup provides a mock function with given fields:
func (_m *DBClient) Cleanup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanupOld provides a mock function with given fields: age
func (_m *DBClient) CleanupOld(age int) error {
	ret := _m.Called(age)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(age)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteNotificationById provides a mock function with given fields: id
func (_m *DBClient) DeleteNotificationById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotificationBySlug provides a mock function with given fields: id
func (_m *DBClient) DeleteNotificationBySlug(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotificationsOld provides a mock function with given fields: age
func (_m *DBClient) DeleteNotificationsOld(age int) error {
	ret := _m.Called(age)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(age)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubscriptionById provides a mock function with given fields: id
func (_m *DBClient) DeleteSubscriptionById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubscriptionBySlug provides a mock function with given fields: id
func (_m *DBClient) DeleteSubscriptionBySlug(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransmission provides a mock function with given fields: age, status
func (_m *DBClient) DeleteTransmission(age int64, status models.TransmissionStatus) error {
	ret := _m.Called(age, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, models.TransmissionStatus) error); ok {
		r0 = rf(age, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNewNormalNotifications provides a mock function with given fields: limit
func (_m *DBClient) GetNewNormalNotifications(limit int) ([]models.Notification, error) {
	ret := _m.Called(limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(int) []models.Notification); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewNotifications provides a mock function with given fields: limit
func (_m *DBClient) GetNewNotifications(limit int) ([]models.Notification, error) {
	ret := _m.Called(limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(int) []models.Notification); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationById provides a mock function with given fields: id
func (_m *DBClient) GetNotificationById(id string) (models.Notification, error) {
	ret := _m.Called(id)

	var r0 models.Notification
	if rf, ok := ret.Get(0).(func(string) models.Notification); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Notification)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationBySender provides a mock function with given fields: sender, limit
func (_m *DBClient) GetNotificationBySender(sender string, limit int) ([]models.Notification, error) {
	ret := _m.Called(sender, limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(string, int) []models.Notification); ok {
		r0 = rf(sender, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(sender, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationBySlug provides a mock function with given fields: slug
func (_m *DBClient) GetNotificationBySlug(slug string) (models.Notification, error) {
	ret := _m.Called(slug)

	var r0 models.Notification
	if rf, ok := ret.Get(0).(func(string) models.Notification); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(models.Notification)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotifications provides a mock function with given fields:
func (_m *DBClient) GetNotifications() ([]models.Notification, error) {
	ret := _m.Called()

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func() []models.Notification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
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

// GetNotificationsByEnd provides a mock function with given fields: end, limit
func (_m *DBClient) GetNotificationsByEnd(end int64, limit int) ([]models.Notification, error) {
	ret := _m.Called(end, limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(int64, int) []models.Notification); ok {
		r0 = rf(end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationsByLabels provides a mock function with given fields: labels, limit
func (_m *DBClient) GetNotificationsByLabels(labels []string, limit int) ([]models.Notification, error) {
	ret := _m.Called(labels, limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func([]string, int) []models.Notification); ok {
		r0 = rf(labels, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, int) error); ok {
		r1 = rf(labels, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationsByStart provides a mock function with given fields: start, limit
func (_m *DBClient) GetNotificationsByStart(start int64, limit int) ([]models.Notification, error) {
	ret := _m.Called(start, limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(int64, int) []models.Notification); ok {
		r0 = rf(start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(start, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotificationsByStartEnd provides a mock function with given fields: start, end, limit
func (_m *DBClient) GetNotificationsByStartEnd(start int64, end int64, limit int) ([]models.Notification, error) {
	ret := _m.Called(start, end, limit)

	var r0 []models.Notification
	if rf, ok := ret.Get(0).(func(int64, int64, int) []models.Notification); ok {
		r0 = rf(start, end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(start, end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByCategories provides a mock function with given fields: categories
func (_m *DBClient) GetSubscriptionByCategories(categories []string) ([]models.Subscription, error) {
	ret := _m.Called(categories)

	var r0 []models.Subscription
	if rf, ok := ret.Get(0).(func([]string) []models.Subscription); ok {
		r0 = rf(categories)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(categories)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByCategoriesLabels provides a mock function with given fields: categories, labels
func (_m *DBClient) GetSubscriptionByCategoriesLabels(categories []string, labels []string) ([]models.Subscription, error) {
	ret := _m.Called(categories, labels)

	var r0 []models.Subscription
	if rf, ok := ret.Get(0).(func([]string, []string) []models.Subscription); ok {
		r0 = rf(categories, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, []string) error); ok {
		r1 = rf(categories, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionById provides a mock function with given fields: id
func (_m *DBClient) GetSubscriptionById(id string) (models.Subscription, error) {
	ret := _m.Called(id)

	var r0 models.Subscription
	if rf, ok := ret.Get(0).(func(string) models.Subscription); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Subscription)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByLabels provides a mock function with given fields: labels
func (_m *DBClient) GetSubscriptionByLabels(labels []string) ([]models.Subscription, error) {
	ret := _m.Called(labels)

	var r0 []models.Subscription
	if rf, ok := ret.Get(0).(func([]string) []models.Subscription); ok {
		r0 = rf(labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionByReceiver provides a mock function with given fields: receiver
func (_m *DBClient) GetSubscriptionByReceiver(receiver string) ([]models.Subscription, error) {
	ret := _m.Called(receiver)

	var r0 []models.Subscription
	if rf, ok := ret.Get(0).(func(string) []models.Subscription); ok {
		r0 = rf(receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionBySlug provides a mock function with given fields: slug
func (_m *DBClient) GetSubscriptionBySlug(slug string) (models.Subscription, error) {
	ret := _m.Called(slug)

	var r0 models.Subscription
	if rf, ok := ret.Get(0).(func(string) models.Subscription); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(models.Subscription)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptions provides a mock function with given fields:
func (_m *DBClient) GetSubscriptions() ([]models.Subscription, error) {
	ret := _m.Called()

	var r0 []models.Subscription
	if rf, ok := ret.Get(0).(func() []models.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
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

// GetTransmissionById provides a mock function with given fields: id
func (_m *DBClient) GetTransmissionById(id string) (models.Transmission, error) {
	ret := _m.Called(id)

	var r0 models.Transmission
	if rf, ok := ret.Get(0).(func(string) models.Transmission); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Transmission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByEnd provides a mock function with given fields: end, limit
func (_m *DBClient) GetTransmissionsByEnd(end int64, limit int) ([]models.Transmission, error) {
	ret := _m.Called(end, limit)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(int64, int) []models.Transmission); ok {
		r0 = rf(end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByNotificationSlug provides a mock function with given fields: slug, limit
func (_m *DBClient) GetTransmissionsByNotificationSlug(slug string, limit int) ([]models.Transmission, error) {
	ret := _m.Called(slug, limit)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(string, int) []models.Transmission); ok {
		r0 = rf(slug, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(slug, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByNotificationSlugAndStartEnd provides a mock function with given fields: slug, start, end, limit
func (_m *DBClient) GetTransmissionsByNotificationSlugAndStartEnd(slug string, start int64, end int64, limit int) ([]models.Transmission, error) {
	ret := _m.Called(slug, start, end, limit)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(string, int64, int64, int) []models.Transmission); ok {
		r0 = rf(slug, start, end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, int64, int) error); ok {
		r1 = rf(slug, start, end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByStart provides a mock function with given fields: start, limit
func (_m *DBClient) GetTransmissionsByStart(start int64, limit int) ([]models.Transmission, error) {
	ret := _m.Called(start, limit)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(int64, int) []models.Transmission); ok {
		r0 = rf(start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(start, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByStartEnd provides a mock function with given fields: start, end, limit
func (_m *DBClient) GetTransmissionsByStartEnd(start int64, end int64, limit int) ([]models.Transmission, error) {
	ret := _m.Called(start, end, limit)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(int64, int64, int) []models.Transmission); ok {
		r0 = rf(start, end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(start, end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransmissionsByStatus provides a mock function with given fields: limit, status
func (_m *DBClient) GetTransmissionsByStatus(limit int, status models.TransmissionStatus) ([]models.Transmission, error) {
	ret := _m.Called(limit, status)

	var r0 []models.Transmission
	if rf, ok := ret.Get(0).(func(int, models.TransmissionStatus) []models.Transmission); ok {
		r0 = rf(limit, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, models.TransmissionStatus) error); ok {
		r1 = rf(limit, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkNotificationProcessed provides a mock function with given fields: n
func (_m *DBClient) MarkNotificationProcessed(n models.Notification) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Notification) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotification provides a mock function with given fields: n
func (_m *DBClient) UpdateNotification(n models.Notification) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Notification) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSubscription provides a mock function with given fields: s
func (_m *DBClient) UpdateSubscription(s models.Subscription) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Subscription) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTransmission provides a mock function with given fields: t
func (_m *DBClient) UpdateTransmission(t models.Transmission) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Transmission) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
