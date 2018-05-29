package main

import (
	"log"
	"sync"
)

func start(args ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, arg := range args {
			out <- arg * arg

		}
		close(out)
	}()
	//	log.Printf("the x in start is %v\n", <-out)

	return out
}

func read(arcs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 3)

	hander := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(arcs))
	for _, arc := range arcs {
		go hander(arc)
		//		log.Printf("the x in read is %v\n", <-out)

	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	out1 := start(3, 7, 8, 9)
	out3 := start(1)
	out2 := start(5, 4, 6)
	result := read(out3, out2, out1)

	switch interface{}(result).(type) {
	case chan int:
		log.Printf("1")
		// we can't use fallthrough in type switch
	case <-chan int:
		log.Printf("2 \n")
	default:
		log.Printf("3")
	}
	for i := range result {

		log.Printf("the x final %v\n", i)
	}
	//errorf return a error type, like Sprintf	log.Errorf("print type %v \n", <-out1)
}
