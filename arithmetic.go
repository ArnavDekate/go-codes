package main

import "fmt"

// Main Function
func main() {
	a := 10
	b := 20

	// Printing variables

	fmt.Printf("a : %d\tb : %d\n", a, b)

	// Addition

	sum := a + b
	fmt.Println("Sum = ", sum)

	// Substraction

	diff := b - a
	fmt.Println("Diffrence = ", diff)

	// Multiplication

	prod := a * b
	fmt.Println("Product = ", prod)

	// Division

	div := b / a
	fmt.Println("Quotient = ", div)

	// Increment
	a++
	fmt.Print("Incremented Value : ", a, "\n")

	var x int
	fmt.Scanf("\n%d", x)

	x += 4
	fmt.Print("\n", x)

	x -= 5
	fmt.Print("\n", x)

	x *= 2
	fmt.Print("\n", x)

	x /= 5
	fmt.Print("\n", x)

}
