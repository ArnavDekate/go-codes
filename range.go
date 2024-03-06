package main

import "fmt"

func main(){
	colors := []string{"arnav","manas","paras"}
	for i,v := range colors{
		fmt.Println(i,"-",v)
	}
	array()
}

func array(){
	arr := [...]int{1,3,5,8,6,7,8,5}
	fmt.Println(arr)
	arr[6]=3
	fmt.Println(arr)
}