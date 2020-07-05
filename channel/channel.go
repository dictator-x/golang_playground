package channel

import "fmt"

func Run() {
	fmt.Println(">>> demo chan")
	c := make(chan int)
	go func() {
		c <- 1
	}()
	<-c
	fmt.Println(make(chan struct{}) == make(chan struct{}))
}
