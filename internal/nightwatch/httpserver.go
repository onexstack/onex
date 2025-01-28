// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package nightwatch

import (
	"context"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/core"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	"gorm.io/gorm"

	mw "github.com/onexstack/onex/internal/pkg/middleware/gin"
	"github.com/onexstack/onex/pkg/api/errno"
)

// RESTServer 定义一个使用 Gin 框架开发的 HTTP 服务器.
type RESTServer struct {
	srv server.Server
}

// 确保 *RESTServer 实现了 server.Server 接口.
var _ server.Server = (*RESTServer)(nil)

func NewRESTServer(httpOptions *genericoptions.HTTPOptions, tlsOptions *genericoptions.TLSOptions, db *gorm.DB) *RESTServer {
	// 创建 Gin 引擎
	engine := gin.New()

	// 注册全局中间件，用于恢复 panic、设置 HTTP 头、添加请求 ID 等
	engine.Use(gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.TraceID())

	// 注册业务无关的 API 接口
	InstallGenericAPI(engine)

	// 注册 REST API 路由
	InstallJobAPI(engine, db)

	httpsrv := server.NewHTTPServer(httpOptions, tlsOptions, engine)

	return &RESTServer{srv: httpsrv}
}

func InstallJobAPI(engine *gin.Engine, db *gorm.DB) {
	handler := InitializeHandler(db)

	v1 := engine.Group("/v1")
	{
		// 用户相关路由
		cronjobv1 := v1.Group("/cronjobs")
		{
			cronjobv1.POST("", handler.CreateCronJob)
			cronjobv1.PUT(":cronJobID", handler.UpdateCronJob)
			cronjobv1.DELETE("", handler.DeleteCronJob)
			cronjobv1.GET(":cronJobID", handler.GetCronJob)
			cronjobv1.GET("", handler.ListCronJob)
		}

		// 博客相关路由
		jobv1 := v1.Group("/jobs")
		{
			jobv1.POST("", handler.CreateJob)
			jobv1.PUT(":jobID", handler.UpdateJob)
			jobv1.DELETE("", handler.DeleteJob)
			jobv1.GET(":jobID", handler.GetJob)
			jobv1.GET("", handler.ListJob)
		}
	}
}

// InstallGenericAPI 注册业务无关的路由，例如 pprof、404 处理等.
func InstallGenericAPI(engine *gin.Engine) {
	// 注册 pprof 路由
	pprof.Register(engine)

	// 注册 404 路由处理
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrorPageNotFound("Page Not Found"), nil)
	})
}

// RunOrDie 启动 Gin 服务器，出错则程序崩溃退出.
func (s *RESTServer) RunOrDie() {
	s.srv.RunOrDie()
}

// GracefulStop 优雅停止服务器.
func (s *RESTServer) GracefulStop(ctx context.Context) {
	s.srv.GracefulStop(ctx)
}
