package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter 3 Numbers : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	if a > b {
		if c > a {
			println(c, "is the largest number")
		} else {
			println(a, "is the largest number")
		}
	} else {
		if c > b {
			println(c, "is the largest number")
		} else {
			println(b, "is the largest number")
		}
	}
}
