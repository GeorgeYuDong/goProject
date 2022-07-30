package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Printf("go routine terminated\n")
		for i := 0; i < 6; i++ {
			c <- i * 100
			fmt.Println("go routine, send element", i*100)
		}
	}()

	time.Sleep(2 * time.Second)

	var i int
	for i = 0; i < 6; i++ {
		num := <-c
		fmt.Println("num =", num)
	}

	fmt.Println("main routine terminated")
}
