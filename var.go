package main

import "fmt"

var a = 10
func main(){
// 	fmt.Println("Hello World!")
// 	var a int = 50;
// 	b := 10;
// 	x := 20.5;
// 	y := "Sameer";
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(x)
// 	fmt.Println(y)

// 	var C int8 = 127
// 	C = C + 5
// 	fmt.Println(C)

// Variable Declaration without Initial Value
    var p string
	var q int
	var r bool

	fmt.Print(p)
	fmt.Print(q)
	fmt.Print(r)

// Mixed Variable Declaration in Go
    var f, g, h = 3, 4, "foo"
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Printf("f is of type %T\n",f)
	fmt.Printf("g is of type %T\n",g)
	fmt.Printf("h is of type %T\n",h)


// 	var m, n, o int = 1, 3, 5
// 	fmt.Println(m)
// 	fmt.Println(n)
// 	fmt.Println(o)

// Different types of variables in same line
    var d, e = 6.5, "Hello"
	i, j := 7.5, "World!"
	fmt.Print(d,"\t",e,"\n")
	fmt.Print(i,"\t",j,"\n")
	print(i,"\t",j,"\n")

// Dynamic Type Declaration	/ Type inference in Go
    var w float64 = 20.0
	k := 42
	fmt.Println(w)
	fmt.Println(k)
	fmt.Printf("Value of w is %f\n",w)
	fmt.Printf("Value of k is %d\n",k)
	fmt.Printf("w is of type %T\n",w)
	fmt.Printf("k is of type %T\n",k)

// // Go Variable Declaration in a block
//     var(
// 		a int
// 		b int = 1
// 		c string = "Hello"
// 	)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)

	fmt.Printf("%d\n",a)
}