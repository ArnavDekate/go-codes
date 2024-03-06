package main

import "fmt"

// initialize func rectangle (type)
type rec func(int, int) int

type rectanglepara struct {
	len   int
	bdt   int
	color string
	//function feild as a function
	rect rec
}

func main() {
	result := rectanglepara{
		len:   10,
		bdt:   20,
		color: "red",
		rect: func(len int, bdt int) int {
			return len * bdt
		},
	}

	fmt.Println("color of rectangle : ", result.color)
	fmt.Println("Area of rectangle : ", result.rect(result.len, result.bdt))

}
