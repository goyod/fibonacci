package main

import (
	"fmt"
)

func main() {
	chNum := make(chan int)
	chQuit := make(chan struct{})

	go fibonacci(chNum, chQuit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-chNum)
	}
	chQuit <- struct{}{}
}

func fibonacci(chNum chan int, chQuit chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case <-chQuit:
			return
		case chNum <- a:
			a, b = b, a+b
		}
	}
}
