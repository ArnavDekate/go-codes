package main

import (
	"fmt"
)

func main() {

	var x, y int
	print("enter 2 numbers : ")
	fmt.Scan(&x, &y)

	if x%2 == 0 {
		print("x is even\n")
	} else {
		print("x is odd\n")
	}

	// fmt.Printf("%d is abs value of x\n", int(math.Abs(float64(x))))

	if x > 0 {
		fmt.Printf("%d is abs value of x\n", x)
	} else {
		fmt.Printf("%d is abs value of x\n", x*-1)
	}

}

func add(a, b) {
	a = a + b
	print(a)
}
