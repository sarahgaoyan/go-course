package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sarahgaoyan/go-course/10-todolist/model"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddToDo(c *gin.Context) {
	var todo model.ToDo
	c.BindJSON(&todo)
	err := model.CreateAToDo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "create success",
			"data": todo,
		})
	}
}

func GetAll(c *gin.Context) {
	var todoList []model.ToDo
	if err := model.GetAllList(&todoList); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateAToDo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "id not exist"})
	}
	var todo model.ToDo
	if err := model.FindAToDoById(id, &todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	ctx.BindJSON(&todo)
	if err := model.SaveAToDo(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteAToDo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "id not exist"})
	}
	if err := model.DeleteAToDoById(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id": "deleted"})
	}
}
