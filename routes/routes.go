package routes

import (
	"github.com/JuanCampbsi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/index", controllers.ViewPageIndex)
	r.GET("/students", controllers.ViewStudents)
	r.GET("/students/:id", controllers.SearchStudentsId)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentsCPF)
	r.POST("/students", controllers.CreateStudents)
	r.DELETE("/students/:id", controllers.DeleteStudents)
	r.PATCH("/students/:id", controllers.EditStudents)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
