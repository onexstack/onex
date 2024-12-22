package xerrors

import (
	"errors"
	"fmt"
	"net/http"

	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

const (
	// UnknownCode 定义了未知错误的代码.
	UnknownCode = http.StatusInternalServerError
	// UnknownReason 定义了未知错误的原因.
	UnknownReason = ""
)

// XError 定义了 OneX 项目体系中使用的错误类型，用于描述错误的详细信息。
type XError struct {
	// Code 表示错误的 HTTP 状态码，用于与客户端进行交互时标识错误的类型。
	Code int `json:"code,omitempty"`

	// Reason 表示错误发生的原因，通常为业务错误码，用于精准定位问题。
	Reason string `json:"reason,omitempty"`

	// Message 表示简短的错误信息，通常可直接暴露给用户查看。
	Message string `json:"message,omitempty"`

	// Metadata 用于存储与该错误相关的额外元信息，可以包含上下文或调试信息。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// New 创建一个新的错误.
func New(code int, reason string, format string, args ...any) *XError {
	return &XError{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error 实现 error 接口中的 `Error` 方法.
func (err *XError) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v", err.Code, err.Reason, err.Message, err.Metadata)
}

// WithMessage 设置错误的 Message 字段.  
func (err *XError) WithMessage(format string, args ...any) *XError {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// WithMetadata 设置元数据.
func (err *XError) WithMetadata(md map[string]string) *XError {
	err.Metadata = md
	return err
}

// KV 使用 key-value 对设置元数据.
func (err *XError) KV(kvs ...string) *XError {
	if err.Metadata == nil {
		err.Metadata = make(map[string]string) // 初始化元数据映射
	}

	for i := 0; i < len(kvs); i += 2 {
		// kvs 必须是成对的
		if i+1 < len(kvs) {
			err.Metadata[kvs[i]] = kvs[i+1]
		}
	}
	return err
}

// GRPCStatus 返回 gRPC 状态表示.
func (err *XError) GRPCStatus() *status.Status {
	details := errdetails.ErrorInfo{Reason: err.Reason, Metadata: err.Metadata}
	s, _ := status.New(httpstatus.ToGRPCCode(err.Code), err.Message).WithDetails(&details)
	return s
}

// WithRequestID 设置请求 ID.  
func (err *XError) WithRequestID(requestID string) *XError {
	return err.KV("X-Request-ID", requestID) // 设置请求 ID  
}

// Is matches each error in the chain with the target value.
func (err *XError) Is(target error) bool {
	if xerr := new(XError); errors.As(target, &xerr) {
		return xerr.Code == err.Code && xerr.Reason == err.Reason
	}
	return false
}

// Code 返回错误的 HTTP 代码.
func Code(err error) int {
	if err == nil {
		return http.StatusOK //nolint:mnd
	}
	return FromError(err).Code
}

// Reason 返回特定错误的原因.
func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}

// FromError 尝试将一个通用的 error 转换为自定义的 *XError 类型.
func FromError(err error) *XError {
	// 如果传入的错误是 nil，则直接返回 nil，表示没有错误需要处理。
	if err == nil {
		return nil
	}

	// 检查传入的 error 是否已经是 XError 类型的实例.
	// 如果错误可以通过 errors.As 转换为 *XError 类型，则直接返回该实例.
	if xerr := new(XError); errors.As(err, &xerr) {
		return xerr
	}

	// gRPC 的 status.FromError 方法尝试将 error 转换为 gRPC 错误的 status 对象.
	// 如果 err 不能转换为 gRPC 错误（即不是 gRPC 的 status 错误），
	// 则返回一个带有默认值的 XError，表示是一个未知类型的错误.
	gs, ok := status.FromError(err)
	if !ok {
		return New(UnknownCode, UnknownReason, err.Error())
	}

	// 如果 err 是 gRPC 的错误类型，会成功返回一个 gRPC status 对象（gs）。
	// 使用 gRPC 状态中的错误代码和消息创建一个 XError。
	ret := New(httpstatus.FromGRPCCode(gs.Code()), UnknownReason, gs.Message())

	// 遍历 gRPC 错误详情中的所有附加信息（Details）。
	for _, detail := range gs.Details() {
		if typed, ok := detail.(*errdetails.ErrorInfo); ok {
			ret.Reason = typed.Reason
			return ret.WithMetadata(typed.Metadata)
		}
	}

	return ret
}
