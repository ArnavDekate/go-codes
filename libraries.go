package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var num1 int = 10
	var num2 string = "gdhdhsskddj"

	fmt.Printf("Size of Num1 : %d\n", unsafe.Sizeof(num1))
	fmt.Printf("Size of Num2 : %d\n", unsafe.Sizeof(num2))

	a := 10
	b := 20.10
	c := "hello"
	d := true

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	fmt.Println("Type of a = ", reflect.TypeOf(a))
	fmt.Println("Type of b = ", reflect.TypeOf(b))
	fmt.Println("Type of c = ", reflect.TypeOf(c))
	fmt.Println("Type of d = ", reflect.TypeOf(d))

	defer fmt.Println("Type of a = ", reflect.ValueOf(a).Kind())
	defer fmt.Println("Type of b = ", reflect.ValueOf(b).Kind())
	defer fmt.Println("Type of c = ", reflect.ValueOf(c).Kind())
	fmt.Println("\nType of d = ", reflect.ValueOf(d).Kind())

}
