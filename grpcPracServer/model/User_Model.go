package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Player int
}

func Login(username ,password string)error{
	var u User
	//DB.Where("username = ?", f.Username).Where("password = ?", f.Password).First(&u)
	DB.Where(User{
		Username: username,
		Password: password,
	}).First(&u)

	if u.ID == 0 {
		return errors.New("user name or password error.")
	} else {
		return nil
	}
}

func Register(username ,password string)error{
	var u User
	DB.Where("username = ?",username).First(&u)
	if u.ID!=0 {
		return errors.New("username exist!")
	}

	u = User{
		Username: username,
		Password: password,
		Player: 0,
	}

	DB.Create(&u)
	return nil
}


