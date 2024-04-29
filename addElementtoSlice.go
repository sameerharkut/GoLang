package main

import "fmt"

func main() {
	//	primeNumbers := []int{2,3}

	evenNumbers := []int{2, 4}

	oddNumbers := []int{1, 3}

	//	primeNumbers = append(primeNumbers,5,7)

	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers:", evenNumbers)
	printSlice(evenNumbers)
}
func printSlice(x []int) {
	fmt.Printf("len:%d , Capacity:%d , Slice:%v", len(x), cap(x), x)
}
