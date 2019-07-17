package main

import "fmt"

type MyType int

var x MyType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%t", y)
}
