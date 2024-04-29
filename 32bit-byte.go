package main

import "fmt"

func main() {
	var intNum32 int32 = 123
	var byteNum byte = 0

	byteNum = byte(intNum32)
	fmt.Println("byte int number:", byteNum)
}
