// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/Mldlr/storety/internal/server/models"

	uuid "github.com/google/uuid"
)

// DataService is an autogenerated mock type for the Service type
type DataService struct {
	mock.Mock
}

type DataService_Expecter struct {
	mock *mock.Mock
}

func (_m *DataService) EXPECT() *DataService_Expecter {
	return &DataService_Expecter{mock: &_m.Mock}
}

// CreateBatch provides a mock function with given fields: ctx, userID, dataBatch
func (_m *DataService) CreateBatch(ctx context.Context, userID uuid.UUID, dataBatch []models.Data) error {
	ret := _m.Called(ctx, userID, dataBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.Data) error); ok {
		r0 = rf(ctx, userID, dataBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataService_CreateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBatch'
type DataService_CreateBatch_Call struct {
	*mock.Call
}

// CreateBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - dataBatch []models.Data
func (_e *DataService_Expecter) CreateBatch(ctx interface{}, userID interface{}, dataBatch interface{}) *DataService_CreateBatch_Call {
	return &DataService_CreateBatch_Call{Call: _e.mock.On("CreateBatch", ctx, userID, dataBatch)}
}

func (_c *DataService_CreateBatch_Call) Run(run func(ctx context.Context, userID uuid.UUID, dataBatch []models.Data)) *DataService_CreateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.Data))
	})
	return _c
}

func (_c *DataService_CreateBatch_Call) Return(_a0 error) *DataService_CreateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataService_CreateBatch_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.Data) error) *DataService_CreateBatch_Call {
	_c.Call.Return(run)
	return _c
}

// CreateData provides a mock function with given fields: ctx, userID, _a2
func (_m *DataService) CreateData(ctx context.Context, userID uuid.UUID, _a2 *models.Data) error {
	ret := _m.Called(ctx, userID, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *models.Data) error); ok {
		r0 = rf(ctx, userID, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataService_CreateData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateData'
type DataService_CreateData_Call struct {
	*mock.Call
}

// CreateData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - _a2 *models.Data
func (_e *DataService_Expecter) CreateData(ctx interface{}, userID interface{}, _a2 interface{}) *DataService_CreateData_Call {
	return &DataService_CreateData_Call{Call: _e.mock.On("CreateData", ctx, userID, _a2)}
}

func (_c *DataService_CreateData_Call) Run(run func(ctx context.Context, userID uuid.UUID, _a2 *models.Data)) *DataService_CreateData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(*models.Data))
	})
	return _c
}

func (_c *DataService_CreateData_Call) Return(_a0 error) *DataService_CreateData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataService_CreateData_Call) RunAndReturn(run func(context.Context, uuid.UUID, *models.Data) error) *DataService_CreateData_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteData provides a mock function with given fields: ctx, userID, name
func (_m *DataService) DeleteData(ctx context.Context, userID uuid.UUID, name string) error {
	ret := _m.Called(ctx, userID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, userID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataService_DeleteData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteData'
type DataService_DeleteData_Call struct {
	*mock.Call
}

// DeleteData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - name string
func (_e *DataService_Expecter) DeleteData(ctx interface{}, userID interface{}, name interface{}) *DataService_DeleteData_Call {
	return &DataService_DeleteData_Call{Call: _e.mock.On("DeleteData", ctx, userID, name)}
}

func (_c *DataService_DeleteData_Call) Run(run func(ctx context.Context, userID uuid.UUID, name string)) *DataService_DeleteData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *DataService_DeleteData_Call) Return(_a0 error) *DataService_DeleteData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataService_DeleteData_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) error) *DataService_DeleteData_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataContent provides a mock function with given fields: ctx, userID, name
func (_m *DataService) GetDataContent(ctx context.Context, userID uuid.UUID, name string) ([]byte, string, error) {
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

// DataService_GetDataContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataContent'
type DataService_GetDataContent_Call struct {
	*mock.Call
}

// GetDataContent is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - name string
func (_e *DataService_Expecter) GetDataContent(ctx interface{}, userID interface{}, name interface{}) *DataService_GetDataContent_Call {
	return &DataService_GetDataContent_Call{Call: _e.mock.On("GetDataContent", ctx, userID, name)}
}

func (_c *DataService_GetDataContent_Call) Run(run func(ctx context.Context, userID uuid.UUID, name string)) *DataService_GetDataContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *DataService_GetDataContent_Call) Return(_a0 []byte, _a1 string, _a2 error) *DataService_GetDataContent_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DataService_GetDataContent_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) ([]byte, string, error)) *DataService_GetDataContent_Call {
	_c.Call.Return(run)
	return _c
}

