package main

import (
	"fmt"
	"sync"
	"time"
)

// channel to limit concurrency count
func worker(id int, wg *sync.WaitGroup, limiter chan struct{}) {
	defer wg.Done()

	// 向 limiter 发送一个值，表示占用一个并发槽
	limiter <- struct{}{}

	// 模拟工作
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// 从 limiter 接收一个值，表示释放一个并发槽
	<-limiter
}

func run() {
	// 限制并发数为 3
	limiter := make(chan struct{}, 3)

	// 使用 WaitGroup 等待所有 Goroutine 完成
	var wg sync.WaitGroup

	// 启动 10 个 worker
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg, limiter)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()
}
