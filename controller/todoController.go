package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/models"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	// 请求中获取参数
	var todo models.Todo
	c.BindJSON(&todo)
	// 存入数据库
	err := models.CreateATodo(&todo)
	if err == nil {
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
}

func GetTodoList(c *gin.Context) {
	todos, err := models.GetTodoList()
	if err == nil {
		c.JSON(http.StatusOK, todos)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdateStatus(c *gin.Context) {
	id, _ := c.Params.Get("id")
	todo, err := models.FindTodoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.UpdateTodoStatus(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	err := models.DeleteTodoByID(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
}
