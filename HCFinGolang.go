package main

import "fmt"

func main() {
	var num1 int = 0
	var num2 int = 0
	var temp int = 0

	fmt.Printf("Enter number1: ")
	fmt.Scanf("%d", &num1)

	fmt.Printf("Enter number2: ")
	fmt.Scan(&num2)

	for num2 != 0 {
		temp = num1 % num2
		num1 = num2
		num2 = temp
	}

	fmt.Printf("Highest common factor is:%d", num1)
}
