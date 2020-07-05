package interfacedemo

import "fmt"

type Context interface {
	Done()
}

type ContextA struct{}

func (*ContextA) Done() {
	fmt.Println("ContextA done")
}

func useContext(context Context) {
	fmt.Println(context)
}

func Run() {
	fmt.Println(">>> interface demo")
	useContext(&ContextA{})
}
