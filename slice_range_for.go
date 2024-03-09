package main

import "fmt"

func main() {

	//create two slices
	primeNumbers := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3}

	//copy elements of primeNumbers to numbers
	copy(numbers, primeNumbers)

	fmt.Println("Numbers:", numbers)

	numbers = []int{2, 4, 6, 8, 10}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Iterating a Slice using RANGE in FOR loop

	//create an integer array
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice from 2 to 4(5)
	intSlice := arr[2:5]

	fmt.Println("Slice elements: ")
	for index, ele := range intSlice {
		fmt.Printf("Index = %d,element = %d\n", index, ele)
	}

}
