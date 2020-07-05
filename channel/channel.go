package channel

import "fmt"

func Run() {
	fmt.Println(">>> demo chan")
	c := make(chan int)
	go func() {
		for {
			c <- 1
		}
	}()
	go func() {
		for {
			c <- 2
		}
	}()
	go func() {
		for {
			fmt.Printf("gorount1: %d\n", <-c)
		}
	}()
	go func() {
		for {
			fmt.Printf("gorount2: %d\n", <-c)
		}
	}()
	// for {

	// }
	fmt.Println(make(chan struct{}) == make(chan struct{}))
}
