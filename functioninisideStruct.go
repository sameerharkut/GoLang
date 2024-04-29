package main

import "fmt"

type Rectangle func(int, int) int

type rectanglePara struct {
	length  int
	breadth int
	color   string

	area Rectangle
	perimeter Rectangle
}

func main() {

	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		area: func(length int, breadth int) int {
			return length * breadth
		},
		perimeter: func(length int, breadth int) int {
			return 2 * (length + breadth)
		},
	}
	// result1 := rectanglePara{
	// 	length:  10,
	// 	breadth: 20,
	// 	color:   "Red",
	// 	rect: func(length int, breadth int) int {
	// 		return 2 * (length + breadth)
	// 	},
	// }
	fmt.Println("Color of Rectangle:", result.color)
	fmt.Println("Area of Rectangle:", result.area(result.length, result.breadth))
	fmt.Println("Perimeter of Rectangle:", result.perimeter(result.length, result.breadth))


}
