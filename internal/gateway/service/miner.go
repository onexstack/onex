// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package service

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/onexstack/onex/internal/pkg/contextx"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
)

func (s *GatewayService) CreateMiner(ctx context.Context, m *v1beta1.Miner) (*emptypb.Empty, error) {
	if err := s.biz.Miners().Create(ctx, contextx.FromUserID(ctx), m); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *GatewayService) ListMiner(ctx context.Context, rq *v1.ListMinerRequest) (*v1.ListMinerResponse, error) {
	ms, err := s.biz.Miners().List(ctx, contextx.FromUserID(ctx), rq)
	if err != nil {
		return &v1.ListMinerResponse{}, err
	}

	return ms, nil
}

func (s *GatewayService) GetMiner(ctx context.Context, rq *v1.GetMinerRequest) (*v1beta1.Miner, error) {
	m, err := s.biz.Miners().Get(ctx, contextx.FromUserID(ctx), rq.Name)
	if err != nil {
		return &v1beta1.Miner{}, err
	}

	return m, nil
}

func (s *GatewayService) UpdateMiner(ctx context.Context, m *v1beta1.Miner) (*emptypb.Empty, error) {
	if err := s.biz.Miners().Update(ctx, contextx.FromUserID(ctx), m); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *GatewayService) DeleteMiner(ctx context.Context, rq *v1.DeleteMinerRequest) (*emptypb.Empty, error) {
	if err := s.biz.Miners().Delete(ctx, contextx.FromUserID(ctx), rq.Name); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
