package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web-demo/dao"
	"go-web-demo/models"
	"go-web-demo/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SerUpRouter()

	r.Run(":8888")
}
