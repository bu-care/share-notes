package main

import (
	"fmt"
	"sync"
)

// 发送给独立的goroutine处理程序, http://docs.lvrui.io/2020/03/26/go%E8%AF%AD%E8%A8%80%E5%9C%A8goroutine%E4%B8%AD%E6%8B%BF%E5%88%B0%E8%BF%94%E5%9B%9E%E5%80%BC
var responseChannel = make(chan string, 15)

func httpGet01(url int, limiter chan bool, wg *sync.WaitGroup) {
	// 函数执行完毕时 计数器-1
	defer wg.Done()
	fmt.Println("http get:", url)
	responseChannel <- fmt.Sprintf("Hello Go %d", url)
	// 释放一个坑位
	<-limiter
}

func ResponseController() {
	for rc := range responseChannel {
		fmt.Println("response: ", rc)
	}
}

func init() {

	// 启动接收response的控制器
	go ResponseController()

	wg := &sync.WaitGroup{}
	// 控制并发数为10
	limiter := make(chan bool, 20)

	for i := 0; i < 99; i++ {
		// 计数器+1
		wg.Add(1)
		limiter <- true
		go httpGet01(i, limiter, wg)
	}
	// 等待所以协程执行完毕
	wg.Wait() // 当计数器为0时, 不再阻塞
	fmt.Println("所有协程已执行完毕")
}
