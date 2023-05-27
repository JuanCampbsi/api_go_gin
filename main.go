package main

import (
	"github.com/JuanCampbsi/api-go-gin/database"
	"github.com/JuanCampbsi/api-go-gin/routes"
)

func main() {
	database.ConnectDataBase()
	routes.HandlerRequests()
}
