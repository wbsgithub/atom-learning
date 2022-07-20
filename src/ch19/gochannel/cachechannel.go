package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("len(ch):", len(ch), "cap(ch):", cap(ch))
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子go程：写入", i, "len(ch):", len(ch), "cap(ch):", cap(ch))
		}
	}()
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("主go程：读取", num)
	}

}
