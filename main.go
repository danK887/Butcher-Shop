package main

import (
	"html/template"
	"log"
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
<<<<<<< HEAD
	//userBob := User{"Bob", "+7-123-456-7890", "bob@gmail.com"}
	fmt.Fprintf(w, "contact page")
	//fmt.Fprintf(w, "User data: %s %s %s", userBob.name, userBob.phone, userBob.email)
=======
	userBob := User{"Bob", "+7-123-456-7890", "bob@gmail.com"}
	tmpl, _ := template.ParseFiles("templates/contact.html")
	tmpl.Execute(w, userBob)
}
func chefsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/stuff.html")
	tmpl.Execute(w, nil)
}
func galleryPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/gallery.html")
	tmpl.Execute(w, nil)
>>>>>>> template
}

func handleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/menu/", menuPage)
	mux.HandleFunc("/about/", aboutPage)
	mux.HandleFunc("/contacts/", contactsPage)
	mux.HandleFunc("/chefs/", chefsPage)
	mux.HandleFunc("/gallery/", galleryPage)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func main() {

	handleRequest()

}
