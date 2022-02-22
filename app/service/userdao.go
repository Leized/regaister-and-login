package service

import (
	"Project/app/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func QueryName(username string) model.Users {
	user := model.Users{}
	db, err := gorm.Open(mysql.Open("root:990820@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Where("username = ?", username).First(&user)
	return user
}

func InsertData(username string, password string) model.Users {
	user := model.Users{}
	db, err := gorm.Open(mysql.Open("root:990820@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&model.Users{
		Username: username,
		Password: password,
	})
	return user
}
func Encryption(password string) model.Users {
	user := model.Users{}
	salt := time.Now().Unix()
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	encryption := hex.EncodeToString(st)
	user.Password = encryption
	return user
}
func Encryption1(password string) model.RegisterReq {
	req := model.RegisterReq{}
	salt := time.Now().Unix()
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	encryption := hex.EncodeToString(st)
	req.Password = encryption
	return req
}
