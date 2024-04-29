package main

import "fmt"

func main(){
	x:=10+
	y:=10

	if(x!=y){
		if(x<y){
			fmt.Println("x is less than y")
		} else if(x>y){
			fmt.Println("x is greater than y")
		}
	} else{
		fmt.Println("x is equal to y")
	}
}