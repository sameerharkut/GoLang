package main

import (
	"fmt"
	"time"
)

func main() {

	YYYY, MM, DD := time.Now().Date()
	fmt.Println("Date : ", DD)
	fmt.Println("Month : ", MM)
	fmt.Println("Year : ", YYYY)
}
