package main

import "fmt"

func main()  {
	// var a int
	// var b int

	// fmt.Scanf("%d",&a)
	// fmt.Scanln("%d",&b)

	// var fname,lname string

	// fmt.Print("Enter your first name and last name: ")
	// fmt.Scanln(&fname,&lname)
	// fmt.Println("Hello",fname,lname)

// Addition of two strings
    fmt.Print("Enter First String: ")
	var first string
	fmt.Scanln(&first)
	fmt.Printf("Enter Second String: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first + second)
}