package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		fmt.Printf("new Goroutine : i = %d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go newTask()

	//main 退出， goroutine也退出
	/*
		i := 0
		for {
			fmt.Printf("main Goroutine : i = %d\n", i)
			i++
			time.Sleep(1 * time.Second)
		}
	*/
	go func() {
		defer println("china")
		fmt.Println("b")
		runtime.Goexit() //runtime.Goexit() 退出
		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println("a is", a, ", b is", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
