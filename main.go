package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

type CustomerReviews struct {
	Id                               uint16
	Name, Email, SiteRating, Content string
}

var reviews = []CustomerReviews{}

// подключение к бд, для получения данных о фитбеке пользователей
func dataFromDB() []CustomerReviews {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/feetback")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `usersdata2`")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	// проходит по таблице и забираем все данные из указанных полей
	reviews = []CustomerReviews{}
	for res.Next() {
		var review CustomerReviews
		err := res.Scan(&review.Id, &review.Name, &review.Email, &review.SiteRating, &review.Content)
		if err != nil {
			panic(err)
		}
		reviews = append(reviews, review)
	}
	return reviews

}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/customerReviews.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", dataFromDB())

}
func menuPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/menu.html", "templates/footer.html", "templates/customerReviews.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "menu", dataFromDB())
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.Execute(w, nil)
}
func contactsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/contact.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "contact", nil)
}
func chefsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/stuff.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.Execute(w, nil)
}
func galleryPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/gallery.html", "templates/footer.html", "templates/customerReviews.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "gallery", dataFromDB())
}

func saveFeetback(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	siteRating := r.FormValue("siterating")
	content := r.FormValue("content")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/feetback")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new user data
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `usersdata2` (`name`, `email`, `siteRating`, `content`) VALUES ('%s', '%s', '%s', '%s')", name, email, siteRating, content))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/menu/", menuPage)
	mux.HandleFunc("/about/", aboutPage)
	mux.HandleFunc("/contacts/", contactsPage)
	mux.HandleFunc("/chefs/", chefsPage)
	mux.HandleFunc("/gallery/", galleryPage)
	mux.HandleFunc("/saveFeetback/", saveFeetback)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func main() {

	handleRequest()

}
