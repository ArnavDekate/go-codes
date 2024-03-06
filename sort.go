package main

import (
	"fmt"
	"sort"
)

func main() {
	str := [6]int{1, 5, 3, 6, 2, 7}
	srt(str)

	str1 := []float64{1.2, 3.4, 43.3, 4.3, 2.3, 5.4, 1.5}
	srtf(str1)
}

func srt(str [6]int) {
	fmt.Println("Unsorted : ", str)
	sort.Ints(str[:])
	fmt.Println("Sorted : ", str)

}

func srtf(str1 []float64) {
	fmt.Println("Unsorted : ", str1)
	sort.Float64s(str1)
	fmt.Println("Sorted : ", str1)

}
