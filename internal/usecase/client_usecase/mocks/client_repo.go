// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "medicalCenter/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ClientRepo is an autogenerated mock type for the clientRepo type
type ClientRepo struct {
	mock.Mock
}

type ClientRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientRepo) EXPECT() *ClientRepo_Expecter {
	return &ClientRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: client
func (_m *ClientRepo) Create(client *domain.Client) (*domain.Client, error) {
	ret := _m.Called(client)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Client) (*domain.Client, error)); ok {
		return rf(client)
	}
	if rf, ok := ret.Get(0).(func(*domain.Client) *domain.Client); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ClientRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - client *domain.Client
func (_e *ClientRepo_Expecter) Create(client interface{}) *ClientRepo_Create_Call {
	return &ClientRepo_Create_Call{Call: _e.mock.On("Create", client)}
}

func (_c *ClientRepo_Create_Call) Run(run func(client *domain.Client)) *ClientRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Client))
	})
	return _c
}

func (_c *ClientRepo_Create_Call) Return(_a0 *domain.Client, _a1 error) *ClientRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientRepo_Create_Call) RunAndReturn(run func(*domain.Client) (*domain.Client, error)) *ClientRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *ClientRepo) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ClientRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int
func (_e *ClientRepo_Expecter) Delete(id interface{}) *ClientRepo_Delete_Call {
	return &ClientRepo_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *ClientRepo_Delete_Call) Run(run func(id int)) *ClientRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ClientRepo_Delete_Call) Return(_a0 error) *ClientRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientRepo_Delete_Call) RunAndReturn(run func(int) error) *ClientRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetClientByID provides a mock function with given fields: id
func (_m *ClientRepo) GetClientByID(id int) (*domain.Client, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetClientByID")
	}

	var r0 *domain.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Client, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Client); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRepo_GetClientByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClientByID'
type ClientRepo_GetClientByID_Call struct {
	*mock.Call
}

// GetClientByID is a helper method to define mock.On call
//   - id int
func (_e *ClientRepo_Expecter) GetClientByID(id interface{}) *ClientRepo_GetClientByID_Call {
	return &ClientRepo_GetClientByID_Call{Call: _e.mock.On("GetClientByID", id)}
}

func (_c *ClientRepo_GetClientByID_Call) Run(run func(id int)) *ClientRepo_GetClientByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ClientRepo_GetClientByID_Call) Return(_a0 *domain.Client, _a1 error) *ClientRepo_GetClientByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientRepo_GetClientByID_Call) RunAndReturn(run func(int) (*domain.Client, error)) *ClientRepo_GetClientByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: client
func (_m *ClientRepo) Update(client *domain.Client) (*domain.Client, error) {
	ret := _m.Called(client)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Client) (*domain.Client, error)); ok {
		return rf(client)
	}
	if rf, ok := ret.Get(0).(func(*domain.Client) *domain.Client); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ClientRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - client *domain.Client
func (_e *ClientRepo_Expecter) Update(client interface{}) *ClientRepo_Update_Call {
	return &ClientRepo_Update_Call{Call: _e.mock.On("Update", client)}
}

func (_c *ClientRepo_Update_Call) Run(run func(client *domain.Client)) *ClientRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Client))
	})
	return _c
}

func (_c *ClientRepo_Update_Call) Return(_a0 *domain.Client, _a1 error) *ClientRepo_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientRepo_Update_Call) RunAndReturn(run func(*domain.Client) (*domain.Client, error)) *ClientRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientRepo creates a new instance of ClientRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientRepo {
	mock := &ClientRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
