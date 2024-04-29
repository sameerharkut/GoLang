package main

import "fmt"

func main() {
	x := 3

	switch x {
	case 1:
		fmt.Println("1")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	case 5:
		fmt.Println("3")

	}
}
