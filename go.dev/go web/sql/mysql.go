package sql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// start mysql
func StartMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/demo?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTable(db *sql.DB) {
	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			password TEXT NOT NULL,
			username TEXT NOT NULL,
			department TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func InsertRow(db *sql.DB) {
	username := "Jack"
	password := 123456
	createdAt := time.Now()
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`,
		username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func InsertRow2(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	CheckErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Println("id:", id)

}

func QueryRow(db *sql.DB) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, username, password, createdAt)
}

func QueryRows(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)

		fmt.Println(users)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func QueryAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)

		CheckErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

}

func DeleteRow(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
