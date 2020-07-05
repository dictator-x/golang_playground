package function

import "fmt"

type person struct {
}

func (person) sayHello() {
	fmt.Println("say hello")
}
func Run() {
	fmt.Println(">>> function demo")
	p := person{}
	p.sayHello()
}
