package main

import (
	"fmt"
	"net/http"
	//"html/template"
)

type User struct {
	name  string
	phone string
	email string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the jungle!")
}
func menuPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "menu page")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
func contactsPage(w http.ResponseWriter, r *http.Request) {
	//userBob := User{"Bob", "+7-123-456-7890", "bob@gmail.com"}
	fmt.Fprintf(w, "contact page")
	//fmt.Fprintf(w, "User data: %s %s %s", userBob.name, userBob.phone, userBob.email)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/menu/", menuPage)
	http.HandleFunc("/about/", aboutPage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()

}
