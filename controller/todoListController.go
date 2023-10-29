package controller

import (
	"net/http"
	"strconv"
	"sync"
	"todolist/database"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

var (
	lock = sync.Mutex{}
)

func ListTodo(c echo.Context) error {
	lock.Lock()
	page := c.QueryParam("page")
	pageSize := c.QueryParam("pageSize")
	searchTitle := c.QueryParam("searchTitle")

	var Todolist []models.List

	offset := 0
	limit := 10

	if page != "" && pageSize != "" {
		pageNumber, _ := strconv.Atoi(page)
		size, _ := strconv.Atoi(pageSize)
		offset = (pageNumber - 1) * size
		limit = size
	}

	defer lock.Unlock()

	query := database.Db.Preload("SubList").Offset(offset).Limit(limit)

	if searchTitle != "" {
		query = query.Where("title LIKE ?", "%"+searchTitle+"%")
	}

	if err := query.Find(&Todolist); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Todolist)
}

func AddTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	Title := c.FormValue("Title")

	list := models.List{
		Title: Title,
	}

	// if err := c.Bind(&list); err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error)
	// }
	database.Db.Create(&list)
	return c.JSON(http.StatusOK, list)
}

func UpdateTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id := c.Param("id")
	var todolist models.List

	if err := database.Db.First(&todolist, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	todolist.Title = c.FormValue("Title")

	// if err := c.Bind(&todolist); err != nil {
	// 	return err
	// }

	database.Db.Save(&todolist)
	return c.JSON(http.StatusOK, todolist)
}

func DeleteTodo(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var todolist models.List

	id := c.Param("id")

	if err := database.Db.First(&todolist, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}
	database.Db.Delete(&todolist)
	return c.JSON(http.StatusOK, "Data Berhasil Terhapus")
}
