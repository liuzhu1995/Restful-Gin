package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// 使用bcrypt算法生成密码的哈希值
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


func VerifyPassword(hashPassword, password string) bool {
	// 使用bcrypt算法验证密码是否匹配哈希值
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
 
	return err == nil
}