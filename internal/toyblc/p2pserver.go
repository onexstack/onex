// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package toyblc

import (
	"context"

	"github.com/gin-gonic/gin"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	"golang.org/x/net/websocket"

	"github.com/onexstack/onex/internal/toyblc/handler"
)

// p2pServer 定义一个使用 Gin 框架开发的 HTTP 服务器.
type p2pServer struct {
	srv server.Server
}

// 确保 *p2pServer 实现了 server.Server 接口.
var _ server.Server = (*p2pServer)(nil)

func (c *ServerConfig) NewP2PServer() *p2pServer {
	// 创建 Gin 引擎
	engine := gin.New()

	handler := handler.NewHandler(c.bs, c.ss)
	engine.Use(gin.WrapH(websocket.Handler(handler.WSHandler)))

	httpOptions := &genericoptions.HTTPOptions{ Addr: c.p2paddr, }
	httpsrv := server.NewHTTPServer(httpOptions, nil, engine)

	return &p2pServer{srv: httpsrv}
}

// RunOrDie 启动 Gin 服务器，出错则程序崩溃退出.
func (s *p2pServer) RunOrDie() {
	s.srv.RunOrDie()
}

// GracefulStop 优雅停止服务器.
func (s *p2pServer) GracefulStop(ctx context.Context) {
	s.srv.GracefulStop(ctx)
}
