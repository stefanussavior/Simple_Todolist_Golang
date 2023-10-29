package handler

import (
	"net/http"
	"strconv"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

var todos []models.Todo
var List []models.List1

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, list := range List {
		if list.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.NoContent(http.StatusNotFound)
}
