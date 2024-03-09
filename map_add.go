package main

import "fmt"

func main() {
	students := map[int]string{1: "john", 2: "lily"}
	fmt.Println("Initial Map :", students)

	students[5] = "Robin"

	students[3] = "Rob"

	fmt.Println("Updated Map :", students)

	delete(students, 1)
	fmt.Println("Updated Map :", students)

}
