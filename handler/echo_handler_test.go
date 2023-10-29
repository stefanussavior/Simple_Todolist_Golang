package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/controller"
	"todolist/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/listTodo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controller.ListTodo(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteTodo(t *testing.T) {
	e := echo.New()

	List = []models.List1{
		{ID: 1, Title: "Task 1"},
		{ID: 2, Title: "Task 2"},
	}

	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, DeleteTodo(c))
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
