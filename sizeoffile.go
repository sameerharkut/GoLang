package main
import "fmt"
import "os"

func main() {
	MyFile, err := os.Stat("test.txt")
	if err != nil{
		fmt.Println("File does not exist")
	}
	fmt.Println("Size of files in bytes: ",MyFile.Size())
}