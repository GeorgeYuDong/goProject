package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("go routine terminated")
		for i := 0; i < 5; i++ {
			c <- (i + 1) * 10
			fmt.Println("go routine send", (i+1)*10)
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("received data is", num)
	}

	fmt.Println("main routine terminated")
}
