package function

import "fmt"

type person struct {
}

func (person) sayHello() {
	fmt.Println("say hello")
}

type BaseNum struct {
	num1 int
	num2 int
}

type Add struct {
	BaseNum
}

type Sub struct {
	BaseNum
}

func use(base *BaseNum) {
	fmt.Println(base)
}

func Run() {
	fmt.Println(">>> function demo")
	p := person{}
	p.sayHello()

	// data := BaseNum{2, 3}
	// add := Add{data}
	// sub := Sub{data}

	// use(&add)
	// use(&sub)
}
