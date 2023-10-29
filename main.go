package main

import (
	routes "todolist/Routes"
	"todolist/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.ConnectionDatabase()
	routes.Route(echo.New())
}
