// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/go-containerregistry/pkg/name (interfaces: Reference)

// Package testmock is a generated GoMock package.
package testmock

import (
	gomock "github.com/golang/mock/gomock"
	name "github.com/google/go-containerregistry/pkg/name"
	reflect "reflect"
)

// MockReference is a mock of Reference interface
type MockReference struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceMockRecorder
}

// MockReferenceMockRecorder is the mock recorder for MockReference
type MockReferenceMockRecorder struct {
	mock *MockReference
}

// NewMockReference creates a new mock instance
func NewMockReference(ctrl *gomock.Controller) *MockReference {
	mock := &MockReference{ctrl: ctrl}
	mock.recorder = &MockReferenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReference) EXPECT() *MockReferenceMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockReference) Context() name.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(name.Repository)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockReferenceMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockReference)(nil).Context))
}

// Identifier mocks base method
func (m *MockReference) Identifier() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockReferenceMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockReference)(nil).Identifier))
}

// Name mocks base method
func (m *MockReference) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockReferenceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockReference)(nil).Name))
}

// Scope mocks base method
func (m *MockReference) Scope(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scope", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Scope indicates an expected call of Scope
func (mr *MockReferenceMockRecorder) Scope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scope", reflect.TypeOf((*MockReference)(nil).Scope), arg0)
}

// String mocks base method
func (m *MockReference) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockReferenceMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockReference)(nil).String))
}
