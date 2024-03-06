package main

import (
	"fmt"
)

var c, x int = 0, 1

func main() {
	fmt.Println("Enter the no : ")
	var a, result int

	fmt.Scan(&a)

	result = fib(a)

	fmt.Printf("%d is last element of fib series", result)
}

func fib(a int) int {

	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return fibhelp(c, x, a)
	}

}

func fibhelp(num1 int, num2 int, a int) int {
	var c, t int = a, 0
	if c >= 1 {
		t = num2
		num2 = num1 + num2
		num1 = t
		c--
		num2 = fibhelp(num1, num2, a)

	}
	return num2

}
