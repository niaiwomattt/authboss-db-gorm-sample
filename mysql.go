package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func openDb() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", "dev:123@/gyys_authboss?charset=utf8mb4") // &parseTime=True&loc=Local
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 全局禁用表名复数
	db.SingularTable(true)

	return db
}



func initCreate() {
	db.AutoMigrate(&User{})
}