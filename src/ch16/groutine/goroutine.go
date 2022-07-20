package main

import (
	"fmt"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("正在唱歌")
		//time.Sleep(10 * time.Millisecond)
	}
}

func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("正在跳舞")
		//time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go sing()
	go dance()
	for {

	}
}
