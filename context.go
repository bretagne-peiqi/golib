package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancelFunc()

	go func(ctx context.Context) {
		fmt.Println("routine ...")
		//		<-ctx.Done()
		fmt.Println("tested in goroutine")
	}(ctx)

	if ctx.Err() == nil {
		time.Sleep(1 * time.Second)
		fmt.Println("err is nil")
	}
	if ctx.Err() != nil {
		fmt.Println("exited")
	}
}
