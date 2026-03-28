// Exercício 2 - Declarando 3 variáveis em package scope e(int, string, bool) com valor zero

package main

	import "fmt"

	var x int
	var y string
	var z bool

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z)

}