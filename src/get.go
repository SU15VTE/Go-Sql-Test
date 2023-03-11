package main

import (
	"fmt"
	"log"
)

func GetRows(id int) {
	Get, err := selectRow(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("id: %d, username: %s, password: %s\n", Get.id, Get.username, Get.password)
}
func GetAll(id int) {
	GetAll, err := selecetAll(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(GetAll)
	fmt.Printf("%s     %8s     %s\n", idstr, usernamestr, passwordstr)
	for _, users := range GetAll {
		fmt.Printf("%d     %10s     %3s\n", users.id, users.username, users.password)
	}
}

func prepareGet(id int) {
	Get, err := prepareSelectRow(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("%s     %8s     %s\n", idstr, usernamestr, passwordstr)
	for _, users := range Get {
		fmt.Printf("%d     %10s     %3s\n", users.id, users.username, users.password)
	}
}
func prepareGetAll(id int) {
	GetAll, err := prepareSelectAll(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(GetAll)
	fmt.Printf("%s     %8s     %s\n", idstr, usernamestr, passwordstr)
	for _, users := range GetAll {
		fmt.Printf("%d     %10s     %3s\n", users.id, users.username, users.password)
	}
}
