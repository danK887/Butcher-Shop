package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name  string
	Phone string
	Email string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)

}
func menuPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/menu.html")
	tmpl.Execute(w, nil)
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/about.html")
	tmpl.Execute(w, nil)
}
func contactsPage(w http.ResponseWriter, r *http.Request) {
	userBob := User{"Bob", "+7-123-456-7890", "bob@gmail.com"}
	tmpl, _ := template.ParseFiles("templates/contact.html")
	tmpl.Execute(w, userBob)
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
