package main

import "fmt"

func main() {
	x := 8
	y := 10

	if x < y {
		fmt.Println("x is less than y")
	} else // else cannot be on next line
	{
		fmt.Println("x is greater than or equal to y")
	}
}
