// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package rid

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"os"
)

// Salt 计算机器 ID 的哈希值并返回一个 uint64 类型的盐值.
func Salt() uint64 {
	// 使用 FNV-1a 哈希算法计算字符串的哈希值
	hasher := fnv.New64a()
	hasher.Write(ReadMachineID())

	// 将哈希值转换为 uint64 型的盐
	hashValue := hasher.Sum64()
	return hashValue
}

// ReadMachineID 获取机器 ID，如果无法获取，则生成随机 ID.
func ReadMachineID() []byte {
	id := make([]byte, 3)
	machineID, err := readPlatformMachineID()
	if err != nil || len(machineID) == 0 {
		machineID, err = os.Hostname()
	}

	if err == nil && len(machineID) != 0 {
		hasher := sha256.New()
		hasher.Write([]byte(machineID))
		copy(id, hasher.Sum(nil))
	} else {
		// 如果无法收集机器 ID，则回退到生成随机数
		if _, randErr := rand.Reader.Read(id); randErr != nil {
			panic(fmt.Errorf("id: cannot get hostname nor generate a random number: %w; %w", err, randErr))
		}
	}
	return id
}

// readPlatformMachineID 尝试读取平台特定的机器 ID.
func readPlatformMachineID() (string, error) {
	data, err := os.ReadFile("/etc/machine-id")
	if err != nil || len(data) == 0 {
		data, err = os.ReadFile("/sys/class/dmi/id/product_uuid")
	}
	return string(data), err
}
