package main

import "fmt"

func addnum(n1 int, n2 int) (int, int) {
	// n1 := 12
	// n2 := 20

	n3 := n1 + n2
	n4 := n1 - n2

	// println(n3)

	return n3, n4

	// println("After return")
}

func main() {
	var a, b int
	print("Enter 2 numbers : ")
	fmt.Scan(&a, &b)
	sum, diff := addnum(a, b)

	fmt.Println("Result : ", sum, diff)

}
