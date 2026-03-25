package main

	import "fmt"

	var x int = 44
	var y string = "Jesus Cristo"
	var z bool = true

func main() {
	s :=fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}
