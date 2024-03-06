package main

import "fmt"

var c, x int = 1, 0

func main() {
	fmt.Println("Enter the no : ")
	var a, result int

	fmt.Scan(&a)

	result = fact(a)

	fmt.Printf("%d is factorial of %d", result, a)
}

func fact(a int) int {

	if a <= 0 {
		return 1
	} else {
		return a * fact(a-1)
	}
	return c
}
