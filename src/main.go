package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go-db?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := Init()
	if err != nil {
		fmt.Println("Init database error: ", err)
		return
	} else {
		fmt.Println("Init database success")
	}
	transaction("SU15VTE", 1, "123456")
	GetRows(1)
}
