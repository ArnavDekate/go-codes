package main

import "fmt"

func main() {
	integer := 23
	fmt.Printf("%T %T\n", integer, &integer)
	fmt.Print("%T %T\n", integer, &integer)
	fmt.Printf("\n%T %v\n", integer, &integer)
	fmt.Printf("%T %d\n", integer, &integer)

	fl := 1234.38396443
	fmt.Printf("%5.4f\n", fl)

	str := "txt text"
	fmt.Printf("%*s %s\n", 40, str)
	fmt.Printf("%040s %s\n", str)

	var val, st = 'A', "A"
	fmt.Printf("Char value : %c \n", val)
	fmt.Printf("ASCII value : %d \n", val)
	fmt.Printf("%T , %T \n", val, st)

}
