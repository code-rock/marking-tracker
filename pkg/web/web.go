package web

import (
	"fmt"
	"net/http"
	"html/template"
	// "webcam"
	// "gocv.io/x/gocv"
	// "database/sql"
	// "github.com/go-sql-driver/mysql"
)

type Marking struct {
	Name, Folder, Id, Description string
	Avg_price uint16
	Images []string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	mark := Marking{"with urn", "./man-with-urn", "man-with-urn", "не мусорите", 3, []string{"1.png", "2.png","sharp-image.png"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, mark)
}


func search_page(w http.ResponseWriter, r *http.Request) {
	mark := Marking{"with urn", "./man-with-urn", "man-with-urn", "не мусорите", 3, []string{"1.png", "2.png","sharp-image.png"}}
	tmpl, _ := template.ParseFiles("templates/search_page.html")
    tmpl.Execute(w, mark)
}

func (m *Marking) getAllInfo() string {
	return fmt.Sprintf("user name is %s." +
	"It costs %d. %s", m.Name, m.Avg_price, m.Description)
}
func (m *Marking) setNewName(newName string) {
	m.Name = newName
}
func markings_page(w http.ResponseWriter, r *http.Request) {
	mark := Marking{"with urn", "./man-with-urn", "man-with-urn", "не мусорите", 3, []string{"1.png", "2.png","sharp-image.png"}}
	tmpl, _ := template.ParseFiles("templates/markings_page.html")
	tmpl.Execute(w, mark)
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/search/", search_page)
	http.HandleFunc("/markings/", markings_page)
	http.ListenAndServe(":3001", nil)
}

func Listen() {
	handleRequest()
}