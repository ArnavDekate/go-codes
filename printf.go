package main

import "fmt"

func main() {
	const name, dept = "gfg", "cs"

	fmt.Printf("%s is a %s portal\n", name, dept)

	var i, j = 10, "Arnav"

	fmt.Print(i, j)

	println("\nUsing println instead of fmt.Println")
	print("Using print instead of fmt.Println")
}
