package main

	import "fmt"

	var w float64 = 3.1416
	var x int = 44
	var y string = "Jesus Cristo"
	var z bool = true

func main() {
	s :=fmt.Sprintf("%v\t%v\t%v\t%v", w, x, y, z)
	fmt.Println(s)
}
