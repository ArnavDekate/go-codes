package main

import (
	"fmt"
	"os"
)

func main() {
	// if !(20 > 10) {
	// 	fmt.Println("20 is greater than 10")
	// }

	// if true {
	// 	fmt.Println("TRUE")
	// }

	// x := 10
	// if y := 20; x < y {
	// 	fmt.Println("x is less then y")
	// }

	/* handling error*/
	var name string
	var age int

	if _, err := fmt.Scan(&name, &age); err != nil {
		fmt.Print(err, "\n")
		os.Exit(1)
	}
	fmt.Printf("Your name is : %s\n", name)
	fmt.Printf("Your age is : %d\n", age)

}
