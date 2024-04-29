package main

import (
	"fmt"
	"time"
)

func main() {
	currentDateTime := time.Now()

	ANSIC_FORMAT := currentDateTime.Format(time.ANSIC)
	UnixDate_Format := currentDateTime.Format(time.UnixDate)

}
