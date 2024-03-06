package main

import "fmt"

func main() {
	ogarr := []int{3, 4, 5, 2, 2, 3, 1}
	newitem := 6
	index := 3
	newarr := insertintoarr(ogarr, newitem, index)

	fmt.Println("Original Array : ", ogarr)
	fmt.Println("Item to be inserted : ", newitem)
	fmt.Println("At Index : ", index)
	fmt.Println("New Array : ", newarr)

}

func insertintoarr(arr []int, item int, index int) []int {
	newarr := make([]int, len(arr)+1)

	copy(newarr[:index], arr[:index])

	newarr[index] = item

	copy(newarr[index+1:], arr[index:])

	return newarr
}
