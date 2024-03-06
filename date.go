package main
import (
	"fmt"
	"time"
)

func main(){
	cdt := time.Now()
	
	day := cdt.Day()
	mon := cdt.Month()
	year := cdt.Year()

	hour := cdt.Hour()
	min := cdt.Minute()
	sec := cdt.Second()

	fmt.Println(day)
	fmt.Println(mon)
	fmt.Println(year)
	fmt.Println(hour)
	fmt.Println(min)
	fmt.Println(sec)
}

