package main
import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}


func main() {
	var book1 Books
	var book2 Books

	book1.title = "Go Programming"
	book1.author = "Mahesh Kumar"
	book1.subject = "Go Programming Tutorial"
	book1.book_id = 452762

	book2.title = "Competitive Programming"
	book2.author = "Zara Ali"
	book2.subject = "Competitive Programming Tutorial"
	book2.book_id = 783632

	printBook(book1)  
	printBook(book2)


}
func printBook(book Books){
fmt.Printf("Book title : %s\n", book.title)
fmt.Printf("Book author : %s\n", book.author)
fmt.Printf("Book subject : %s\n", book.subject)
fmt.Printf("Book book_id : %d\n", book.book_id)
}