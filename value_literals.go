package main

import "fmt"

func main() {
	println(15 == 017)
	fmt.Println(15 == 0b1111)
	fmt.Println(15 == 0xF)

	//increments in integers
	a := 10
	println(a)
	a++
	println(a)
	a--
	print(a) //built-in functions

	b := 10.4
	b++
	fmt.Printf("\n%f", b)

	// ++a : gives error
	// println(a)

	fmt.Printf("\nBinary of a is : %b", a)
	fmt.Printf("\nOctal of a is : %o", a)
	a++
	fmt.Printf("\nHexadecimal of a is : %x", a)

	fmt.Printf("\n%t", a == 0b1011)
	fmt.Printf("\nExpression is %t", 0xF == 15)
	fmt.Printf("\nExpression is %T\n", 0xF == 15)

	var i, j string = "Hello", "World"

	print(i)
	print(j, "\n")

	fmt.Printf("%s", i)
	fmt.Printf("%s\n", j)

	name := "Arnav"
	lname := "Dekate"

	fmt.Print("My name is ", name, " and last name is ", lname)
	fmt.Printf("\nMy name is %s and last name is %s\n", name, lname)

	fmt.Print(i, " ", j)
	fmt.Printf("\n%s %s", i, j)

}
