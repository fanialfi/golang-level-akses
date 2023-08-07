package main

import (
	"fmt"
	"golang-level-akses/library"
)

func main() {
	fmt.Printf("Name : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
func init() {
	fmt.Println("Apakah yang pertama")
	// sayHello("fani")
}
