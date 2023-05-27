package routes

import (
	"github.com/JuanCampbsi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ViewStudents)
	r.GET("/students/:id", controllers.SearchStudentsId)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentsCPF)
	r.POST("/students", controllers.CreateStudents)
	r.DELETE("/students/:id", controllers.DeleteStudents)
	r.PATCH("/students/:id", controllers.EditStudents)
	r.Run()
}
