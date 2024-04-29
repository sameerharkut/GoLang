package main

import "fmt"

func main() {
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	sliceNumbers := numbers[4:7]
	fmt.Printf("Type of slice numbers:%T\n", sliceNumbers)
	fmt.Printf("Type of array numbers:%T\n", numbers)
	fmt.Println(sliceNumbers)
}
