package main
import "fmt"
import "os"
func main() {
	var name string
	var age uint

	if _,err := fmt.Scan(&name,&age); err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Your name is:%s\n",name)
	fmt.Printf("Your age is:%d\n",age)
}