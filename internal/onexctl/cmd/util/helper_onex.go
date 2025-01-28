// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package util

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/olekukonko/tablewriter"

	jwtauthn "github.com/onexstack/onexstack/pkg/authn/jwt"
)

func TableWriterDefaultConfig(table *tablewriter.Table) *tablewriter.Table {
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("  ") // pad with two space
	table.SetNoWhiteSpace(true)

	return table
}

// SignToken is ued to sign a jwt token with some default options.
func SignToken(secretID string, secretKey string) (string, error) {
	opts := []jwtauthn.Option{
		jwtauthn.WithSigningMethod(jwt.GetSigningMethod("HS256")),
		jwtauthn.WithIssuer("onexctl"),
		jwtauthn.WithTokenHeader(map[string]any{"kid": secretID}),
		jwtauthn.WithExpired(2 * time.Hour),
		jwtauthn.WithSigningKey([]byte(secretKey)),
	}
	j, err := jwtauthn.New(nil, opts...).Sign(context.Background(), "")
	if err != nil {
		return "", err
	}

	return j.GetToken(), nil
}
