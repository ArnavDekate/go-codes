package main
import "fmt"

func main()  {
	var x int 
	print("Enter the number : ")
	fmt.Scanf("%d\n",&x)
	mul(x)
}

func mul(x int){
	mult := 1
	for mult<=10{
		product := x*mult

		fmt.Printf("%d * %d = %d\n",x,mult,product)
		mult++
	}
}