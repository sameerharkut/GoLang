package main

import "fmt"

func main() {
	var floatNum64 float64 = 76.28
	var floatNum32 float32 = 0

	floatNum32 = float32(floatNum64)
	fmt.Println("32-Bit float number:", floatNum32)
}
