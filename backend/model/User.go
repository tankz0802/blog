package model

import (
	error_msg "blog/error"
	"blog/utils"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" from:"username" binding:"required,UsernameValidator" json:"username"`
	Password string `gorm:"type:varchar(64);not null" from:"password" binding:"required" json:"password"`
	Tel string `gorm:"type:char(11);not null" from:"tel" binding:"required,TelValidator" json:"tel"`
	Email string `gorm:"type:varchar(20);not null" from:"email" binding:"required,email" json:"email"`
	Avatar string `gorm:"type:varchar(255)" from:"avatar" json:"avatar"`
	Salt string `gorm:"type:char(8)"`
}

type UpdateUserData struct {
	Username string `from:"username" json:"username" binding:"required,UsernameValidator"`
	Tel string `from:"tel" json:"tel" binding:"required,TelValidator"`
	Email string `from:"email" json:"email" binding:"required,email"`
}

func UsernameIsExist(id uint, username string) bool {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID != id {
		return true
	}
	return false
}

func UserIsExist(username string) bool {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func (u *User) BeforeCreate() {
	salt := utils.GenSalt()
	u.Salt = string(salt)
	pwd, err := scrypt.Key([]byte(u.Password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = base64.StdEncoding.EncodeToString(pwd)
}

func CreateUser(user *User) (code int) {
	if err := db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return error_msg.ERROR
	}
	return error_msg.SUCCESS
}

func CheckPwd(username,password string) (*User, int) {
	var u User
	db.Where("username = ?", username).First(&u)
	pwd, err := scrypt.Key([]byte(password), []byte(u.Salt), 1<<15, 8, 1, 32)
	if err != nil {
		return nil, error_msg.ERROR
	}
	if base64.StdEncoding.EncodeToString(pwd) != u.Password {
		return nil, error_msg.ERROR_PASSWORD_WRONG
	}
	return &u, error_msg.SUCCESS
}

func UpdateUser(id int, data *UpdateUserData) int {
	var user User
	m := make(map[string]interface{})
	m["username"] = data.Username
	m["tel"] = data.Tel
	m["email"] = user.Email
	err := db.Model(&user).Where("id = ?", id).Updates(m).Error
	if err != nil {
		log.Println(err.Error())
		return error_msg.ERROR
	}
	return error_msg.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := db.Delete(&user, id).Error
	if err != nil {
		log.Println(err)
		return error_msg.ERROR
	}
	return error_msg.SUCCESS
}