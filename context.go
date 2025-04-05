package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey string

func main() {
	// 1. context.Background()
	fmt.Println("== Background ==")
	rootCtx := context.Background()

	// 2. context.TODO()
	fmt.Println("== TODO ==")
	todoCtx := context.TODO()
	_ = todoCtx // just a placeholder

	// 3. context.WithCancel()
	fmt.Println("== WithCancel ==")
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		<-cancelCtx.Done()
		fmt.Println("WithCancel cancelled:", cancelCtx.Err())
	}()
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)

	// 4. context.WithTimeout()
	fmt.Println("== WithTimeout ==")
	timeoutCtx, timeoutCancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer timeoutCancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-timeoutCtx.Done():
		fmt.Println("WithTimeout triggered:", timeoutCtx.Err())
	}

	// 5. context.WithDeadline()
	fmt.Println("== WithDeadline ==")
	deadline := time.Now().Add(1 * time.Second)
	deadlineCtx, deadlineCancel := context.WithDeadline(rootCtx, deadline)
	defer deadlineCancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation completed")
	case <-deadlineCtx.Done():
		fmt.Println("WithDeadline triggered:", deadlineCtx.Err())
	}

	// 6. context.WithValue()
	fmt.Println("== WithValue ==")
	valueCtx := context.WithValue(rootCtx, ctxKey("userID"), 42)
	doSomething(valueCtx)
}

func doSomething(ctx context.Context) {
	userID := ctx.Value(ctxKey("userID"))
	fmt.Println("Value from context:", userID)
}
