// Code generated by mockery v2.43.2. DO NOT EDIT.

package repositories_mocks

import (
	models "hex/mjolnir-core/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockUserRepository) CreateUser(user models.User) (models.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(models.User) (models.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user models.User
func (_e *MockUserRepository_Expecter) CreateUser(user interface{}) *MockUserRepository_CreateUser_Call {
	return &MockUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockUserRepository_CreateUser_Call) Run(run func(user models.User)) *MockUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.User))
	})
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) Return(_a0 models.User, _a1 error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) RunAndReturn(run func(models.User) (models.User, error)) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetByInviteId provides a mock function with given fields: inviteId
func (_m *MockUserRepository) GetByInviteId(inviteId string) (models.User, error) {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for GetByInviteId")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(inviteId)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(inviteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetByInviteId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByInviteId'
type MockUserRepository_GetByInviteId_Call struct {
	*mock.Call
}

// GetByInviteId is a helper method to define mock.On call
//   - inviteId string
func (_e *MockUserRepository_Expecter) GetByInviteId(inviteId interface{}) *MockUserRepository_GetByInviteId_Call {
	return &MockUserRepository_GetByInviteId_Call{Call: _e.mock.On("GetByInviteId", inviteId)}
}

func (_c *MockUserRepository_GetByInviteId_Call) Run(run func(inviteId string)) *MockUserRepository_GetByInviteId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetByInviteId_Call) Return(_a0 models.User, _a1 error) *MockUserRepository_GetByInviteId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetByInviteId_Call) RunAndReturn(run func(string) (models.User, error)) *MockUserRepository_GetByInviteId_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetUserByEmail(email string) (models.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockUserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUserRepository_Expecter) GetUserByEmail(email interface{}) *MockUserRepository_GetUserByEmail_Call {
	return &MockUserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *MockUserRepository_GetUserByEmail_Call) Run(run func(email string)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) Return(_a0 models.User, _a1 error) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) RunAndReturn(run func(string) (models.User, error)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: user
func (_m *MockUserRepository) Update(user models.User) (models.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(models.User) (models.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - user models.User
func (_e *MockUserRepository_Expecter) Update(user interface{}) *MockUserRepository_Update_Call {
	return &MockUserRepository_Update_Call{Call: _e.mock.On("Update", user)}
}

func (_c *MockUserRepository_Update_Call) Run(run func(user models.User)) *MockUserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.User))
	})
	return _c
}

func (_c *MockUserRepository_Update_Call) Return(_a0 models.User, _a1 error) *MockUserRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Update_Call) RunAndReturn(run func(models.User) (models.User, error)) *MockUserRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
