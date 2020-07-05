package main

import (
	// "fmt"
	"github.com/dictator-x/goplayground/channel"
	"github.com/dictator-x/goplayground/context"
	"github.com/dictator-x/goplayground/function"
	"github.com/dictator-x/goplayground/interfacedemo"
	"github.com/dictator-x/goplayground/point"
	"github.com/dictator-x/goplayground/slice"
)

func main() {
	point.Run()
	slice.Run()
	context.Run()
	interfacedemo.Run()
	channel.Run()
	function.Run()
}
