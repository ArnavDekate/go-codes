package main

import "fmt"

func main() {
	num := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}
