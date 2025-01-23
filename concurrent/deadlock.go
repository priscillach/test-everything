package main

import (
	"sync"
	"time"
)

func deadLock1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func() {
		defer wait.Done()
		time.Sleep(100 * time.Millisecond) // 增加延迟
		<-ch1                              // 等待 ch1
		ch2 <- 1
	}()
	go func() {
		defer wait.Done()
		time.Sleep(100 * time.Millisecond) // 增加延迟
		<-ch2                              // 等待 ch2
		ch1 <- 1
	}()
	wait.Wait()
}

func deadLock2() {
	wait := sync.WaitGroup{}
	wait.Add(2)
	var lock1, lock2 sync.Mutex
	go func() {
		defer wait.Done()
		lock1.Lock()
		time.Sleep(100 * time.Millisecond) // 增加延迟
		lock2.Lock()
		lock2.Unlock()
		lock1.Unlock()
	}()
	go func() {
		defer wait.Done()
		lock2.Lock()
		time.Sleep(100 * time.Millisecond) // 增加延迟
		lock1.Lock()
		lock1.Unlock()
		lock2.Unlock()
	}()
	wait.Wait()
}
