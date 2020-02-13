package sqlop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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
	fmt.Println(id)
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
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func DeleteRow(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	if err != nil {
		log.Fatal(err)
	}
}
