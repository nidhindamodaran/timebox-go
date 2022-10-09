package controllers

import (
	"github.com/gin-gonic/gin"
	"timebox.com/models"
	"timebox.com/initializers"
	"time"
)

func ProjectsCreate(c *gin.Context) {
	

	var content struct {
		Name string
		Icon string
		Deadline *time.Time
	}

	c.Bind(&content) 

	project := models.Project {Name: content.Name, Icon: content.Icon, Deadline: content.Deadline}
	result := initializers.DB.Create(&project)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})
}

func ProjectsIndex(c *gin.Context) {
	var projects []models.Project
	result := initializers.DB.Find(&projects)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func ProjectsShow(c *gin.Context) {
	var project models.Project

	id := c.Param("id")
 
	result := initializers.DB.Find(&project, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})
}

func ProjectsUpdate(c *gin.Context) {
	id := c.Param("id")
	
	var project models.Project
	initializers.DB.Find(&project, id)
	
	var content struct {
		Name string
		Icon string
		Hidden bool
		Deadline *time.Time
	}
	c.Bind(&content)

	result := initializers.DB.Model(&project).Updates(models.Project{
		Name: content.Name,
		Icon: content.Icon,
		Hidden: content.Hidden,
		Deadline: content.Deadline,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})
}

func ProjectsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Project{}, id)

	c.Status(200)
}