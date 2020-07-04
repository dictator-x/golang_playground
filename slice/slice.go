package slice

import "fmt"

func Run() {
	fmt.Println(">>>> slice and array demo")
	// primes := []
	var a [3]int
	fmt.Println(a)
	var b *[3]int
	fmt.Println(b)
	b = &[3]int{1, 2, 3}
	fmt.Println(b)
	c := [...]int{1, 2, 3}
	fmt.Println(c)
	arr := [5]*int{0: new(int), 1: new(int)}
	fmt.Println(arr)

	slice := make([]int, 3, 5)
	fmt.Println(slice)

	ss := slice
	ss[0] = 1
	fmt.Println(slice)
	fmt.Println(ss)
	fmt.Printf("memory-slice: %p\n", &slice)
	fmt.Printf("memory-slice: %p\n", &ss)

	sliceAsParam(slice)
	fmt.Println(slice)

	var myNum []int
	fmt.Println(myNum)

	var myNump *[]int
	fmt.Println(myNump)
}

func sliceAsParam(slice []int) {
	slice[2] = 100
	fmt.Printf("memory-slice: %p\n", &slice)
}
