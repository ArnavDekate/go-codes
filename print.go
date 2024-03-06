package main

import "fmt"

func main() {
	var i float32 = 15.5
	var txt = "Hello!!"
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%s\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("'%v'\n", txt)
	fmt.Printf("%T", txt)

}
