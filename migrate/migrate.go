package main

import (
	"timebox.com/initializers"
	"timebox.com/models"
	"log"
)

 func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
 }

 func main() {
		initializers.DB.AutoMigrate(&models.Project{})
		log.Printf("Migration Completed")
 }