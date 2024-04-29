package main

import "fmt"
import "os"
import "strings"

func main() {
	var envDetails []string = os.Environ()

	for envIndex, envVar := range envDetails {
		Value := strings.Split(envVar, "=")
		fmt.Printf("%d -> (%s: %s)\n", envIndex, Value[0], Value[1])
	}
}