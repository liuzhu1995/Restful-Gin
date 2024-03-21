package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
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
		return fmt.Errorf("预编译语句失败:%v", err)
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("生成密码哈希值失败:%v", err)
	}

	result, err := stmt.Exec(user.Email, hashPassword)
	if err != nil {
		return fmt.Errorf("创建用户失败:%v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	user.Password = string(hashPassword)
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

func (u *User) ValidateCredentials() error {
	// 通过email查找密码和id，验证密码是否匹配
	query := "SELECT id, password FROM users WHERE email = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	var retrievedPassword string
	row := stmt.QueryRow(u.Email)
	err = row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}
	passwordIsValid := utils.VerifyPassword(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
 

func GetUserByID(userID int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var user User
	row := stmt.QueryRow(userID)
	err = row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) DeleteUser() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID)

	return err
}