package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 5
	b := 2
	c := a & b
	fmt.Println("& operator : ", c)

	d := a | b
	println("| operator : ", d)

	e := a ^ b
	println("^ operator : ", e)

	var x int = 0b111100
	x >>= 3
	println(x)

	println(x > 5)
	println(x == 7)
	println(x < 5)

	var num1 int = 10
	var num2 string = "gdhdhsskddj"

	fmt.Printf("Size of Num1 : %d\n", unsafe.Sizeof(num1))
	fmt.Printf("Size of Num2 : %d\n", unsafe.Sizeof(num2))
}
