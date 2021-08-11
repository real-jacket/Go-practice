package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-dev/web-server/middleware"
	sqlop "github.com/go-dev/web-server/mysql"
	nbcrypt "github.com/go-dev/web-server/passwordHashing"
	"github.com/go-dev/web-server/session"
	ptemplate "github.com/go-dev/web-server/template"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie!")

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
	}
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./views/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	db := sqlop.StartMysql()
	//sqlop.CreateTable(db)
	r := mux.NewRouter()

	//Todo page
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/form.html"))
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.ExecuteTemplate(w, "index", ptemplate.TempTodoPage())
	// })

	r.HandleFunc("/", sayHelloName)
	r.HandleFunc("/login", login)
	r.HandleFunc("/upload", upload)

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

	//Use websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	r.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/websocket.html")
	})
	r.HandleFunc("/websocket/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	// passwordHashing
	password := "secret"
	hash, _ := nbcrypt.HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password", password)
	fmt.Println("Hash: ", hash)

	match := nbcrypt.CheckPasswordHash(password, hash)
	fmt.Println("Match", match)

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

	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
