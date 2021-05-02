package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/controller"
)

func SerUpRouter() *gin.Engine {
	r := gin.Default()
	// 设定静态文件路径，前后端分离不需要
	r.Static("/static", "static")
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", controller.IndexHandler)

	v1g := r.Group("/v1")
	{
		// 添加
		v1g.POST("/todo", controller.CreateTodo)
		// 查看所有待办事项
		v1g.GET("/todo", controller.GetTodoList)
		// 查看某个待办事项
		v1g.GET("/todo/:id", func(c *gin.Context) {
		})
		// 修改状态
		v1g.PUT("/todo/:id", controller.UpdateStatus)
		// 删除
		v1g.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
