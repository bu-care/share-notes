package main

import (
	"fmt"
	"sync"
)

func main() {
	//i共享内存，但是并发goroutine执行时i的值不确定是多少，所以最后total的值也不确定，
	//想要total的值和sum一样，需要使用通道channel
	total, sum := 0, 0

	// for i := 1; i <= 10; i++ {
	// 	sum += i
	// 	go func() {
	// 		total += i
	// 	}()
	// }
	// fmt.Printf("total:%d sum %d", total, sum)

	var wg sync.WaitGroup
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		sum += i
		wg.Add(1)
		//如果上面初始化ch的时候没有设置容量，就不能先向通道中存数据，不然会死锁- deadlock!，
		//此时可以将ch <- i写在goroutine后面，让goroutine先去等待读取通道数据
		ch <- i
		go func() {
			defer wg.Done()
			total += <-ch
		}()
		// ch <- i
	}
	wg.Wait()
	fmt.Printf("total:%d sum %d", total, sum)

}
