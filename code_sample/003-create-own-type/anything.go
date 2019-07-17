package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 43
	b = 68
	fmt.Println(a)
	fmt.Printf("%t\n", a)
	fmt.Println(b)
	fmt.Printf("%t\n", b)
	a = int(b) //conversion, GO doesnt call it "casting"
	fmt.Println(a)
	fmt.Printf("%t\n", a)
}
