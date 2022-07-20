package main

import (
	"fmt"
	"time"
)

var printchan = make(chan int)

func Printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func person1() {
	Printer("hello")
	printchan <- 8
}

func person2() {
	a := <-printchan
	fmt.Println(a)
	Printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
