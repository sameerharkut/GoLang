package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func WriteData() {
	file, err := os.Create("Sample.txt")
	if err != nil {
		fmt.Printf("Unable to open file:%s", err)
	}

	len, err := file.WriteString("Hello World")

	if err != nil {
		fmt.Printf("Unable to write data:%s", err)
	}
	file.Close()

	fmt.Printf("%d character written successfully into file", len)
}
func ReadData() {
	textData,err:=ioutil.ReadFile("Sample.txt")
	if err != nil {
		
	}
}