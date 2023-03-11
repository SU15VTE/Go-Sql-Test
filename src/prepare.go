package main

import (
	"fmt"
	"log"
)

func prepareSelectRow(id int) (results []User, err error) {
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
		var user User
		err = rows.Scan(&user.id, &user.username, &user.password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		results = append(results, user)
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
		err = rows.Scan(&user.id, &user.username, &user.password)
		if err != nil {
			fmt.Printf("scan error: %v\n", err)
			return
		}
		users = append(users, user)

	}
	return
}

func prepareInsert(username string, password string) (err error) {
	sqlStr := "INSERT INTO USER(username,password) VALUES(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln()
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Fatalln()
	}
	fmt.Println("Insert success.")
	return
}

func prepareDelete(id int) (err error) {
	sqlStr := "DELETE FROM user WHERE id =?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln()
	}
	defer stmt.Close()
	rows, err := stmt.Exec(id)
	if err != nil {
		log.Fatalln()
	}
	n, err := rows.RowsAffected()
	if err != nil {
		log.Fatalln()
	}
	fmt.Printf("Delete success,n = %d\n", n)
	return
}

func prepareupdate(id int, username string, password string) (err error) {
	sqlStr := "UPDATE user SET username =?, password =? WHERE id =?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln()
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, password, id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatalln()
	}
	fmt.Printf("update success,n= %d \n", n)
	return
}
