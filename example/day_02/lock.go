package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
)

// 加锁连加操作
func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}

// 不加锁连加操作
func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x++
	}
}

// 五个协程连续操作
func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	fmt.Println("withLock:", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	fmt.Println("withoutLock:", x)
}

func main() {
	Add()
}
