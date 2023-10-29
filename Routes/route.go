package routes

import (
	"net/http"

	"todolist/controller"
	"todolist/handler"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Testing")
	})

	//todo route
	e.GET("/listTodo", controller.ListTodo)
	e.POST("/addTodo", controller.AddTodo)
	e.PUT("/updateTodo/:id", controller.UpdateTodo)
	e.DELETE("/deleteTodo/:id", controller.DeleteTodo)

	//sublist route
	e.GET("/listSubListTodo", controller.ListSubList)
	e.POST("/addSubListTodo", controller.AddSubList)
	e.PUT("/updateSubListTodo/:id", controller.UpdateSubList)
	e.DELETE("/deleteSubList/:id", controller.DeleteSubList)

	//unit test
	e.GET("/echo", handler.GetTodos)
	e.DELETE("/TestdeleteSubList/:id", handler.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
