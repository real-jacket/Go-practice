package main

import (
	"fmt"
	sqlop "github.com/go-dev/web-server/mysql"
	palmate "github.com/go-dev/web-server/template"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	db := sqlop.StartMysql()
	//sqlop.CreateTable(db)
	r := mux.NewRouter()

	//Todo page
	tmpl := template.Must(template.ParseFiles("layout.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := palmate.TempTodoPage()
		tmpl.Execute(w, data)
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

	http.ListenAndServe(":80", r)
}
