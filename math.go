package main

import (
	"fmt"
	"math"
)

func main() {
	var val float64 = -12.37

	fmt.Printf("Abs value of %.2f is %.2f\n", val, math.Abs(val))

	num1, num2 := 10.23, 73.
	large := math.Max(num1, num2)
	small := math.Min(num1, num2)

	println(large)
	println(small)

	n1, n2 := 10.4, 3.4
	pow := math.Pow(n1, n2)

	println(pow)

	a, b := 10.25, 20.52

	var r float64 = 0

	r = math.Floor(a)
	fmt.Printf("Result is : %f", r)
	r = math.Ceil(b)
	fmt.Printf("Result is : %f", r)

}
