package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	sqlDB, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test_practice?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error", err)
	}
	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err2 != nil {
		fmt.Println("Error", err2)
	}
	fmt.Println("gormDB Existing:", *gormDB)

}
