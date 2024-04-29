package main
import "fmt"

func main() {

	type Rectangle struct {
		length int
		breadth int
	}

	rect := Rectangle{20,22}

	fmt.Println("Length:",rect.length)
	fmt.Println("Breadth:",rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area:",area)
}