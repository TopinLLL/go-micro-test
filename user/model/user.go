package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PasswordDigest string
}
const(
	PASSWORDCOST  = 12
)

//加密密码
func (user *User)EncodePassword(password string)string{
	encodePassword, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORDCOST)
	if err!=nil{
		fmt.Println("密码加密出错",err)
	}
	return string(encodePassword)
}


func (user *User)DecodePassword(password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.PasswordDigest))
	return err==nil
}


