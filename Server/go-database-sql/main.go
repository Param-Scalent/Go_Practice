package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		id         int
		first_name string
	)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test_practice")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error Connecting Database", err)
	}
	// Preapring Queries => stmt, err := db.Prepare("select id, name from users where id = ?")
	// Passing value as placeholder or bind parameter => rows, err := stmt.Query(1)
	rows, err := db.Query("select id, first_name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &first_name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, first_name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// Single-Row Queries
	err = db.QueryRow("select first_name from users where id = ?", 1).Scan(&first_name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(first_name)
	// On Prepared statements
	// stmt, err := db.Prepare("select first_name from users where id = ?")
	// err = stmt.QueryRow(1).Scan(&first_name)

	// Statements that Modify Data
	stmt, err := db.Prepare("insert into users(first_name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowcnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %d affected = %d \n", lastId, rowcnt)

	// Working with Transactions
}
