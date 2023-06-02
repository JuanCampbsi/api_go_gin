package routes

import (
	"github.com/JuanCampbsi/api-go-gin/controllers"
	docs "github.com/JuanCampbsi/api-go-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandlerRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/index", controllers.ViewPageIndex)
	r.GET("/students", controllers.ViewStudents)
	r.GET("/students/:id", controllers.SearchStudentsId)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentsCPF)
	r.POST("/students", controllers.CreateStudents)
	r.DELETE("/students/:id", controllers.DeleteStudents)
	r.PATCH("/students/:id", controllers.EditStudents)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
