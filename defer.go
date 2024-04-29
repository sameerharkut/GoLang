package main
import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// 	defer fmt.Println("Bye Bye")
// 	defer fmt.Println("wow jaishree krishna ekdum krishna lageche ")
// 	fmt.Println("Welcome to Golang")
// }

func show() {
	fmt.Println("Last")
}
func main() {
	defer show()
	fmt.Println("Hello")
	fmt.Println("World")
	defer display()
}
func display() {
	fmt.Println("Line")
}