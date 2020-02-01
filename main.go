package main

import (
	"fmt"
)

func main() {
	fn := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}
