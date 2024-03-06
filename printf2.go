package main

import "fmt"

func main() {
	const n = "gfg"

	fmt.Printf("%s is a portal\n", n)

	var n1 int = 32
	fmt.Printf("Int value : %d\n", n1)

	var n2 float32 = 9.332
	fmt.Printf("Decimal value : %g\n", n2)

	var n3 int = 11
	fmt.Printf("Binary value : %b\n", n3)

	var n4 int = 4 + 11
	fmt.Printf("Scientefic value : %e", n4)

	var a int = 10
	var b float32 = 10.6
	var name string = "Rajesh"
	fmt.Printf("%d %f %s\n", a, b, name)
	fmt.Printf("%v %v %v", a, b, name)

}
