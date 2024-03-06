package main

import "fmt"

func sum(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + sum(num-1)
	}
}

func main() {
	var num int
	print("Enter Number : ")
	fmt.Scan(&num)

	if num <= 0 {
		print("Invalid Input!!!")
	} else {
		// func call
		var result = sum(num)
		fmt.Println("\nSum : ", result)
	}
}
