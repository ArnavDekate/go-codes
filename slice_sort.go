package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"honsety", "is", "the", "best", "policy"}
	sort.Strings(slice)

	fmt.Println("Sorted Slice : ")
	for _, item := range slice {
		fmt.Printf("%s ", item)
	}
}
