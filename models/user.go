package models

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type user User
	safeUser := struct {
		user
		Password string `json:"password,omitempty"`
	}{
		user: user(*u),
	}
	return json.Marshal(safeUser)
}

func (u *User) CreateUser() (err error) {
	err = u.GenerateFromPassword(u.Password)
	if err != nil {
		return err
	}
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GenerateFromPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CompareHashAndPassword(password string) (err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (user *User, err error) {
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
