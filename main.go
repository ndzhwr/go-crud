package main

import (
	"go-crud/controllers"
	"go-crud/utils"

	"github.com/gin-gonic/gin"
)

/*
Loads environment variables and connects
to database using the utils package
*/
func Init() {
	utils.LoadEnv()
	utils.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.MainRouteController)
	r.GET("/posts", controllers.AllPosts)
	r.GET("/posts/:id", controllers.SinglePost)
	r.POST("/posts/new", controllers.PostCreate)
	r.PATCH("/posts/update/:id", controllers.UpdatePost)
	r.DELETE("/posts/delete/:id", controllers.DeletePost)
	Init()

	r.Run()
}
