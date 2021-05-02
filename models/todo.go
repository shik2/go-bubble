package models

import (
	"go-web-demo/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

func GetTodoList() (todoList []Todo, err error) {
	var todos []Todo
	err = dao.DB.Find(&todos).Error
	return todos, err
}

func FindTodoByID(id string) (todo Todo, err error) {
	err = dao.DB.Where("id=?", id).Find(&todo).Error
	return todo, err
}

func UpdateTodoStatus(todo Todo) (err error) {
	if todo.Status {
		todo.Status = false
	} else {
		todo.Status = true
	}
	err = dao.DB.Model(&todo).Update("status", todo.Status).Error
	return err
}

func DeleteTodoByID(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return err
}
