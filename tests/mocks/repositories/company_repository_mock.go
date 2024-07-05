// Code generated by mockery v2.43.2. DO NOT EDIT.

package repositories_mocks

import (
	models "hex/mjolnir-core/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

// MockCompanyRepository is an autogenerated mock type for the CompanyRepository type
type MockCompanyRepository struct {
	mock.Mock
}

type MockCompanyRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompanyRepository) EXPECT() *MockCompanyRepository_Expecter {
	return &MockCompanyRepository_Expecter{mock: &_m.Mock}
}

// FindByNameOrCreate provides a mock function with given fields: name
func (_m *MockCompanyRepository) FindByNameOrCreate(name string) (models.Company, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindByNameOrCreate")
	}

	var r0 models.Company
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Company, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.Company); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Company)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCompanyRepository_FindByNameOrCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByNameOrCreate'
type MockCompanyRepository_FindByNameOrCreate_Call struct {
	*mock.Call
}

// FindByNameOrCreate is a helper method to define mock.On call
//   - name string
func (_e *MockCompanyRepository_Expecter) FindByNameOrCreate(name interface{}) *MockCompanyRepository_FindByNameOrCreate_Call {
	return &MockCompanyRepository_FindByNameOrCreate_Call{Call: _e.mock.On("FindByNameOrCreate", name)}
}

func (_c *MockCompanyRepository_FindByNameOrCreate_Call) Run(run func(name string)) *MockCompanyRepository_FindByNameOrCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCompanyRepository_FindByNameOrCreate_Call) Return(_a0 models.Company, _a1 error) *MockCompanyRepository_FindByNameOrCreate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCompanyRepository_FindByNameOrCreate_Call) RunAndReturn(run func(string) (models.Company, error)) *MockCompanyRepository_FindByNameOrCreate_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: company
func (_m *MockCompanyRepository) Update(company models.Company) (models.Company, error) {
	ret := _m.Called(company)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.Company
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Company) (models.Company, error)); ok {
		return rf(company)
	}
	if rf, ok := ret.Get(0).(func(models.Company) models.Company); ok {
		r0 = rf(company)
	} else {
		r0 = ret.Get(0).(models.Company)
	}

	if rf, ok := ret.Get(1).(func(models.Company) error); ok {
		r1 = rf(company)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCompanyRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCompanyRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - company models.Company
func (_e *MockCompanyRepository_Expecter) Update(company interface{}) *MockCompanyRepository_Update_Call {
	return &MockCompanyRepository_Update_Call{Call: _e.mock.On("Update", company)}
}

func (_c *MockCompanyRepository_Update_Call) Run(run func(company models.Company)) *MockCompanyRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Company))
	})
	return _c
}

func (_c *MockCompanyRepository_Update_Call) Return(_a0 models.Company, _a1 error) *MockCompanyRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCompanyRepository_Update_Call) RunAndReturn(run func(models.Company) (models.Company, error)) *MockCompanyRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompanyRepository creates a new instance of MockCompanyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompanyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompanyRepository {
	mock := &MockCompanyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
