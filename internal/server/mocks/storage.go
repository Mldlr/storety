// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/Mldlr/storety/internal/server/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// CreateBatch provides a mock function with given fields: ctx, userID, dataBatch
func (_m *Storage) CreateBatch(ctx context.Context, userID uuid.UUID, dataBatch []models.Data) error {
	ret := _m.Called(ctx, userID, dataBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.Data) error); ok {
		r0 = rf(ctx, userID, dataBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBatch'
type Storage_CreateBatch_Call struct {
	*mock.Call
}

// CreateBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - dataBatch []models.Data
func (_e *Storage_Expecter) CreateBatch(ctx interface{}, userID interface{}, dataBatch interface{}) *Storage_CreateBatch_Call {
	return &Storage_CreateBatch_Call{Call: _e.mock.On("CreateBatch", ctx, userID, dataBatch)}
}

func (_c *Storage_CreateBatch_Call) Run(run func(ctx context.Context, userID uuid.UUID, dataBatch []models.Data)) *Storage_CreateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.Data))
	})
	return _c
}

func (_c *Storage_CreateBatch_Call) Return(_a0 error) *Storage_CreateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_CreateBatch_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.Data) error) *Storage_CreateBatch_Call {
	_c.Call.Return(run)
	return _c
}

// CreateData provides a mock function with given fields: ctx, userID, data
func (_m *Storage) CreateData(ctx context.Context, userID uuid.UUID, data *models.Data) error {
	ret := _m.Called(ctx, userID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *models.Data) error); ok {
		r0 = rf(ctx, userID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateData'
type Storage_CreateData_Call struct {
	*mock.Call
}

// CreateData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - data *models.Data
func (_e *Storage_Expecter) CreateData(ctx interface{}, userID interface{}, data interface{}) *Storage_CreateData_Call {
	return &Storage_CreateData_Call{Call: _e.mock.On("CreateData", ctx, userID, data)}
}

func (_c *Storage_CreateData_Call) Run(run func(ctx context.Context, userID uuid.UUID, data *models.Data)) *Storage_CreateData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(*models.Data))
	})
	return _c
}

func (_c *Storage_CreateData_Call) Return(_a0 error) *Storage_CreateData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_CreateData_Call) RunAndReturn(run func(context.Context, uuid.UUID, *models.Data) error) *Storage_CreateData_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSession provides a mock function with given fields: ctx, session, oldSession
func (_m *Storage) CreateSession(ctx context.Context, session *models.Session, oldSession *models.Session) error {
	ret := _m.Called(ctx, session, oldSession)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Session, *models.Session) error); ok {
		r0 = rf(ctx, session, oldSession)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSession'
type Storage_CreateSession_Call struct {
	*mock.Call
}

// CreateSession is a helper method to define mock.On call
//   - ctx context.Context
//   - session *models.Session
//   - oldSession *models.Session
func (_e *Storage_Expecter) CreateSession(ctx interface{}, session interface{}, oldSession interface{}) *Storage_CreateSession_Call {
	return &Storage_CreateSession_Call{Call: _e.mock.On("CreateSession", ctx, session, oldSession)}
}

func (_c *Storage_CreateSession_Call) Run(run func(ctx context.Context, session *models.Session, oldSession *models.Session)) *Storage_CreateSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Session), args[2].(*models.Session))
	})
	return _c
}

func (_c *Storage_CreateSession_Call) Return(_a0 error) *Storage_CreateSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_CreateSession_Call) RunAndReturn(run func(context.Context, *models.Session, *models.Session) error) *Storage_CreateSession_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storage) CreateUser(ctx context.Context, user *models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Storage_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *models.User
func (_e *Storage_Expecter) CreateUser(ctx interface{}, user interface{}) *Storage_CreateUser_Call {
	return &Storage_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *Storage_CreateUser_Call) Run(run func(ctx context.Context, user *models.User)) *Storage_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *Storage_CreateUser_Call) Return(_a0 error) *Storage_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_CreateUser_Call) RunAndReturn(run func(context.Context, *models.User) error) *Storage_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDataByName provides a mock function with given fields: ctx, userID, name
func (_m *Storage) DeleteDataByName(ctx context.Context, userID uuid.UUID, name string) error {
	ret := _m.Called(ctx, userID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, userID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_DeleteDataByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDataByName'
type Storage_DeleteDataByName_Call struct {
	*mock.Call
}

// DeleteDataByName is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - name string
func (_e *Storage_Expecter) DeleteDataByName(ctx interface{}, userID interface{}, name interface{}) *Storage_DeleteDataByName_Call {
	return &Storage_DeleteDataByName_Call{Call: _e.mock.On("DeleteDataByName", ctx, userID, name)}
}

func (_c *Storage_DeleteDataByName_Call) Run(run func(ctx context.Context, userID uuid.UUID, name string)) *Storage_DeleteDataByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *Storage_DeleteDataByName_Call) Return(_a0 error) *Storage_DeleteDataByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_DeleteDataByName_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) error) *Storage_DeleteDataByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllDataInfo provides a mock function with given fields: ctx, userID
func (_m *Storage) GetAllDataInfo(ctx context.Context, userID uuid.UUID) ([]models.DataInfo, error) {
	ret := _m.Called(ctx, userID)

	var r0 []models.DataInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]models.DataInfo, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []models.DataInfo); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DataInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetAllDataInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllDataInfo'
