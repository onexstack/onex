package core

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/onex/pkg/errorsx"
)

// ValidatorFn 定义验证函数类型。用于对绑定的数据结构进行验证。  
type ValidatorFn[T any] func(context.Context, T) error

// BindFn 定义绑定函数类型，接收通用类型参数并返回错误。  
type BindFn func(any) error

// ErrorResponse 定义错误响应结构体。  
// 用于 API 请求中发生错误时，按照统一格式返回给客户端。  
type ErrorResponse struct {
	Reason   string            `json:"reason,omitempty"`   // 错误原因，标识错误类型  
	Message  string            `json:"message,omitempty"`  // 错误的具体描述信息  
	Metadata map[string]string `json:"metadata,omitempty"` // 元数据，包含额外的上下文信息  
}

// ShouldBindJSON 使用 JSON 格式的绑定函数绑定请求参数并执行验证。  
func ShouldBindJSON[T any](c *gin.Context, rq T, validators ...ValidatorFn[T]) error {
	return ReadRequest(c, rq, c.ShouldBindJSON, validators...)
}

// ShouldBindQuery 使用 Query 格式的绑定函数绑定请求参数并执行验证。
func ShouldBindQuery[T any](c *gin.Context, rq T, validators ...ValidatorFn[T]) error {
	return ReadRequest(c, rq, c.ShouldBindQuery, validators...)
}

// ShouldBindUri 使用 URI 格式的绑定函数绑定请求参数并执行验证。
func ShouldBindUri[T any](c *gin.Context, rq T, validators ...ValidatorFn[T]) error {
	return ReadRequest(c, rq, c.ShouldBindUri, validators...)
}

// ReadRequest 是通用的请求绑定和验证工具函数。
// 它会对请求进行参数绑定，初始化默认值（如果目标结构体实现了 Default 接口），并执行验证函数（可选）。
func ReadRequest[T any](c *gin.Context, rq T, bindFn BindFn, validators ...ValidatorFn[T]) error {
	// 调用绑定函数绑定请求数据
	if err := bindFn(rq); err != nil {
		return errorsx.ErrBind.WithMessage(err.Error())
	}

	// 如果目标结构体实现了 Default 接口，则调用其 Default 方法设置默认值
	if defaulter, ok := any(rq).(interface{ Default() }); ok {
		defaulter.Default()
	}

	// 遍历所有验证函数并执行它们
	for _, validator := range validators {
		if validator == nil { // 跳过 nil 的验证函数  
			continue
		}
		if err := validator(c.Request.Context(), rq); err != nil {
			return err
		}
	}

	return nil
}

// WriteResponse 是统一的响应处理函数。  
// 根据返回值是否发生错误，返回成功响应或错误响应。  
func WriteResponse(c *gin.Context, result any, err error) {
	// 如果发生错误，生成错误响应  
	if err != nil {
		errx := errorsx.FromError(err) // 从错误对象中解析详细错误信息  
		c.JSON(errx.Code, &ErrorResponse{
			Reason:   errx.Reason,
			Message:  errx.Message,
			Metadata: errx.Metadata,
		})
		return
	}

	// 如果没有错误，生成成功响应  
	c.JSON(http.StatusOK, result)
}
