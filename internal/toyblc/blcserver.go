// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package toyblc

import (
	"context"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/core"
	"github.com/onexstack/onexstack/pkg/server"

	"github.com/onexstack/onex/internal/toyblc/handler"
	mw "github.com/onexstack/onex/internal/toyblc/pkg/middleware"
	"github.com/onexstack/onex/pkg/api/errno"
)

// blcServer 定义一个使用 Gin 框架开发的 HTTP 服务器.
type blcServer struct {
	srv server.Server
}

// 确保 *blcServer 实现了 server.Server 接口.
var _ server.Server = (*blcServer)(nil)

func (c *ServerConfig) NewBlockChainServer() *blcServer {
	// 创建 Gin 引擎
	engine := gin.New()

	// 注册全局中间件，用于恢复 panic、设置 HTTP 头、添加请求 ID 等
	engine.Use(gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.TraceID())

	// 注册 REST API 路由
	c.InstallRESTAPI(engine)

	httpsrv := server.NewHTTPServer(c.cfg.HTTPOptions, c.cfg.TLSOptions, engine)

	return &blcServer{srv: httpsrv}
}

// 注册 API 路由。路由的路径和 HTTP 方法，严格遵循 REST 规范.
func (c *ServerConfig) InstallRESTAPI(engine *gin.Engine) {
	// 注册业务无关的 API 接口
	InstallGenericAPI(engine)

	// 创建核心业务处理器
	handler := handler.NewHandler(c.bs, c.ss)

	// 注册 v1 版本 API 路由分组
	v1 := engine.Group("/v1", mw.BasicAuth(c.accounts))
	{
		// 创建 blocks 路由分组
		blockv1 := v1.Group("/blocks")
		{
			blockv1.POST("", handler.CreateBlock)
			blockv1.GET("", handler.ListBlock)
		}

		// 创建 peers 路由分组
		peerv1 := v1.Group("/peers")
		{
			peerv1.POST("", handler.CreatePeer)
			peerv1.GET("", handler.ListPeer)
		}
	}
}

// InstallGenericAPI 注册业务无关的路由，例如 pprof、404 处理等.
func InstallGenericAPI(engine *gin.Engine) {
	// 注册 pprof 路由
	pprof.Register(engine)

	// 注册 404 路由处理
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, nil, errno.ErrorPageNotFound("Page Not Found"))
	})
}

// RunOrDie 启动 Gin 服务器，出错则程序崩溃退出.
func (s *blcServer) RunOrDie() {
	s.srv.RunOrDie()
}

// GracefulStop 优雅停止服务器.
func (s *blcServer) GracefulStop(ctx context.Context) {
	s.srv.GracefulStop(ctx)
}
