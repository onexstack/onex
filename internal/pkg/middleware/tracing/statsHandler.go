// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
)

// ClientHandler is tracing ClientHandler.
type ClientHandler struct{}

// HandleConn exists to satisfy gRPC stats.Handler.
func (c *ClientHandler) HandleConn(ctx context.Context, cs stats.ConnStats) {
	fmt.Println("Handle connection.")
}

// TagConn exists to satisfy gRPC stats.Handler.
func (c *ClientHandler) TagConn(ctx context.Context, cti *stats.ConnTagInfo) context.Context {
	return ctx
}

// HandleRPC implements per-RPC tracing and stats instrumentation.
func (c *ClientHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	if _, ok := rs.(*stats.OutHeader); !ok {
		return
	}
	p, ok := peer.FromContext(ctx)
	if !ok {
		return
	}
	remoteAddr := p.Addr.String()
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
		span.SetAttributes(peerAttr(remoteAddr)...)
	}
}

// TagRPC implements per-RPC context management.
func (c *ClientHandler) TagRPC(ctx context.Context, rti *stats.RPCTagInfo) context.Context {
	return ctx
}
