// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package peer

import (
	"github.com/gin-gonic/gin"

	"github.com/onexstack/onex/pkg/core"
	"github.com/onexstack/onex/internal/toyblc/ws"
	v1 "github.com/onexstack/onex/pkg/api/toyblc/v1"
)

func (b *PeerController) Create(c *gin.Context) {
	var r v1.CreatePeerRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.Respond(c, err, nil)
		return
	}

	ws.ConnectToPeers(c, b.bs, b.ss, []string{r.Peer})

	core.Respond(c, nil, nil)
}
