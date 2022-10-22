package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sarahgaoyan/go-course/10-todolist/controller"
)

func Register() (r *gin.Engine) {
	r = gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "static")

	r.GET("/", controller.Home)

	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", controller.AddToDo)
		v1Group.GET("/todo", controller.GetAll)
		v1Group.PUT("/todo/:id", controller.UpdateAToDo)
		v1Group.DELETE("/todo/:id", controller.DeleteAToDo)
	}
	return
}
