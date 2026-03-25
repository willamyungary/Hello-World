package main

import "fmt"

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 55
	fmt.Println(x)
	fmt.Println("↑ foi x.\n↓ foi y.")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}