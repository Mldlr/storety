// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/Mldlr/storety/internal/server/models"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the Service type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, _a1
func (_m *UserService) CreateUser(ctx context.Context, _a1 *models.User) (*models.Session, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *models.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) (*models.Session, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.Session); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserService_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *models.User
func (_e *UserService_Expecter) CreateUser(ctx interface{}, _a1 interface{}) *UserService_CreateUser_Call {
	return &UserService_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, _a1)}
}

func (_c *UserService_CreateUser_Call) Run(run func(ctx context.Context, _a1 *models.User)) *UserService_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserService_CreateUser_Call) Return(_a0 *models.Session, _a1 error) *UserService_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_CreateUser_Call) RunAndReturn(run func(context.Context, *models.User) (*models.Session, error)) *UserService_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// LogInUser provides a mock function with given fields: ctx, _a1
func (_m *UserService) LogInUser(ctx context.Context, _a1 *models.User) (*models.Session, string, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *models.Session
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) (*models.Session, string, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.Session); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User) string); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *models.User) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UserService_LogInUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogInUser'
type UserService_LogInUser_Call struct {
	*mock.Call
}

// LogInUser is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *models.User
func (_e *UserService_Expecter) LogInUser(ctx interface{}, _a1 interface{}) *UserService_LogInUser_Call {
	return &UserService_LogInUser_Call{Call: _e.mock.On("LogInUser", ctx, _a1)}
}

func (_c *UserService_LogInUser_Call) Run(run func(ctx context.Context, _a1 *models.User)) *UserService_LogInUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserService_LogInUser_Call) Return(_a0 *models.Session, _a1 string, _a2 error) *UserService_LogInUser_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *UserService_LogInUser_Call) RunAndReturn(run func(context.Context, *models.User) (*models.Session, string, error)) *UserService_LogInUser_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshUserSession provides a mock function with given fields: ctx, oldSession
func (_m *UserService) RefreshUserSession(ctx context.Context, oldSession *models.Session) (*models.Session, error) {
	ret := _m.Called(ctx, oldSession)

	var r0 *models.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Session) (*models.Session, error)); ok {
		return rf(ctx, oldSession)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Session) *models.Session); ok {
		r0 = rf(ctx, oldSession)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Session) error); ok {
		r1 = rf(ctx, oldSession)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_RefreshUserSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshUserSession'
type UserService_RefreshUserSession_Call struct {
	*mock.Call
}

// RefreshUserSession is a helper method to define mock.On call
//   - ctx context.Context
//   - oldSession *models.Session
func (_e *UserService_Expecter) RefreshUserSession(ctx interface{}, oldSession interface{}) *UserService_RefreshUserSession_Call {
	return &UserService_RefreshUserSession_Call{Call: _e.mock.On("RefreshUserSession", ctx, oldSession)}
}

func (_c *UserService_RefreshUserSession_Call) Run(run func(ctx context.Context, oldSession *models.Session)) *UserService_RefreshUserSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Session))
	})
	return _c
}

func (_c *UserService_RefreshUserSession_Call) Return(_a0 *models.Session, _a1 error) *UserService_RefreshUserSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_RefreshUserSession_Call) RunAndReturn(run func(context.Context, *models.Session) (*models.Session, error)) *UserService_RefreshUserSession_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
