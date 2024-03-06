package main

import "fmt"

var sum int

func sq(n int) {
	fmt.Printf("%d is square of %d\n", n*n, n)

	sum = 5 + 9

}
func main() {
	sq(3)
	sq(9)
	sq(18)
	println(sum)

}
