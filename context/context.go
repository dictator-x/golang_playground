package context

import (
	"context"
	"fmt"
	"math/rand"
)

func Run() {
	fmt.Println("Golang Context demo")
	fmt.Printf("%T\n", context.Background())
	fmt.Printf("%p\n", context.Background())
	startEatBurger()
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
