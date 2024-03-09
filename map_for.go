package main

import "fmt"

func main() {

	sqnum := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	for number, squared := range sqnum {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}
}
