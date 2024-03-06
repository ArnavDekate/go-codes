package main

import "fmt"

func main() {

	// var greet = func() {
	// 	fmt.Println("Hello , How are you!!")
	// }

	// greet()

	// welcome := greet
	// welcome()

	// fmt.Printf("%T is type of greet\n", greet)

	//anonymous func with parameters

	var sum = func(num1, num2 int) int {
		result := num1 + num2
		return result
	}

	r := sum(3, 4)
	fmt.Println("Sum : ", r)
}
