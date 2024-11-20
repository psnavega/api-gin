package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psnavega/api-go-gin/src/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/student/:id", controllers.GetStudent)
	r.GET("/greetings/:name", controllers.Greetings)
	r.POST("/students", controllers.CreateStudents)
	r.Run(":8080")
}
