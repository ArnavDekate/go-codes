package main

import "fmt"

func main() {
	name := "Arnav Dekate"
	roll := 30
	branch := "Cyber Security"
	year_sem := "Third/Sixth"
	paper1 := 75
	paper2 := 85
	paper3 := 90

	fmt.Println("\t\t     CERTIFICATE")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Name : %s", name)
	fmt.Printf("\t\tRoll No : %v\n", roll)
	fmt.Printf("Branch : %s", branch)
	fmt.Printf("\t\tYear & Semester : %s\n", year_sem)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Score :-\n")
	fmt.Println("Paper 1 : ", paper1, "/100")
	fmt.Println("Paper 2 : ", paper2, "/100")
	fmt.Println("Paper 3 : ", paper3, "/100")
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Total Score : %d/300", paper1+paper2+paper3)
	fmt.Printf("\t\tPercentage : %d\n", (paper1+paper2+paper3)/3)
	fmt.Println("--------------------------------------------------------------")

	if paper1 > 40 && paper2 > 40 && paper3 > 40 {
		print("PASS")
	} else {
		print("FAIL")
	}
}
