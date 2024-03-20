package models

import (
	"fmt"

	"example.com/rest-api/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users (email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("UserSaveFunc预编译语句失败:%v", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("UserSaveFunc创建用户失败:%v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}


func GetAllUsers()([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}