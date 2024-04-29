package main

import "fmt"

// Main function
func main()  {
	const PI = 3.14
	fmt.Println(PI)

// Typed constant 
    const A int = 1
	const B float32 = 5.4
    fmt.Println(A)
	fmt.Println(B)

// Untyped constant 
    const C = 2
    fmt.Println(C)

// Multiple Constants Declaration
    const (
		X int = 1
		Y = 3.14
		Z = "Hi!"
	)
	fmt.Println(X)
	fmt.Print(Y)
    fmt.printf(Z)

// Type of Constant
    const V int = 1
	fmt.Printf("%T",V)
	const A = 5.32
	fmt.Printf("%T",A)

// Unchangeable and Read-only
    const B = 1
	B = 2
	fmt.Println(B)
}