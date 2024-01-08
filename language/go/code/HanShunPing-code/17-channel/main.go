package main

import "fmt"

func send(ch chan<- int, exite_chan chan struct{}) {
	for i := 0; i < 10; i++ {
		fmt.Println("send: ", i)
		ch <- i
	}
	close(ch)
	var a struct{}
	exite_chan <- a
}

func recv(ch <-chan int, exite_chan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("recv: ", v)
	}
	var a struct{}
	exite_chan <- a
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	exite_chan := make(chan struct{}, 2)
	go send(ch, exite_chan)
	go recv(ch, exite_chan)

	var total = 0
	for range exite_chan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("end!!!")
}
