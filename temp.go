package main

import "fmt"

func main() {
	var ctemp, ftemp, f, c float64

	fmt.Print("Enter temprature in Farenheit :")
	fmt.Scanf("%f", &ftemp)
	ctemp = (ftemp - 32) / 1.8
	fmt.Printf("Temprature in Celcius : %.2f", ctemp)

	fmt.Print("Enter temprature in Celcius :")
	fmt.Scanf("%f", &c)
	f = (c * 1.8) + 32
	fmt.Printf("Temprature in Celcius : %.2f", f)
}
