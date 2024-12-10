// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/miniblog.

package xerrors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/status"
)

// Errno 定义了 miniblog 使用的错误类型.
type Errno struct {
	Code      Code
	Message   string
	RequestID string
}

func New(code Code, format string, args ...any) *Errno {
	return &Errno{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error 实现 error 接口中的 `Error` 方法.
func (err *Errno) Error() string {
	message := fmt.Sprintf("[%s]%s", err.Code.String(), err.Message)
	if err.RequestID != "" {
		message = fmt.Sprintf("[%s]%s", err.RequestID, message)
	}
	return message
}

// M 设置 Errno 类型错误中的 Message 字段.
// 注意这里是直接替换 err.Message 的内容.
func (err *Errno) M(format string, args ...any) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// This member function is used by grpc when converting an error into a status
func (err Errno) GRPCStatus() *status.Status {
	return status.New(err.Code.GRPC(), err.Error())
}

func (err Errno) HTTPStatus() int {
	return err.Code.HTTP()
}

func (err *Errno) Is(target error) bool {
	return errors.Is(err, target)
}

func (err *Errno) SetRequestID(requestID string) *Errno {
	err.RequestID = requestID
	return err
}

func Parse(err error) *Errno {
	if err == nil {
		return &Errno{Code: OK, Message: ""}
	}

	// 尝试进行类型断言，提取 *Errno 错误
	if errno, ok := err.(*Errno); ok {
		return errno
	}

	return &Errno{Code: Unknown, Message: err.Error()}
}
