package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go test(table)
	go player("ping", table)
	go player("pong", table)

	//table <- new(Ball) // game on; toss the ball
	//fmt.Println("main", 999)
	time.Sleep(1 * time.Second)
	x := <-table // game over; grab the ball
	fmt.Printf("x is %v \n", x.hits)

	//panic("show me the stacks")
}

func test(table chan *Ball) {

	//fmt.Printf("waiting for ...")
	tmp := new(Ball)
	tmp.hits = 88
	table <- tmp
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
