package main

import (
	"fmt"
	"time"
)

func main() {
	dateTime := time.Now()

	fmt.Println(dateTime.Format("01-02-2006 15:04:05.000000000"))
	fmt.Println(dateTime.Format("01-02-2006"))
	fmt.Println(dateTime.Format("01-02-2006 15:04:05 Mon"))
	fmt.Println(dateTime.Format("01-02-2006 15:04:05"))
	fmt.Println(dateTime.Format("01-02-2006 15:04:05.000000"))
	fmt.Println(dateTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(dateTime.Format("01-02-2006 15:04:05 Tue"))
}
