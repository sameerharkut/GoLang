package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// colors := []string{"red", "green", "blue"}
	// result,_ := json.Marshal(colors)
	// fmt.Println(string(result))
	// fmt.Println(result)

	boolVal,_ := json.Marshal(true)
	fmt.Println(string(boolVal))

	intVal,_ := json.Marshal(25)
	fmt.Println(string(intVal))

	floatVal,_ := json.Marshal(3.58)  
	fmt.Println(string(floatVal))
}