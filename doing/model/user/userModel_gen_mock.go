// Code generated by MockGen. DO NOT EDIT.
// Source: userModel_gen.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockuserModel is a mock of userModel interface.
type MockuserModel struct {
	ctrl     *gomock.Controller
	recorder *MockuserModelMockRecorder
}

// MockuserModelMockRecorder is the mock recorder for MockuserModel.
type MockuserModelMockRecorder struct {
	mock *MockuserModel
}

// NewMockuserModel creates a new mock instance.
func NewMockuserModel(ctrl *gomock.Controller) *MockuserModel {
	mock := &MockuserModel{ctrl: ctrl}
	mock.recorder = &MockuserModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserModel) EXPECT() *MockuserModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockuserModel) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockuserModelMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockuserModel)(nil).Delete), ctx, id)
}

// FindOne mocks base method.
func (m *MockuserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockuserModelMockRecorder) FindOne(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockuserModel)(nil).FindOne), ctx, id)
}

// FindOneByMobile mocks base method.
func (m *MockuserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByMobile", ctx, mobile)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByMobile indicates an expected call of FindOneByMobile.
func (mr *MockuserModelMockRecorder) FindOneByMobile(ctx, mobile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByMobile", reflect.TypeOf((*MockuserModel)(nil).FindOneByMobile), ctx, mobile)
}

// FindOneByName mocks base method.
func (m *MockuserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByName", ctx, name)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByName indicates an expected call of FindOneByName.
func (mr *MockuserModelMockRecorder) FindOneByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByName", reflect.TypeOf((*MockuserModel)(nil).FindOneByName), ctx, name)
}

// Insert mocks base method.
func (m *MockuserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockuserModelMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockuserModel)(nil).Insert), ctx, data)
}

// Update mocks base method.
func (m *MockuserModel) Update(ctx context.Context, data *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockuserModelMockRecorder) Update(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockuserModel)(nil).Update), ctx, data)
}