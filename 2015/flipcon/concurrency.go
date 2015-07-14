package main

import (
	"fmt"
	"time"
)

func f(msg string, count int, delay time.Duration, ch chan string) {
	for i := 0; i < count; i++ {
		ch <- msg
		time.Sleep(delay)
	}
}

func main() {
	ch := make(chan string)
	go f("flip", 10, 300*time.Millisecond, ch)
	go f("BOARD", 10, 500*time.Millisecond, ch)
	go f(" !!!", 10, 1100*time.Millisecond, ch)

	for i := 0; i < 30; i++ {
		fmt.Println(i, <-ch)
	}
}
