package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	// fatal error: all goroutines are asleep - deadlock!
	// 要 sender 跟 receiver 都 ready 才能送訊息
	// c <- "hello"
	// val := <-c
	// fmt.Println(val)

	go func() {
		c <- "hello"
	}()

	val := <-c
	fmt.Println(val)
}
