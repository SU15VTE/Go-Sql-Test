package main

func selectall(id int) (u User, err error) {
	u = User{}
	selectStr := "SELECT * FROM user WHERE id =?"
	err = db.QueryRow(selectStr, 1).Scan(&u.id, &u.name, &u.age)
	return
}
