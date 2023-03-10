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
	GetRows(10)
}

func GetRows(id int) {
	Get, err := selectRow(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", Get.id, Get.name, Get.age)
}
func GetAll(id int) {
	GetAll, err := selecetAll(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(GetAll)
	fmt.Printf("%s     %8s     %s\n", idstr, namestr, agestr)
	for _, users := range GetAll {
		fmt.Printf("%d     %10s     %3d\n", users.id, users.name, users.age)
	}
}

func prepareGet(id int) {
	Get, err := prepareSelectRow(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", Get.id, Get.name, Get.age)
}

func prepareGetAll(id int) {
	GetAll, err := prepareSelectAll(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(GetAll)
	fmt.Printf("%s     %8s     %s\n", idstr, namestr, agestr)
	for _, users := range GetAll {
		fmt.Printf("%d     %10s     %3d\n", users.id, users.name, users.age)
	}
}
