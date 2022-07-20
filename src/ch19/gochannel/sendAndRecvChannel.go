package main

import "fmt"

/**
单向读通道：只能读数据，不能写
单向写通道：只能写数据，不能读
*/
func recv(in <-chan int) {
	num := <-in
	//in <- 100
	fmt.Println("读取：", num)
}

func send(out chan<- int) {
	out <- 100
	//num := <-out
	close(out)
}

func main() {
	ch := make(chan int)
	go func() {
		send(ch)
	}()
	recv(ch)
}
