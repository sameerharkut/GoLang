package main
import "fmt"
import "os"

func main() {
	empFile, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Unable to create file")
	} else {
		fmt.Println("File creation done")
	}
	empFile.Close()
}