package models

import (
	"gin_notes/helpers"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserCheckAvailability(email string) bool {
	var user User
	DB.Where("username = ?", email).First(&user)
	return (user.ID == 0) //if ID == 0 email is not signed up, hence avalable
}

func UserCreate(email string, password string) *User {
	hashPasswd, _ := helpers.HashPassword(password)

	entry := User{Username: email, Password: hashPasswd}
	DB.Create(&entry)
	return &entry
}
func UserFind(id uint64) *User {
	var user User
	DB.Where("id = ?", id).First(&user)
	return &user
}

func UserCheck(email string, password string) *User {
	var user User

	DB.Where("username = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}

	match := helpers.CheckPasswordHash(password, user.Password)

	if match {
		return &user
	} else {
		return nil
	}
}