type Storage_GetAllDataInfo_Call struct {
	*mock.Call
}

// GetAllDataInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *Storage_Expecter) GetAllDataInfo(ctx interface{}, userID interface{}) *Storage_GetAllDataInfo_Call {
	return &Storage_GetAllDataInfo_Call{Call: _e.mock.On("GetAllDataInfo", ctx, userID)}
}

func (_c *Storage_GetAllDataInfo_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *Storage_GetAllDataInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Storage_GetAllDataInfo_Call) Return(_a0 []models.DataInfo, _a1 error) *Storage_GetAllDataInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetAllDataInfo_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]models.DataInfo, error)) *Storage_GetAllDataInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataByUpdateAndHash provides a mock function with given fields: ctx, userID, syncData
func (_m *Storage) GetDataByUpdateAndHash(ctx context.Context, userID uuid.UUID, syncData []models.SyncData) ([]models.Data, []string, error) {
	ret := _m.Called(ctx, userID, syncData)

	var r0 []models.Data
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.SyncData) ([]models.Data, []string, error)); ok {
		return rf(ctx, userID, syncData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.SyncData) []models.Data); ok {
		r0 = rf(ctx, userID, syncData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Data)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, []models.SyncData) []string); ok {
		r1 = rf(ctx, userID, syncData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, []models.SyncData) error); ok {
		r2 = rf(ctx, userID, syncData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Storage_GetDataByUpdateAndHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataByUpdateAndHash'
type Storage_GetDataByUpdateAndHash_Call struct {
	*mock.Call
}

// GetDataByUpdateAndHash is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - syncData []models.SyncData
func (_e *Storage_Expecter) GetDataByUpdateAndHash(ctx interface{}, userID interface{}, syncData interface{}) *Storage_GetDataByUpdateAndHash_Call {
	return &Storage_GetDataByUpdateAndHash_Call{Call: _e.mock.On("GetDataByUpdateAndHash", ctx, userID, syncData)}
}

func (_c *Storage_GetDataByUpdateAndHash_Call) Run(run func(ctx context.Context, userID uuid.UUID, syncData []models.SyncData)) *Storage_GetDataByUpdateAndHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.SyncData))
	})
	return _c
}

func (_c *Storage_GetDataByUpdateAndHash_Call) Return(_a0 []models.Data, _a1 []string, _a2 error) *Storage_GetDataByUpdateAndHash_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Storage_GetDataByUpdateAndHash_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.SyncData) ([]models.Data, []string, error)) *Storage_GetDataByUpdateAndHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataContentByName provides a mock function with given fields: ctx, userID, name
func (_m *Storage) GetDataContentByName(ctx context.Context, userID uuid.UUID, name string) ([]byte, string, error) {
	ret := _m.Called(ctx, userID, name)

	var r0 []byte
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) ([]byte, string, error)); ok {
		return rf(ctx, userID, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) []byte); ok {
		r0 = rf(ctx, userID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) string); ok {
		r1 = rf(ctx, userID, name)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, string) error); ok {
		r2 = rf(ctx, userID, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Storage_GetDataContentByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataContentByName'
type Storage_GetDataContentByName_Call struct {
	*mock.Call
}

// GetDataContentByName is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - name string
func (_e *Storage_Expecter) GetDataContentByName(ctx interface{}, userID interface{}, name interface{}) *Storage_GetDataContentByName_Call {
	return &Storage_GetDataContentByName_Call{Call: _e.mock.On("GetDataContentByName", ctx, userID, name)}
}

func (_c *Storage_GetDataContentByName_Call) Run(run func(ctx context.Context, userID uuid.UUID, name string)) *Storage_GetDataContentByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *Storage_GetDataContentByName_Call) Return(_a0 []byte, _a1 string, _a2 error) *Storage_GetDataContentByName_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Storage_GetDataContentByName_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) ([]byte, string, error)) *Storage_GetDataContentByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetNewData provides a mock function with given fields: ctx, userID, ids
func (_m *Storage) GetNewData(ctx context.Context, userID uuid.UUID, ids []uuid.UUID) ([]models.Data, error) {
	ret := _m.Called(ctx, userID, ids)

	var r0 []models.Data
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []uuid.UUID) ([]models.Data, error)); ok {
		return rf(ctx, userID, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []uuid.UUID) []models.Data); ok {
		r0 = rf(ctx, userID, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Data)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, []uuid.UUID) error); ok {
		r1 = rf(ctx, userID, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetNewData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNewData'
type Storage_GetNewData_Call struct {
	*mock.Call
}

// GetNewData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - ids []uuid.UUID
func (_e *Storage_Expecter) GetNewData(ctx interface{}, userID interface{}, ids interface{}) *Storage_GetNewData_Call {
	return &Storage_GetNewData_Call{Call: _e.mock.On("GetNewData", ctx, userID, ids)}
}

func (_c *Storage_GetNewData_Call) Run(run func(ctx context.Context, userID uuid.UUID, ids []uuid.UUID)) *Storage_GetNewData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]uuid.UUID))
	})
	return _c
}

