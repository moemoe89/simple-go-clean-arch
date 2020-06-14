// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/moemoe89/go-clean-arch-ratu/api/v1/api_struct/model"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Count provides a mock function with given fields: filter, where
func (_m *Repository) Count(filter map[string]interface{}, where string) (int, error) {
	ret := _m.Called(filter, where)

	var r0 int
	if rf, ok := ret.Get(0).(func(map[string]interface{}, string) int); ok {
		r0 = rf(filter, where)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, string) error); ok {
		r1 = rf(filter, where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: userReq
func (_m *Repository) Create(userReq *model.UserModel) (*model.UserModel, error) {
	ret := _m.Called(userReq)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(*model.UserModel) *model.UserModel); ok {
		r0 = rf(userReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserModel) error); ok {
		r1 = rf(userReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: filter, where, orderBy, selectField
func (_m *Repository) Get(filter map[string]interface{}, where string, orderBy string, selectField string) ([]*model.UserModel, error) {
	ret := _m.Called(filter, where, orderBy, selectField)

	var r0 []*model.UserModel
	if rf, ok := ret.Get(0).(func(map[string]interface{}, string, string, string) []*model.UserModel); ok {
		r0 = rf(filter, where, orderBy, selectField)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, string, string, string) error); ok {
		r1 = rf(filter, where, orderBy, selectField)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id, selectField
func (_m *Repository) GetByID(id string, selectField string) (*model.UserModel, error) {
	ret := _m.Called(id, selectField)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(string, string) *model.UserModel); ok {
		r0 = rf(id, selectField)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, selectField)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userReq
func (_m *Repository) Update(userReq *model.UserModel) (*model.UserModel, error) {
	ret := _m.Called(userReq)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(*model.UserModel) *model.UserModel); ok {
		r0 = rf(userReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserModel) error); ok {
		r1 = rf(userReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
