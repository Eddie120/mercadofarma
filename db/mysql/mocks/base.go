// Code generated by MockGen. DO NOT EDIT.
// Source: base.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDataAccess is a mock of DataAccess interface.
type MockDataAccess struct {
	ctrl     *gomock.Controller
	recorder *MockDataAccessMockRecorder
}

// MockDataAccessMockRecorder is the mock recorder for MockDataAccess.
type MockDataAccessMockRecorder struct {
	mock *MockDataAccess
}

// NewMockDataAccess creates a new mock instance.
func NewMockDataAccess(ctrl *gomock.Controller) *MockDataAccess {
	mock := &MockDataAccess{ctrl: ctrl}
	mock.recorder = &MockDataAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataAccess) EXPECT() *MockDataAccessMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockDataAccess) BeginTx(ctx context.Context) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDataAccessMockRecorder) BeginTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDataAccess)(nil).BeginTx), ctx)
}

// Close mocks base method.
func (m *MockDataAccess) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDataAccessMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDataAccess)(nil).Close))
}

// ExecWithContext mocks base method.
func (m *MockDataAccess) ExecWithContext(ctx context.Context, query string, arg ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecWithContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecWithContext indicates an expected call of ExecWithContext.
func (mr *MockDataAccessMockRecorder) ExecWithContext(ctx, query interface{}, arg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecWithContext", reflect.TypeOf((*MockDataAccess)(nil).ExecWithContext), varargs...)
}

// Ping mocks base method.
func (m *MockDataAccess) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDataAccessMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDataAccess)(nil).Ping))
}

// QueryWithContext mocks base method.
func (m *MockDataAccess) QueryWithContext(ctx context.Context, query string, arg ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryWithContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWithContext indicates an expected call of QueryWithContext.
func (mr *MockDataAccessMockRecorder) QueryWithContext(ctx, query interface{}, arg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWithContext", reflect.TypeOf((*MockDataAccess)(nil).QueryWithContext), varargs...)
}
