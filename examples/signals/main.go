package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.TODO(), 10 * time.Second)
	defer cancel()

	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<- ctx.Done()

	fmt.Println("Stop signal received. Shutting down")
}