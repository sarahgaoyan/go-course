package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func main() {
	DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	DB.AutoMigrate(&ToDo{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", func(ctx *gin.Context) {
			var todo ToDo
			ctx.BindJSON(&todo)
			if err := DB.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"msg":  "create success",
					"data": todo,
				})
			}
		})
		v1Group.GET("/todo", func(ctx *gin.Context) {
			var todoList []ToDo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}

		})
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "id not exist"})
			}
			var todo ToDo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			ctx.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}
		})
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "id not exist"})
			}
			if err = DB.Where("id=?", id).Delete(ToDo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"id": "deleted"})
			}

		})
	}

	r.Run()
}
