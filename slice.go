package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}

	fmt.Println("Numbers : ", num)
	fmt.Println("Numbers : ", num[0:])
	fmt.Println("Numbers : ", num[1:4])
	fmt.Println("Numbers : ", num[:])

	//slice if unspecified size
	var nums []int

	nums = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)

	nums[3] = 10
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)

}
