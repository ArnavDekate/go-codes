package main

import "fmt"

func main() {
	fmt.Println("Enter 1st String : ")
	var first, second string
	fmt.Scanln(&first) //input 1
	fmt.Println("Enter 2nd String : ")
	fmt.Scanln(&second)

	fmt.Println(first + " " + second) // addn of 2 strs

	var a, b int = 10, 20
	fmt.Println(a, b)
	fmt.Print(a + b)

}
