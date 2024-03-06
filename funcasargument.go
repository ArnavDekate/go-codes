package main

import "fmt"

func findsq(num int) int {
	sq := num * num
	return sq
}

func main() {
	sum := func(n1, n2 int) int {
		return n1 + n2
	}

	result := findsq(sum(2, 4))
	fmt.Println("Result : ", result)
}
