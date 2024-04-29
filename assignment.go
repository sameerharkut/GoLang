package main

import "fmt"

func main(){
	var roll_no int
	var name string
	var paper1 int
	var paper2 int
	var paper3 int
    var total int
	var percentage float64

	fmt.Printf("Roll_no: \tName: \n")
	fmt.Scanf("%d %s\n",&roll_no, &name)
	fmt.Printf("Marks of Paper1: ")
	fmt.Scan(&paper1)
	fmt.Printf("Marks of Paper2: ")
	fmt.Scan(&paper2)
	fmt.Printf("Marks of Paper3: ")
	fmt.Scan(&paper3)

	total = paper1 + paper2 + paper3
	percentage = float64(total) / 3.0

	fmt.Println("Total: \t\tPercentage: \n",total, percentage)
}