package main
import "fmt"

func main() {
	i:=0
	// loop: //Label for the goto statement
	// fmt.Println(i)
	// i++
	// if i<5{
	// 	goto loop // Jump to the loop label
	// }
	// fmt.Println("Loop ends here")



	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i>4 {
	// 		break
	// 	}
	// }


	loop:
	if i<=4 {
		fmt.Println(i)
		i++
		goto loop
	}


}