package main

import (
	"fmt"
)

func main() {
	var name, lname string

	fmt.Print("Enter name : ")
	fmt.Scanf("%s %s", &name, &lname)
	fmt.Printf("%s %s", name, lname)

}
