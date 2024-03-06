package main

import "fmt"

func count(num int) {
	fmt.Println(num)

	if num > 1 {
		count(num - 1)
	}
}

func main() {
	fmt.Println("Countdown Starts : ")
	count(5)

}
