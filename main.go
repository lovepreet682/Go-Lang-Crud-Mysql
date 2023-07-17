package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// create a DB
var DB *gorm.DB

var temp *template.Template

type crud struct {
	id    string
	Name  string
	email string
	age   string
	city  string
}

func getMySQLDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:rehan200@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection status
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database.")
	} else {
		fmt.Println("Connected to the database successfully!")
	}
	return db
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new.html", nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "show.html", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "edit.html", nil)
}

func init() {
	temp = template.Must(template.ParseGlob("templates/*.html"))
}

//var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Insert the Value
func Insert(w http.ResponseWriter, r *http.Request) {
	db := getMySQLDB()

	if r.Method == "POST" {

		name := r.FormValue("name")
		email := r.FormValue("email")
		age := r.FormValue("age")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO crud (name, email, age, city) VALUES (?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}

		//Exceute the Query using form
		insForm.Exec(name, email, age, city)
		log.Println("Insert Data: name " + name + " | Email " + email + " | Age " + age + " | City " + city)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// // show the Data
// func Show(w http.ResponseWriter, r *http.Request) {
// 	db := getMySQLDB()
// 	selectForDB, err := db.Query("select * from curd")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	emp := crud{}
// 	res := []crud{}
// 	for selectForDB.Next() {
// 		var id string
// 		var name string
// 		var email string
// 		var age string
// 		var city string

// 		err = selectForDB.Scan(&id, &name, &email, &age, &city)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		emp.id = id
// 		emp.Name = name
// 		emp.email = email
// 		emp.age = age
// 		emp.city = city

// 		res = append(res, emp)
// 	}

// 	// Create an HTML template
// 	// tmpl.ExecuteTemplate(w, "index", res)
// 	defer db.Close()
// }

func main() {

	getMySQLDB()

	r := mux.NewRouter()
	r.HandleFunc("/insert", Insert)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/show", showHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/edit", editHandler)

	//action in the
	r.HandleFunc("/insert", Insert) // INSERT :: New register
	http.ListenAndServe(":9090", r)
}