func (_c *Storage_GetNewData_Call) Return(_a0 []models.Data, _a1 error) *Storage_GetNewData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetNewData_Call) RunAndReturn(run func(context.Context, uuid.UUID, []uuid.UUID) ([]models.Data, error)) *Storage_GetNewData_Call {
	_c.Call.Return(run)
	return _c
}

// GetSession provides a mock function with given fields: ctx, sessionID, refreshToken
func (_m *Storage) GetSession(ctx context.Context, sessionID uuid.UUID, refreshToken string) (uuid.UUID, error) {
	ret := _m.Called(ctx, sessionID, refreshToken)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) (uuid.UUID, error)); ok {
		return rf(ctx, sessionID, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) uuid.UUID); ok {
		r0 = rf(ctx, sessionID, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, sessionID, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSession'
type Storage_GetSession_Call struct {
	*mock.Call
}

// GetSession is a helper method to define mock.On call
//   - ctx context.Context
//   - sessionID uuid.UUID
//   - refreshToken string
func (_e *Storage_Expecter) GetSession(ctx interface{}, sessionID interface{}, refreshToken interface{}) *Storage_GetSession_Call {
	return &Storage_GetSession_Call{Call: _e.mock.On("GetSession", ctx, sessionID, refreshToken)}
}

func (_c *Storage_GetSession_Call) Run(run func(ctx context.Context, sessionID uuid.UUID, refreshToken string)) *Storage_GetSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *Storage_GetSession_Call) Return(_a0 uuid.UUID, _a1 error) *Storage_GetSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_GetSession_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) (uuid.UUID, error)) *Storage_GetSession_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserDataByName provides a mock function with given fields: ctx, username
func (_m *Storage) GetUserDataByName(ctx context.Context, username string) (uuid.UUID, string, string, error) {
	ret := _m.Called(ctx, username)

	var r0 uuid.UUID
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uuid.UUID, string, string, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uuid.UUID); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) string); ok {
		r2 = rf(ctx, username)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string) error); ok {
		r3 = rf(ctx, username)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Storage_GetUserDataByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserDataByName'
type Storage_GetUserDataByName_Call struct {
	*mock.Call
}

// GetUserDataByName is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *Storage_Expecter) GetUserDataByName(ctx interface{}, username interface{}) *Storage_GetUserDataByName_Call {
	return &Storage_GetUserDataByName_Call{Call: _e.mock.On("GetUserDataByName", ctx, username)}
}

func (_c *Storage_GetUserDataByName_Call) Run(run func(ctx context.Context, username string)) *Storage_GetUserDataByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_GetUserDataByName_Call) Return(_a0 uuid.UUID, _a1 string, _a2 string, _a3 error) *Storage_GetUserDataByName_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *Storage_GetUserDataByName_Call) RunAndReturn(run func(context.Context, string) (uuid.UUID, string, string, error)) *Storage_GetUserDataByName_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBatch provides a mock function with given fields: ctx, userID, dataBatch
func (_m *Storage) UpdateBatch(ctx context.Context, userID uuid.UUID, dataBatch []models.Data) error {
	ret := _m.Called(ctx, userID, dataBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.Data) error); ok {
		r0 = rf(ctx, userID, dataBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_UpdateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBatch'
type Storage_UpdateBatch_Call struct {
	*mock.Call
}

// UpdateBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - dataBatch []models.Data
func (_e *Storage_Expecter) UpdateBatch(ctx interface{}, userID interface{}, dataBatch interface{}) *Storage_UpdateBatch_Call {
	return &Storage_UpdateBatch_Call{Call: _e.mock.On("UpdateBatch", ctx, userID, dataBatch)}
}

func (_c *Storage_UpdateBatch_Call) Run(run func(ctx context.Context, userID uuid.UUID, dataBatch []models.Data)) *Storage_UpdateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.Data))
	})
	return _c
}

func (_c *Storage_UpdateBatch_Call) Return(_a0 error) *Storage_UpdateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_UpdateBatch_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.Data) error) *Storage_UpdateBatch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