// GetSyncData provides a mock function with given fields: ctx, userID, syncData
func (_m *DataService) GetSyncData(ctx context.Context, userID uuid.UUID, syncData []models.SyncData) ([]models.Data, []string, error) {
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

// DataService_GetSyncData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSyncData'
type DataService_GetSyncData_Call struct {
	*mock.Call
}

// GetSyncData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - syncData []models.SyncData
func (_e *DataService_Expecter) GetSyncData(ctx interface{}, userID interface{}, syncData interface{}) *DataService_GetSyncData_Call {
	return &DataService_GetSyncData_Call{Call: _e.mock.On("GetSyncData", ctx, userID, syncData)}
}

func (_c *DataService_GetSyncData_Call) Run(run func(ctx context.Context, userID uuid.UUID, syncData []models.SyncData)) *DataService_GetSyncData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.SyncData))
	})
	return _c
}

func (_c *DataService_GetSyncData_Call) Return(_a0 []models.Data, _a1 []string, _a2 error) *DataService_GetSyncData_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DataService_GetSyncData_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.SyncData) ([]models.Data, []string, error)) *DataService_GetSyncData_Call {
	_c.Call.Return(run)
	return _c
}

// ListData provides a mock function with given fields: ctx, userID
func (_m *DataService) ListData(ctx context.Context, userID uuid.UUID) ([]models.DataInfo, error) {
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

// DataService_ListData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListData'
type DataService_ListData_Call struct {
	*mock.Call
}

// ListData is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *DataService_Expecter) ListData(ctx interface{}, userID interface{}) *DataService_ListData_Call {
	return &DataService_ListData_Call{Call: _e.mock.On("ListData", ctx, userID)}
}

func (_c *DataService_ListData_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *DataService_ListData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *DataService_ListData_Call) Return(_a0 []models.DataInfo, _a1 error) *DataService_ListData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataService_ListData_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]models.DataInfo, error)) *DataService_ListData_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBatch provides a mock function with given fields: ctx, userID, dataBatch
func (_m *DataService) UpdateBatch(ctx context.Context, userID uuid.UUID, dataBatch []models.Data) error {
	ret := _m.Called(ctx, userID, dataBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []models.Data) error); ok {
		r0 = rf(ctx, userID, dataBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataService_UpdateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBatch'
type DataService_UpdateBatch_Call struct {
	*mock.Call
}

// UpdateBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - dataBatch []models.Data
func (_e *DataService_Expecter) UpdateBatch(ctx interface{}, userID interface{}, dataBatch interface{}) *DataService_UpdateBatch_Call {
	return &DataService_UpdateBatch_Call{Call: _e.mock.On("UpdateBatch", ctx, userID, dataBatch)}
}

func (_c *DataService_UpdateBatch_Call) Run(run func(ctx context.Context, userID uuid.UUID, dataBatch []models.Data)) *DataService_UpdateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].([]models.Data))
	})
	return _c
}

func (_c *DataService_UpdateBatch_Call) Return(_a0 error) *DataService_UpdateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataService_UpdateBatch_Call) RunAndReturn(run func(context.Context, uuid.UUID, []models.Data) error) *DataService_UpdateBatch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDataService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataService creates a new instance of DataService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataService(t mockConstructorTestingTNewDataService) *DataService {
	mock := &DataService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
