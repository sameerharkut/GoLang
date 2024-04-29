package main
import "fmt"

func main() {
	result := display()
	fmt.Println("Welcome to", *result)
}

func display() *string{
	message := "SRKNEC"

	return &message
}