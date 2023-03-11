package github.com/SU15VTE/Go-Sql-Test/src/

import ("log"
		"src")

func transaction(username string) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() // rollback
		}
		log.Fatalln(err)
		return
	}

	sqlStr := "UPDATE `user` SET `username` =? WHERE `id` =?"
	ret, err := db.Exec(sqlStr, username)
	if err != nil {
		tx.Rollback() // rollback
		log.Fatalln(err)
		return
	}

}
