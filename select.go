package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		ch1 <- "foo"
	}()

	go func() {
		time.Sleep(time.Second * 20)
		ch2 <- "bar"
	}()

	//	for i := 0; i < 3; i++ {
	//	time.Sleep(time.Second * time.Duration(i))
	select {
	case msg1 := <-ch1:
		fmt.Println("received", msg1)
	case msg2 := <-ch2:
		fmt.Println("received", msg2)
		//default:
		//	fmt.Println("All channels are empty...")
	}
	//	}
}
