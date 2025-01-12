// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package model

import (
	"gorm.io/gorm"

	"github.com/onexstack/onex/internal/pkg/zid"
)

// AfterCreate runs after creating a OrderM database record and updates the OrderID field.
func (m *OrderM) AfterCreate(tx *gorm.DB) (err error) {
	m.OrderID = zid.Order.New(uint64(m.ID)) // Generate and set a new order ID.

	return tx.Save(m).Error // Save the updated order record to the database.
}
