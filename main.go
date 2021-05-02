package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect mysql failed, err: %v\n", err)
		return
	}
	defer db.Close()
	// 模型绑定
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	// 设定静态文件路径，前后端分离不需要
	r.Static("/static", "static")
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1g := r.Group("/v1")
	{
		// 添加
		v1g.POST("/todo", func(c *gin.Context) {
			// 请求中获取参数
			var todo Todo
			c.BindJSON(&todo)
			// 存入数据库
			if err = db.Create(&todo).Error; err == nil {
				c.JSON(http.StatusOK, todo)
				/*c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "success",
					"data": todo,
				})*/
			} else {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		})
		// 查看所有待办事项
		v1g.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			if err = db.Find(&todos).Error; err == nil {
				c.JSON(http.StatusOK, todos)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		})
		// 查看某个待办事项
		v1g.GET("/todo/:id", func(c *gin.Context) {
		})
		// 修改状态
		v1g.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			if err = db.Where("id=?", id).Find(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			if todo.Status {
				todo.Status = false
			} else {
				todo.Status = true
			}
			db.Model(&todo).Update("status", todo.Status)
			c.JSON(http.StatusOK, todo)

		})
		// 删除
		v1g.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			if err = db.Where("id=?", id).Delete(Todo{}).Error; err == nil {
				c.JSON(http.StatusOK,gin.H{
					"msg":"success",
				})
			}else{
				c.JSON(http.StatusOK,gin.H{
					"error":err.Error(),
				})
			}
		})
	}

	r.Run(":8888")
}
