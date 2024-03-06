package main

import "fmt"

// func test(x int){
// 	fmt.Println("Hello",x)
// }

// func main(){
// 	x := test
// 	x(5)
// 	test(20)
// }

// --------------------------------------

// var test = func(x int){
// 	fmt.Println(x)
// }

// func main()  {

// 	test(4)


// -----------------------------------------------

func main()  {
	test := func (x int)  {
		return x*-1
	}(8)

	println(test)
	
}
