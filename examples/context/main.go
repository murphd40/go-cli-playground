package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		doSomething(ctx)
	}()

	wg.Wait()
}

func doSomething(ctx context.Context) {
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <- tick:
			fmt.Println("tock")
		case <- ctx.Done():
			fmt.Println("finished")
			return
		}
	}

}
