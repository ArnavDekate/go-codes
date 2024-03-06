package main

import "fmt"

func main() {
	var radius float32
	var area float32
	const Pi = 3.14

	fmt.Printf("Enter Radius : ")
	fmt.Scanf("%f", &radius)

	area = Pi * radius * radius

	fmt.Printf("Area of Circle : %f\n", area)
	fmt.Printf("Area of Circle : %g\n", area)
	fmt.Printf("Area of Circle : %.2f\n", area)

	perimeter := 2 * Pi * radius

	fmt.Printf("Perimeter of Circle : %.2f\n", perimeter)
}
