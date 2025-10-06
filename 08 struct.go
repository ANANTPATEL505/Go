package main
import (
	"fmt"
)
func main() {
	
	type Book struct{
		name string
		id int
		author string
	}

	var n int
	fmt.Print("how many books do you want to enter : ")
	fmt.Scan(&n)

	books:=make(map[int]Book)
	var b Book

	for i:=0;i<n;i++{
		fmt.Printf("\nEnter details for book %d:\n", i+1)
		fmt.Print("Enter Book ID : ")
		fmt.Scan(&b.id)
		fmt.Print("Enter Book name : ")
		fmt.Scan(&b.name)
		fmt.Print("Enter Book author : ")
		fmt.Scan(&b.author)
		books[b.id]=b
	}
	var searchID int
	fmt.Print("\nEnter the Book ID to display: ")
	fmt.Scan(&searchID)

	book, exists := books[searchID]
	if exists {
		fmt.Println("\nBook Details:")
		fmt.Println("ID:", book.id)
		fmt.Println("Name:", book.name)
		fmt.Println("Author:", book.author)
	} else {
		fmt.Println("Book not found with ID", searchID)
	}

}
