// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onexstack/onex/internal/cacheserver/biz/secret (interfaces: SecretBiz)

// Package secret is a generated GoMock package.
package secret

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/onexstack/onex/pkg/api/cacheserver/v1"
)

// MockSecretBiz is a mock of SecretBiz interface.
type MockSecretBiz struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBizMockRecorder
}

// MockSecretBizMockRecorder is the mock recorder for MockSecretBiz.
type MockSecretBizMockRecorder struct {
	mock *MockSecretBiz
}

// NewMockSecretBiz creates a new mock instance.
func NewMockSecretBiz(ctrl *gomock.Controller) *MockSecretBiz {
	mock := &MockSecretBiz{ctrl: ctrl}
	mock.recorder = &MockSecretBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBiz) EXPECT() *MockSecretBizMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockSecretBiz) Del(arg0 context.Context, arg1 *v1.DelSecretRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockSecretBizMockRecorder) Del(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockSecretBiz)(nil).Del), arg0, arg1)
}

// Get mocks base method.
func (m *MockSecretBiz) Get(arg0 context.Context, arg1 *v1.GetSecretRequest) (*v1.GetSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSecretBizMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretBiz)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockSecretBiz) Set(arg0 context.Context, arg1 *v1.SetSecretRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSecretBizMockRecorder) Set(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSecretBiz)(nil).Set), arg0, arg1)
}
