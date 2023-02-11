package main

import (
	"fmt"
	"github.com/huangyuuu1999/oranget/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	TestMysql()
}

func TestMysql() {
	username := "root"
	password := "123456"
	host := "43.139.176.247"
	port := 3000
	Dbname := "fruit_buckets"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db.AutoMigrate(&model.UserBasic{})

	user := model.UserBasic{}
	user.Username = "Qxx"
	user.Password = "123456"
	//db.Create(&user)
	//db.First(&user)
	//fmt.Println(user)

	db.Model(&user).Where("id=1").Update("password", "654321")
	db.First(&user)
	fmt.Println(user)
}
