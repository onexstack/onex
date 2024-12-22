package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/superproj/onex/pkg/xerrors"
)

// ErrorResponse 定义了一个错误响应结构体.
type ErrorResponse struct {
	// 理由
	Reason string `json:"reason,omitempty"`
	// 信息
	Message string `json:"message,omitempty"`
	// 元数据
	Metadata map[string]string `json:"metadata,omitempty"`
}

// WriteResponse 处理响应的函数.
func WriteResponse(c *gin.Context, err error, result any) {
	// 判断错误是否存在
	if err != nil {
		xerr := xerrors.FromError(err) // 从错误中获取详细信息
		// 返回错误响应
		c.JSON(xerr.Code, &ErrorResponse{
			Reason:   xerr.Reason,
			Message:  xerr.Message,
			Metadata: xerr.Metadata,
		})
		return
	}

	// 返回正常响应
	c.JSON(http.StatusOK, result)
}
