package controllers

import (
	"net/http"

	"github.com/JuanCampbsi/api-go-gin/database"
	"github.com/JuanCampbsi/api-go-gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ViewStudents godoc
// @Summary        Views Students
// @Description    Route view students
// @Tags           Students
// @Accept         json
// @Produce        json
// @Success        200 {object} models.Student
// @Failure        400 {object} httputil.HTTPError
// @Router         /students [get]
func ViewStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

// SearchStudentsId godoc
// @Summary        Search Views Students
// @Description    Route search view students ID
// @Tags           Students
// @Accept         json
// @Produce        json
// @Param          id    path      int  true  "Student ID"
// @Success        200  {object}   models.Student
// @Failure        400  {object}   httputil.HTTPError
// @Failure        404  {object}   httputil.HTTPError
// @Failure        500  {object}   httputil.HTTPError
// @Router         /students/{id} [get]
func SearchStudentsId(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// CreateStudents  godoc
// @Summary        Created Students
// @Description    Route Created students
// @Tags           Students
// @Accept         json
// @Produce        json
// @Param          student  body  models.Student  true  "Model Student"
// @Success        200  {object}   models.Student
// @Failure        400  {object}   httputil.HTTPError
// @Failure        404  {object}   httputil.HTTPError
// @Failure        500  {object}   httputil.HTTPError
// @Router         /students [post]
func CreateStudents(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// DeleteStudents  godoc
// @Summary        Delete Students
// @Description    Route Delete students
// @Tags           Students
// @Accept         json
// @Produce        json
// @Param          id    path      int  true  "Student ID"
// @Success        200  {object}   models.Student
// @Failure        400  {object}   httputil.HTTPError
// @Failure        404  {object}   httputil.HTTPError
// @Failure        500  {object}   httputil.HTTPError
// @Router         /students/{id} [delete]
func DeleteStudents(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted student"})
}

// EditStudents  godoc
// @Summary        Edit Students
// @Description    Route Edit students
// @Tags           Students
// @Accept         json
// @Produce        json
// @Param          id    path      int  true  "Student ID"
// @Success        200  {object}   models.Student
// @Failure        400  {object}   httputil.HTTPError
// @Failure        404  {object}   httputil.HTTPError
// @Failure        500  {object}   httputil.HTTPError
// @Router         /students/{id} [patch]
func EditStudents(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

// SearchStudentsCPF godoc
// @Summary        Search Views StudentsCPF
// @Description    Route search view students CPF
// @Tags           Students
// @Accept         json
// @Produce        json
// @Param          cpf    path     string  true  "Student CPF"
// @Success        200  {object}   models.Student
// @Failure        400  {object}   httputil.HTTPError
// @Failure        404  {object}   httputil.HTTPError
// @Failure        500  {object}   httputil.HTTPError
// @Router         /students/cpf/{cpf} [get]
func SearchStudentsCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ViewPageIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"not found": "not found",
	})
}
