package main

import (
	"fmt"
	"time"
)

func main() {
	var currentDateAndTime string

	currentDateAndTime = time.Now().String()

	fmt.Println("Current date and time\n", currentDateAndTime)
}
