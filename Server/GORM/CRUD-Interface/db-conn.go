package main

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "root:password@tcp(127.0.0.1:3306)/test_practice?charset=utf8mb4&parseTime=true&loc=Local"

func DBConn() {
	sqlDB, err := sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println("Error", err)
	}

	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with sqlDB. Ping ")
		panic(err.Error())
	}
	insert, err := sqlDB.Query("INSERT INTO `test_practice`.`param`(`id`, `first_name`, `last_name`, `age`, `address`) VALUES('2', 'Ben', 'Ford');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successful Connnection to Database")
	// gormDB, err2 := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	// if err2 != nil {
	// 	fmt.Println("Error", err2)
	// }
	// fmt.Println("gormDB Existing:", *gormDB)
}
