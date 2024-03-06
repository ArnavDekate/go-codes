package main

import "fmt"

func main() {
	var matrix [3][3]int

	fmt.Printf("Enter elements of matrix :\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Elements : matrix[%d][%d]:", i, j)
			fmt.Scanf("%d", &matrix[i][j])
		}

	}

	fmt.Printf("Matrix : \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			fmt.Printf("%d", matrix[i][j])
		}
		print("\n")
	}

	fmt.Printf("Left Diagonal ")
}
