package main

import (
	"fmt"
	"github.com/dictator-x/goplayground/context"
	"github.com/dictator-x/goplayground/point"
	"github.com/dictator-x/goplayground/slice"
)

func main() {
	fmt.Println(context.Run())
	point.Run()
	slice.Run()
}
