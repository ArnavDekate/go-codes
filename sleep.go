package main

import "time"
import "fmt"

func main(){
	println("Hello")
	time.Sleep(10000*time.Millisecond)
	fmt.Println("Hi")
}