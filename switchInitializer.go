package main
import (
	"fmt"
	// "time"
)

func main() {
	// switch today := time.Now(); {
	// case today.Day() < 5:
	// 	fmt.Println("Clean your house.")
	// case today.Day() <= 10:
	// 	fmt.Println("Buy some wine.")
	// case today.Day() <= 15:
	// 	fmt.Println("Visit a Doctor.")
	// case today.Day() <= 25:
	// 	fmt.Println("Buy some food.")
	// default:
	// 	fmt.Println("No information available for that day.")

	// }

	// switch day:=4; day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thrusday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// default:
	// 	fmt.Println("Sunday")
	// }


	var x interface{}="RKNEC"
	switch i := x.(type){
	case nil:
		fmt.Printf("type of x:%T",i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}
}
