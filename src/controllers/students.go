package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/psnavega/api-go-gin/database"
	"github.com/psnavega/api-go-gin/src/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "Id não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "E aí, tudo beleza, " + name + "?",
	})
}

func CreateStudents(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, &student)
}
