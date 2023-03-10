package main

import (
	"fmt"
	"log"
)

func prepareSelectRow(id int) (user User, err error) {
	sqlStr := "SELECT * FROM user WHERE id =?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			log.Fatalln(err.Error())
		}
		return
	}
	return
}

func prepareSelectAll(id int) (users []User, err error) {
	sqlStr := "SELECT * FROM user WHERE id >?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err != nil {
			fmt.Printf("scan error: %v\n", err)
			return
		}
		err = rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			fmt.Printf("scan error: %v\n", err)
			return
		}
		users = append(users, user)

	}
	return
}
