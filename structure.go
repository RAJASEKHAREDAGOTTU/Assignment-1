package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book Books   
   var Book1 Books   
   Book.title = "Python"
   Book.author = "Rajsekhar"
   Book.subject = "python Programming Tutorial"
   Book.book_id = 10007  
   Book1.title = "Golang"
   Book1.author = "raju"
   Book1.subject = "Raju  Tutorial"
   Book1.book_id = 123456
   printBook(&Book)
   printBook(&Book1)
}
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}