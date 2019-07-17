package main

import "fmt"

var a = 15
var b string

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 42
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)

	fmt.Println(a, b)
	foo()
	fmt.Println(b)
}

func foo() {
	b = "Long"
	fmt.Println("foo", b)
}
