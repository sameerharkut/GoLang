package main

import "encoding/binary"
import "os"
import "fmt"

type Str struct {
	intNum uint8
	floatNum float32
}

func main() {
	file, err := os.Create("data.bin")
	if err != nil {
		fmt.Println("Couldn't open file")
	}

	var st = Str{10, 2.3}

	err = binary.Write(file, binary.LittleEndian, st)
	if err != nil {
		fmt.Println("Write failed")
	}
	fmt.Println("Structure written into file successfully")

	file.Close()
}
