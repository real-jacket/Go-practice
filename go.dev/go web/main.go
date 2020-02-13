package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-dev/web-server/middleware"
	sqlop "github.com/go-dev/web-server/mysql"
	"github.com/go-dev/web-server/session"
	ptemplate "github.com/go-dev/web-server/template"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	db := sqlop.StartMysql()
	//sqlop.CreateTable(db)
	r := mux.NewRouter()

	//Todo page
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/form.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", ptemplate.TempTodoPage())
	})

	//Form page
	r.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.ExecuteTemplate(w, "form", nil)
			return
		}

		tmpl.ExecuteTemplate(w, "form", struct {
			Success bool
		}{true})

		details := ptemplate.ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		fmt.Println(details)
	})

	// use middleware
	r.HandleFunc("/middleware", middleware.Chain(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	}, middleware.Logging(), middleware.Method("GET")))

	// use session and cookie
	s := r.PathPrefix("/session").Subrouter()
	s.HandleFunc("/secret", session.Secret)
	s.HandleFunc("/login", session.Login)
	s.HandleFunc("/logout", session.Logout)

	// use  json
	type User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	}

	r.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
	})

	r.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			FirstName: "John",
			LastName:  "Doe",
			Age:       25,
		}
		json.NewEncoder(w).Encode(peter)
	})

	// books 请求示例
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// sql 操作
	r.HandleFunc("/mysql/table/insert", func(w http.ResponseWriter, r *http.Request) {
		sqlop.InsertRow(db)
		fmt.Fprintf(w, "Add one row")
	})

	//  server 静态文件目录
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}
