package main

import "os"
import "io"
import "fmt"

func main() {
	existingFile, err := os.Open("Sample.txt")

	if err != nil {
		fmt.Println("Unable to open file")
	}

	CopyFile, err := os.Create("Copy.txt")
	if err != nil {
		fmt.Println("Unable to create file")
	}
	len, err := io.Copy(CopyFile, existingFile)
	if err !=  nil {
		fmt.Println("Unable to copy file")
	}
	fmt.Printf("%d bytes copied successfully",len)

	existingFile.Close()
	CopyFile.Close()
}