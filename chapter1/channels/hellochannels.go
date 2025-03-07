package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	// 從channel 取得資料
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

// main is the entry point for the program.
func main() {
	//建立一个int通道
	c := make(chan int)
	// 把channel 传入, 让它开始等待资料喂入
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
