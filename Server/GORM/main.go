package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Connection
var dsn = "root:password@tcp(127.0.0.1:3306)/test_practice?charset=utf8mb4"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type GoTestModel struct {
	Name string
	Year string
}

func main() {
	http.HandleFunc("/createstuff", GoDatabaseCreate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	GoTestModel := GoTestModel{
		Name: "Mike",
		Year: "2022",
	}
	db.Create(&GoTestModel)
	if err := db.Create(&GoTestModel).Error; err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(GoTestModel)

	fmt.Println("Fields Added", GoTestModel)
}
