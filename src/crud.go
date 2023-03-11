package main

import (
	"fmt"
	"log"
)

//
func selectRow(id int) (u User, err error) {
	u = User{}
	selectStr := "SELECT * FROM user WHERE id =?"
	err = db.QueryRow(selectStr, id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

//
func selecetAll(id int) (users []User, err error) {
	selectStr := "SELECT * FROM user WHERE id >?"
	rows, err := db.Query(selectStr, id)

	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		users = append(users, u)
	}
	return
}

func insertRow(username string, password string) (err error) {
	insertStr := "INSERT INTO user (username, password) VALUES (?,?)"
	ret, err := db.Exec(insertStr, username, password)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	thisID, err := ret.LastInsertId()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("insert success, thisID = %d\n", thisID)
	return
}

func deleteRow(id int) (err error) {
	deleteStr := "DELETE FROM user WHERE id =?"
	ret, err := db.Exec(deleteStr, id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("delete success, n = %d\n", n)
	return
}

func updateRow(id int, username string, password string) (err error) {
	sqlStr := "UPDATE user SET username =?, password =? WHERE id =?"
	ret, err := db.Exec(sqlStr, username, password, id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("update success, n = %d\n", n)
	return
}
