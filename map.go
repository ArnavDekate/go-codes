package main

import "fmt"

func main() {

	submarks := map[string]float32{"Golang": 84, "Java": 80, "Python": 73}

	fmt.Println(submarks)

	flrcolor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}

	fmt.Println(flrcolor["Sunflower"]) //Yellow

	fmt.Println(flrcolor["Jasmine"]) //White

	fmt.Println(flrcolor["White"]) //No Output

	capital := map[string]string{"India": "Delhi", "Nepal": "Katmandu", "US": "New York"}
	fmt.Println(capital)

	capital["US"] = "Washington DC"

	fmt.Println(capital)
}
