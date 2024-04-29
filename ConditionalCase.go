package main
import "fmt"

func main(){
	x:=20
	y:=30

	switch{
	case x>y:
		fmt.Println("x is greater than y")
	case x<y:
		fmt.Println("x is less than y")
	default:
		fmt.Println("x is equal to y")
	}
}