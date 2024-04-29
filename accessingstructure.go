package main
import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

type Person struct{
	name string
	age int
	job string
	salary int
}

func main() {
	// var book1 Books
	// var book2 Books

	// book1.title = "Go Programming"
	// book1.author = "Mahesh Kumar"
	// book1.subject = "Go Programming Tutorial"
	// book1.book_id = 452762

	// book2.title = "Competitive Programming"
	// book2.author = "Zara Ali"
	// book2.subject = "Competitive Programming Tutorial"
	// book2.book_id = 783632

	// fmt.Printf("Book 1 title : %s\n", book1.title)
	// fmt.Printf("Book 1 author : %s\n", book1.author)
	// fmt.Printf("Book 1 subject : %s\n", book1.subject)
	// fmt.Printf("Book 1 book_id : %d\n", book1.book_id)

	// fmt.Printf("Book 2 title : %s\n", book2.title)
	// fmt.Printf("Book 2 author : %s\n", book2.author)
	// fmt.Printf("Book 2 subject : %s\n", book2.subject)
	// fmt.Printf("Book 2 book_id : %d\n", book2.book_id)


	var person1 Person
	var person2 Person

	person1.name = "Sameer"
	person1.age = 20
	person1.job = "SDE"
	person1.salary = 90000

	person2.name = "Rohit"
	person2.age = 24
	person2.job = "SDE"
	person2.salary = 80000

	fmt.Printf("Book 1 title : %s\n", person1.name)
	fmt.Printf("Book 1 author : %d\n", person1.age)
	fmt.Printf("Book 1 subject : %s\n", person1.job)
	fmt.Printf("Book 1 book_id : %d\n", person1.salary)

	fmt.Printf("Book 2 title : %s\n", person2.name)
	fmt.Printf("Book 2 author : %d\n", person2.age)
	fmt.Printf("Book 2 subject : %s\n", person2.job)
	fmt.Printf("Book 2 book_id : %d\n", person2.salary)

}