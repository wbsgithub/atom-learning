package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 3)
	fmt.Println("len(ch):", len(ch), "cap(ch):", cap(ch))
	go func() {
		for i := 0; i < 10; i++ {
			ch <- strconv.Itoa(i)
		}
		close(ch)
	}()
	for {
		if num, ok := <-ch; ok {
			fmt.Println("主go程：读取", num)
		} else {
			n := <-ch
			if n == "" {
				fmt.Printf("关闭后：%s\n", n)
			}
			break
		}
	}

	for {
	}
}
