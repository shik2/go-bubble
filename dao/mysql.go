package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	// 连接数据库
	dsn := "root:123456@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect mysql failed, err: %v\n", err)
		return
	}
	return DB.DB().Ping()
}
