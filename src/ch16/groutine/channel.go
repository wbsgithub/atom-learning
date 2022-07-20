package main

import "fmt"

func main() {
	ch := make(chan string) //无缓冲channel
	//len(ch) : channel中剩余未读取数据个数
	//cap(ch) : 通道中容量
	fmt.Println("len(ch):", len(ch), "cap(ch):", cap(ch))
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
			fmt.Println("len(ch):", len(ch), "cap(ch):", cap(ch))
		}
		ch <- "adsf"
	}()
	str := <-ch
	fmt.Println(str)

}
