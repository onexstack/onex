package xerrors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"github.com/superproj/onex/pkg/xerrors"
)

func TestXError_NewAndToString(t *testing.T) {
	// 创建一个 XError 错误  
	xerr := xerrors.New(500, "InternalError.DBConnection", "Database connection failed: %s", "timeout")

	// 检查字段值  
	assert.Equal(t, 500, xerr.Code)
	assert.Equal(t, "InternalError.DBConnection", xerr.Reason)
	assert.Equal(t, "Database connection failed: timeout", xerr.Message)

	// 检查字符串表示  
	expected := `error: code = 500 reason = InternalError.DBConnection message = Database connection failed: timeout metadata = map[]`
	assert.Equal(t, expected, xerr.Error())
}

func TestXError_WithMessage(t *testing.T) {
	// 创建一个基础错误  
	xerr := xerrors.New(400, "BadRequest.InvalidInput", "Invalid input for field %s", "username")

	// 更新错误的消息  
	xerr.WithMessage("New error message: %s", "retry failed")

	// 验证变更  
	assert.Equal(t, "New error message: retry failed", xerr.Message)
	assert.Equal(t, 400, xerr.Code)                         // Code 不变
	assert.Equal(t, "BadRequest.InvalidInput", xerr.Reason) // Reason 不变
}

func TestXError_WithMetadata(t *testing.T) {
	// 创建基础错误
	xerr := xerrors.New(400, "BadRequest.InvalidInput", "Invalid input")

	// 添加元数据
	xerr.WithMetadata(map[string]string{
		"field": "username",
		"type":  "empty",
	})

	// 验证元数据
	assert.Equal(t, "username", xerr.Metadata["field"])
	assert.Equal(t, "empty", xerr.Metadata["type"])

	// 动态添加更多元数据
	xerr.KV("user_id", "12345", "trace_id", "xyz-789")
	assert.Equal(t, "12345", xerr.Metadata["user_id"])
	assert.Equal(t, "xyz-789", xerr.Metadata["trace_id"])
}

func TestXError_Is(t *testing.T) {
	// 定义两个预定义错误
	err1 := xerrors.New(404, "NotFound.User", "User not found")
	err2 := xerrors.New(404, "NotFound.User", "Another message")
	err3 := xerrors.New(403, "Forbidden", "Access denied")

	// 验证两个错误均被认为是同一种类型的错误（Code 和 Reason 相等）
	assert.True(t, err1.Is(err2))  // Message 不影响匹配
	assert.False(t, err1.Is(err3)) // Reason 不同
}

func TestXError_FromError_WithPlainError(t *testing.T) {
	// 创建一个普通的 Go 错误
	plainErr := errors.New("Something went wrong")

	// 转换为 XError
	xerr := xerrors.FromError(plainErr)

	// 检查转换后的 XError
	assert.Equal(t, xerrors.UnknownCode, xerr.Code)       // 默认 500
	assert.Equal(t, xerrors.UnknownReason, xerr.Reason)   // 默认 ""
	assert.Equal(t, "Something went wrong", xerr.Message) // 转换时保留原始错误消息
}

func TestXError_FromError_WithGRPCError(t *testing.T) {
	// 创建一个 gRPC 错误
	grpcErr := status.New(3, "Invalid argument").Err() // gRPC INVALID_ARGUMENT = 3

	// 转换为 XError  
	xerr := xerrors.FromError(grpcErr)

	// 检查转换后的 XError  
	assert.Equal(t, 400, xerr.Code) // httpstatus.FromGRPCCode(3) 对应 HTTP 400  
	assert.Equal(t, "Invalid argument", xerr.Message)

	// 没有附加的元数据  
	assert.Nil(t, xerr.Metadata)
}

func TestXError_FromError_WithGRPCErrorDetails(t *testing.T) {
	// 创建带有详细信息的 gRPC 错误  
	st := status.New(3, "Invalid argument")
	grpcErr, err := st.WithDetails(&errdetails.ErrorInfo{
		Reason:   "InvalidInput",
		Metadata: map[string]string{"field": "name", "type": "required"},
	})
	assert.NoError(t, err) // 确保 gRPC 错误创建成功  

	// 转换为 XError  
	xerr := xerrors.FromError(grpcErr.Err())

	// 检查转换后的 XError  
	assert.Equal(t, 400, xerr.Code) // gRPC INVALID_ARGUMENT = HTTP 400
	assert.Equal(t, "Invalid argument", xerr.Message)
	assert.Equal(t, "InvalidInput", xerr.Reason) // 从 gRPC ErrorInfo 中提取  

	// 检查元数据  
	assert.Equal(t, "name", xerr.Metadata["field"])
	assert.Equal(t, "required", xerr.Metadata["type"])
}
