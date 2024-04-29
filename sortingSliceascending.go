package main

import "fmt"
import "sort"

func main() {
	slice := []int{70,20,30,60,50,60,10,80,90,100}
	slice1 := []string{"honesty","is","the","best","policy"}

	sort.Ints(slice)
	sort.Strings(slice1)

	fmt.Println("Sorted slice:")
	for _, ele := range slice {
		fmt.Printf("%d \n",ele)
	}
	for _, item := range slice1 {
		fmt.Printf("%s ",item)
	}
}