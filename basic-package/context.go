package basicpackage

import (
	"context"
	"fmt"
	"time"
)

func goroutine1(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go goroutine1(ctx, ch)

ctxloop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break ctxloop
		case <-ch:
			fmt.Println("success")
			break ctxloop
		}
	}
}
