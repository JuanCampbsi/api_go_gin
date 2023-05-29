package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/JuanCampbsi/api-go-gin/controllers"
	"github.com/JuanCampbsi/api-go-gin/database"
	"github.com/JuanCampbsi/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutsTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreatedStudentMock() {
	student := models.Student{Name: "Name Test", CPF: "12345678911", RG: "123456789"}
	database.DB.Save(&student)
	ID = int(student.ID)
}

func DeleteStudenMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifySearchStudentsAll(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	defer DeleteStudenMock()
	r := SetupRoutsTest()
	r.GET("/students", controllers.ViewStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "They should be the same")
}

func TestVerifyStatusCodeSearchStudentsWithParams(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	defer DeleteStudenMock()
	r := SetupRoutsTest()
	r.GET("/students/:id", controllers.SearchStudentsId)
	req, _ := http.NewRequest("GET", "/students/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "They should be the same")
}

func TestVerifyStatusCodeStudentsWithParamsCPFHandler(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	defer DeleteStudenMock()
	r := SetupRoutsTest()
	r.GET("/students/cpf/:cpf", controllers.SearchStudentsCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/00000000000", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "They should be the same")
}

func TestSearchStudentIDHandler(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	defer DeleteStudenMock()
	r := SetupRoutsTest()
	r.GET("/students/:id", controllers.SearchStudentsId)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "Name Test", studentMock.Name, "Names: they should be the same")
	assert.Equal(t, "12345678911", studentMock.CPF, "CPF: they should be the same")
	assert.Equal(t, "123456789", studentMock.RG, "RG: they should be the same")
	assert.Equal(t, http.StatusOK, response.Code, "Status code: they should be the same")
}

func TestEditStudentHandler(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	defer DeleteStudenMock()
	r := SetupRoutsTest()
	r.PATCH("/students/:id", controllers.EditStudents)
	student := models.Student{Name: "Edit Test", CPF: "42345678911", RG: "423456789"}
	valueJson, _ := json.Marshal(student)
	pathEdit := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(valueJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "Edit Test", studentMock.Name, "Name: they should be the same")
	assert.Equal(t, "42345678911", studentMock.CPF, "CPF: they should be the same")
	assert.Equal(t, "423456789", studentMock.RG, "RG: they should be the same")
	assert.Equal(t, http.StatusOK, response.Code, "Status Code: they should be the same")
}

func TestDeletStudentHandler(t *testing.T) {
	database.ConnectDataBase()
	CreatedStudentMock()
	r := SetupRoutsTest()
	r.DELETE("/students/:id", controllers.DeleteStudents)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
