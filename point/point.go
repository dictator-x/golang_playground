package point

import "fmt"

type Address struct {
	street string
}

type Person struct {
	name    string
	age     uint8
	address Address
}

type Context interface {
}

type myType int

var (
	background = new(myType)
)

func newPerson() Context {
	return background
}
func (p Person) introSelf() {
	fmt.Printf("memory-person: %p\n", &p)
	fmt.Printf("memory-address: %p\n", &(p.address))
	fmt.Printf("memory-address: %p\n", &p.address)
	fmt.Printf("memory-address: %p\n", (&p).address)
}

func (p *Person) introSelf1() {
	fmt.Printf("memory-person: %p\n", p)
	fmt.Printf("memory-address: %p\n", &p.address)
}

func Run() {
	fmt.Println("Golang Struct && Point demo:")
	p := Person{"Jordan", 20, Address{"london"}}
	fmt.Println(p)
	fmt.Printf("memory: %p\n", &p)
	p.introSelf()
	p.introSelf1()

	l := p
	p.address.street = "a"
	fmt.Println(p)
	fmt.Println(l)
	fmt.Println(new(Person))

	fmt.Println("==")
	personAsParam(p)
	fmt.Println(p)

	var newp Person
	newp.name = "kobe"
	fmt.Println(newp)
	var newpp *Person
	fmt.Println(newpp)
	fmt.Printf("nil == nil: %v\n", newpp == nil)
	fmt.Println(newPerson())
}

func personAsParam(p Person) {
	p.address.street = "new street"
	fmt.Println(p)
}
