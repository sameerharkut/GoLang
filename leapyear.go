package main
import "fmt"

func main(){
	var year int

	fmt.Printf("Enter year:")
	fmt.Scanf("%d",&year)

	if(year%4==0 && year%100!=0)
}