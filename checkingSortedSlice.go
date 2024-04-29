package main

import "fmt"
import "sort"

func main() {
	var status bool = false

	slice1 := []int{10,20,30,40,50,60,70,80,90,100}
	slice2 := []int{70,20,80,50,40,10,90,30,100,60}

	slice3 := []float64{1.34,2.76,3.45,4.78,5.89}
	slice4 := []float64{4.56,2.38,1.90,6.89,2.90}

	slice5 := []string{"is","my","name","sameer"}
	slice6 := []string{"my","name","is","sameer"}

	status = sort.IntsAreSorted(slice1)
	if status == true {
		fmt.Println("Slice1 is sorted")
	} else {
		fmt.Println("Slice1 is not sorted")
	}
	status = sort.IntsAreSorted(slice2)
	if status == true {
		fmt.Println("Slice2 is sorted")
	} else {
		fmt.Println("Slice2 is not sorted")
	}

	status = sort.Float64sAreSorted(slice3)
	if status == true {
		fmt.Println("Slice3 is sorted")
	} else {
		fmt.Println("Slice3 is not sorted")
	}
	status = sort.Float64sAreSorted(slice4)
	if status == true {
		fmt.Println("Slice4 is sorted")
	} else {
		fmt.Println("Slice4 is not sorted")
	}
	
	status = sort.StringsAreSorted(slice5)
	if status == true {
		fmt.Println("Slice5 is sorted")
	} else {
		fmt.Println("Slice5 is not sorted")
	}
	status = sort.StringsAreSorted(slice6)
	if status == true {
		fmt.Println("Slice6 is sorted")
	} else {
		fmt.Println("Slice6 is not sorted")
	}
}