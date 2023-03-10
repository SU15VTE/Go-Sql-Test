package main

import (
	"fmt"
	"log"
)

//
func selectRow(id int) (u User, err error) {
	u = User{}
	selectStr := "SELECT * FROM user WHERE id =?"
	err = db.QueryRow(selectStr, id).Scan(&u.id, &u.name, &u.age)
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
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Fatalln(err.Error())
		}
		users = append(users, u)
	}
	return
}

func insertRow(name string, age int) (err error) {
	insertStr := "INSERT INTO user (name, age) VALUES (?,?)"
	ret, err := db.Exec(insertStr, name, age)
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

func updateRow(id int, name string, age int) (err error) {
	sqlStr := "UPDATE user SET name =?, age =? WHERE id =?"
	ret, err := db.Exec(sqlStr, name, age, id)
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
