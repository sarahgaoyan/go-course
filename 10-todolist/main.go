package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sarahgaoyan/go-course/10-todolist/dao"
	"github.com/sarahgaoyan/go-course/10-todolist/model"
	"github.com/sarahgaoyan/go-course/10-todolist/router"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	dao.DB.AutoMigrate(&model.ToDo{})

	r := router.Register()
	r.Run()
}
