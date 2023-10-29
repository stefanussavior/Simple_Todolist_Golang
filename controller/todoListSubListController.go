package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"todolist/database"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

func ListSubList(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var SubList []models.SubList

	if err := database.Db.Find(&SubList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, SubList)
}

func AddSubList(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	idListStr := c.FormValue("id_list")
	description := c.FormValue("description")
	file, err := c.FormFile("file_description")
	if err != nil {
		fmt.Println(err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer src.Close()

	fileUpload, err := io.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// var SubList models.SubList
	idStr, err := strconv.Atoi(idListStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	SubList := models.SubList{
		IdList:          idStr,
		Description:     description,
		FileDescription: fileUpload,
	}

	if err := database.Db.Model(&SubList).Association("SubList").Append(&SubList); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// if err := c.Bind(&SubList); err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// }
	database.Db.Create(&SubList)
	return c.JSON(http.StatusOK, SubList)
}

func UpdateSubList(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var SubList models.SubList
	id := c.Param("id")

	if err := database.Db.First(&SubList, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := database.Db.Model(&SubList).Association("SubList").Replace(&SubList); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	idListStr := c.FormValue("id_list")
	description := c.FormValue("description")

	idListInt, err := strconv.Atoi(idListStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	file, err := c.FormFile("file_description")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	FileDescription, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	// if err := c.Bind(&SubList); err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// SubList = models.SubList{
	// 	IdList:          idListInt,
	// 	Description:     description,
	// 	FileDescription: FileDescription,
	// }

	SubList.IdList = idListInt
	SubList.Description = description
	SubList.FileDescription = FileDescription

	if err := database.Db.Save(&SubList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, SubList)
}

func DeleteSubList(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var Sublist models.SubList

	id := c.Param("id")

	if err := database.Db.First(&Sublist, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	database.Db.Delete(&Sublist)
	return c.JSON(http.StatusOK, "Data Sudah Terhapus")
}
