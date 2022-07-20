package main

import (
	"fmt"
	"testing"
	"time"
)

func Hello() {
	fmt.Println("hello groutine")
}
func TestGroutine(t *testing.T) {
	go Hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main test func")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func TestGroutineLoop(t *testing.T) {
	go numbers()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

func a(i int) {
	fmt.Println(i)
}

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go a(i)
	}
	time.Sleep(1000 * time.Millisecond)
}
