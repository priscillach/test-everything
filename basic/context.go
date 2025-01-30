package basic

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextDeadline() {
	w := &sync.WaitGroup{}
	w.Add(1)
	// same as ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	go Goroutine(ctx, w)
	w.Wait()
	fmt.Println("Context example finished")
}

func Goroutine(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Goroutine run")
		}
		time.Sleep(time.Second)
	}
}
