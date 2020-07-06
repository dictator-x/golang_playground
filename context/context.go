package context

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	fmt.Println("Golang Context demo")
	fmt.Printf("%T\n", context.Background())
	fmt.Printf("%p\n", context.Background())
	startEatBurger()
	withTimeoutDemo()
	for {
	}
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " get message to quit")
			return
		default:
			fmt.Println(name, "is running")
			time.Sleep(time.Second)
		}
	}
}

func withTimeoutDemo() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go work(ctx, "work1")
}

func withCancelDemo() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithCancel(context.Background())
	go work(ctx, "work1")
}

func startEatBurger() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := eatBurger(ctx)
	for n := range eatNum {
		if n >= 10 {
			cancel()
			break
		}
	}
}

func eatBurger(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 0
	t := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("using %d second, eating %d burgers\n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("eating %d burgers\n", n)
			}
		}
	}()

	return c
}
