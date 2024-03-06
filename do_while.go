package main

import "fmt"

func main()  {
	num := 1
	for{
		fmt.Printf("%d\n",num);
		num++
		if num>5{
		break
		}
	}
}