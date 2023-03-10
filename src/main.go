package main

import (
	"database/sql"
	"fmt"
	"log"

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
	GetRows(2)
	GetAll(0)

	updateRow(6, "mio", 16)
	deleteRow(5)
	GetAll(0)
}

func GetRows(id int) {
	Get, err := selectRow(2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", Get.id, Get.name, Get.age)
}
func GetAll(id int) {
	GetAll, err := selecetAll(0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(GetAll)
	fmt.Printf("%s     %8s     %s\n", idstr, namestr, agestr)
	for _, users := range GetAll {
		fmt.Printf("%d     %10s     %3d\n", users.id, users.name, users.age)
	}
}
