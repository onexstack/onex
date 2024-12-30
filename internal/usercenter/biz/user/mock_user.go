// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onexstack/onex/internal/usercenter/biz/user (interfaces: UserBiz)

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// MockUserBiz is a mock of UserBiz interface.
type MockUserBiz struct {
	ctrl     *gomock.Controller
	recorder *MockUserBizMockRecorder
}

// MockUserBizMockRecorder is the mock recorder for MockUserBiz.
type MockUserBizMockRecorder struct {
	mock *MockUserBiz
}

// NewMockUserBiz creates a new mock instance.
func NewMockUserBiz(ctrl *gomock.Controller) *MockUserBiz {
	mock := &MockUserBiz{ctrl: ctrl}
	mock.recorder = &MockUserBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserBiz) EXPECT() *MockUserBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserBiz) Create(arg0 context.Context, arg1 *v1.CreateUserRequest) (*v1.UserReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*v1.UserReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserBizMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserBiz)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUserBiz) Delete(arg0 context.Context, arg1 *v1.DeleteUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserBizMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserBiz)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockUserBiz) Get(arg0 context.Context, arg1 *v1.GetUserRequest) (*v1.UserReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.UserReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserBizMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserBiz)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockUserBiz) List(arg0 context.Context, arg1 *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserBizMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserBiz)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserBiz) Update(arg0 context.Context, arg1 *v1.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserBizMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserBiz)(nil).Update), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockUserBiz) UpdatePassword(arg0 context.Context, arg1 *v1.UpdatePasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserBizMockRecorder) UpdatePassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserBiz)(nil).UpdatePassword), arg0, arg1)
}
