package main

import "fmt"
import "os"

func main() {
	fmt.Println("Total Arguments: ", len(os.Args))
	fmt.Println("Program Name: ",programName)

	fmt.Println("\nArguments:")
	for i:=1; i<len(os.Args); i++ {
		fmt.Println("\tArgument[%d]: %s\n", i, os.Args[i])
	}
}