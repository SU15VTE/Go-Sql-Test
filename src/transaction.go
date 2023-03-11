package main

import (
	"fmt"
	"log"
)

func transaction(username string, id int, password string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() // rollback
			log.Fatalln(err)
		}
		return
	}

	sqlStr := "UPDATE `user` SET `username` =? WHERE `id` =?"
	ret, err := tx.Exec(sqlStr, username, id)
	if err != nil {
		tx.Rollback() // rollback
		log.Fatalln(err)
		return
	}
	Row, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // rollback
		log.Fatalln(err)
		return
	}
	sqlStr2 := "UPDATE `user` SET `password` =? WHERE `id` =?"
	ret2, err := tx.Exec(sqlStr2, password, id)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
		return
	}
	Row2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // rollback
		log.Fatalln(err)
		return
	}

	if Row == 1 && Row2 == 1 {
		fmt.Println("transaction success")
		tx.Commit() //Commit transaction
	} else {
		tx.Rollback()
		fmt.Println("transaction failed,Rollback transaction!")
	}
	return

}
