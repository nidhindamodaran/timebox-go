package main

// import "fmt"
import (
	"github.com/gin-gonic/gin"
	"timebox.com/initializers"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", pingHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
} 