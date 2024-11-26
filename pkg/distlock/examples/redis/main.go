package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/superproj/onex/pkg/db"
	"github.com/superproj/onex/pkg/distlock"
)

func main() {
	opts := &db.RedisOptions{
		Addr: "10.37.91.93:6379",
		//Username: "onex",
		Password: "onex(#)666",
		Database: 0,
	}

	db, err := db.NewRedis(opts)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	locker := distlock.NewRedisLocker(db)

	ctx := context.Background()

	// 持续尝试获取锁，直到成功
	for {
		if err := locker.Lock(ctx); err != nil {
			fmt.Printf("failed to acquire lock: %v, retrying...\n", err)
			// 等待一段时间后重试
			time.Sleep(1 * time.Second) // 可以根据需要调整等待时间
			continue                    // 继续尝试获取锁
		}
		fmt.Println("Lock acquired!")
		break // 成功获取锁后退出循环
	}

	// 模拟业务逻辑
	time.Sleep(1000 * time.Second) // 修改为合理的时间，避免长时间阻塞  

	if err := locker.Unlock(ctx); err != nil {
		fmt.Printf("failed to release lock: %v\n", err)
	} else {
		fmt.Println("Lock released!")
	}
}
