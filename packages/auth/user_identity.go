package auth

import (
	"crypto/md5"
	"encoding/hex"
)

// UserIdentityData структура пользовательских данных для авторизации
type UserIdentityData struct {
	ID       int    // ID номер записи в Бд
	Login    string `json:"login"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// TableName название таблицы БД
func (u UserIdentityData) TableName() string {
	return "identity"
}

// PasswordToMd5 пароль в MD5
func (u UserIdentityData) PasswordToMd5() string {
	passByte := []byte(u.Password)
	passwordHash := md5.Sum(passByte)
	passString := hex.EncodeToString(passwordHash[:])

	return passString
}
