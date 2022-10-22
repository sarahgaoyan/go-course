package model

import (
	"github.com/sarahgaoyan/go-course/10-todolist/dao"
)

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateAToDo(todo *ToDo) (err error) {
	return dao.DB.Create(&todo).Error
}

func GetAllList(todoList *[]ToDo) (err error) {
	return dao.DB.Find(&todoList).Error
}

func FindAToDoById(id string, todo *ToDo) (err error) {
	return dao.DB.Where("id=?", id).First(&todo).Error
}

func SaveAToDo(todo *ToDo) (err error) {
	return dao.DB.Save(&todo).Error
}

func DeleteAToDoById(id string) (err error) {
	return dao.DB.Where("id=?", id).Delete(ToDo{}).Error
}
