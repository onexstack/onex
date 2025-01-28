package core

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/onexstack/onexstack/pkg/errorsx"
)

// Validator 是验证函数的类型，用于对绑定的数据结构进行验证.  
type Validator[T any] func(context.Context, *T) error

// Binder 定义绑定函数的类型，用于绑定请求数据到相应结构体.  
type Binder func(any) error

// Handler 是处理函数的类型，用于处理已经绑定和验证的数据.  
type Handler[T any, R any] func(ctx context.Context, req *T) (R, error)

// ErrorResponse 定义了错误响应的结构，  
// 用于 API 请求中发生错误时返回统一的格式化错误信息.  
type ErrorResponse struct {
	// 错误原因，标识错误类型
	Reason string `json:"reason,omitempty"`
	// 错误详情的描述信息
	Message string `json:"message,omitempty"`
	// 附带的元数据信息
	Metadata map[string]string `json:"metadata,omitempty"`
}

// HandleJSONRequest 是处理 JSON 请求的快捷函数.
func HandleJSONRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindJSON, handler, validators...)
}

// HandleQueryRequest 是处理 Query 参数请求的快捷函数.
func HandleQueryRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindQuery, handler, validators...)
}

// HandleUriRequest 是处理 URI 请求的快捷函数.
func HandleUriRequest[T any, R any](c *gin.Context, handler Handler[T, R], validators ...Validator[T]) {
	HandleRequest(c, c.ShouldBindUri, handler, validators...)
}

// HandleRequest 是通用的请求处理函数.
// 负责绑定请求数据、执行验证、并调用实际的业务处理逻辑函数.
func HandleRequest[T any, R any](c *gin.Context, binder Binder, handler Handler[T, R], validators ...Validator[T]) {
	var request T

	// 绑定和验证请求数据
	if err := ReadRequest(c, &request, binder, validators...); err != nil {
		WriteResponse(c, err, nil)
		return
	}

	// 调用实际的业务逻辑处理函数
	response, err := handler(c.Request.Context(), &request)
	WriteResponse(c, err, response)
}

// ShouldBindJSON 使用 JSON 格式的绑定函数绑定请求参数并执行验证。
func ShouldBindJSON[T any](c *gin.Context, rq *T, validators ...Validator[T]) error {
	return ReadRequest(c, rq, c.ShouldBindJSON, validators...)
}

// ShouldBindQuery 使用 Query 格式的绑定函数绑定请求参数并执行验证。
func ShouldBindQuery[T any](c *gin.Context, rq *T, validators ...Validator[T]) error {
	return ReadRequest(c, rq, c.ShouldBindQuery, validators...)
}

// ShouldBindUri 使用 URI 格式的绑定函数绑定请求参数并执行验证。
func ShouldBindUri[T any](c *gin.Context, rq *T, validators ...Validator[T]) error {
	return ReadRequest(c, rq, c.ShouldBindUri, validators...)
}

// ReadRequest 是用于绑定和验证请求数据的通用工具函数.
// - 它负责调用绑定函数绑定请求数据.
// - 如果目标类型实现了 Default 接口，会调用其 Default 方法设置默认值.
// - 最后执行传入的验证器对数据进行校验.
func ReadRequest[T any](c *gin.Context, rq *T, binder Binder, validators ...Validator[T]) error {
	// 调用绑定函数绑定请求数据
	if err := binder(rq); err != nil {
		return errorsx.ErrBind.WithMessage(err.Error())
	}

	// 如果数据结构实现了 Default 接口，则调用它的 Default 方法
	if defaulter, ok := any(rq).(interface{ Default() }); ok {
		defaulter.Default()
	}

	// 执行所有验证函数
	for _, validate := range validators {
		if validate == nil { // 跳过 nil 的验证器
			continue
		}
		if err := validate(c.Request.Context(), rq); err != nil {
			return err
		}
	}

	return nil
}

// WriteResponse 是通用的响应函数.
// 它会根据是否发生错误，生成成功响应或标准化的错误响应.
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		// 如果发生错误，生成错误响应
		errx := errorsx.FromError(err) // 提取错误详细信息
		c.JSON(errx.Code, ErrorResponse{
			Reason:   errx.Reason,
			Message:  errx.Message,
			Metadata: errx.Metadata,
		})
		return
	}

	// 如果没有错误，返回成功响应  
	c.JSON(http.StatusOK, data)
}
