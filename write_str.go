package main

import (
	"fmt"
	// "io"
	// "os"
)

func main() {

	/* io.WriteString() Function */
	// var dd int = 17
	// var mm int = 04
	// var yy int = 2021
	// var str string
	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy)
	// io.WriteString(os.Stdout, str)

	/* Scan() Function */
	// var s, t string
	// var i int
	// var x float32
	// fmt.Printf("\nEnter String : ")
	// fmt.Scan(&s, &t, &i, &x)
	// fmt.Printf("Result : %s %s %d %g\n", s, t, i, x)

	/* int = unsigned int */
	// var a int = -9223372036854775808
	// var b uint = 0
	// b = uint(a)
	// fmt.Printf("a = %d, b = %d\n", a, b)

	/*math library */
	// var p float64 = 64
	// q := float32(math.Sqrt(float64(p)))
	// fmt.Printf("\n%g", q)

	/* Type Conversion */
	var num1 int = 42
	var num2 float64 = float64(num1)
	var num3 uint = uint(num2)
	fmt.Printf("Value of num1 is %v an num1 type is %T\n", num1, num1)
	fmt.Printf("Value of num1 is %v an num1 type is %T\n", num2, num2)
	fmt.Printf("Value of num1 is %v an num1 type is %T\n", num3, num3)
}
