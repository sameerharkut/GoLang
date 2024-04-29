package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		sum = sum + numbers[i]
	}
	fmt.Println("Sum of all numbers:", sum)
}
