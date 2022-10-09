package main

// import "fmt"
import (
	"github.com/gin-gonic/gin"
	"timebox.com/initializers"
	"timebox.com/controllers"
)



func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/projects", controllers.ProjectsCreate)
	r.GET("/projects", controllers.ProjectsIndex)
	r.GET("/projects/:id", controllers.ProjectsShow)
	r.PUT("/projects/:id", controllers.ProjectsUpdate)
	r.DELETE("/projects/:id", controllers.ProjectsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
} 