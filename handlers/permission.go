package handlers

import (
	"fmt"
	"github.com/gaia-pipeline/gaia"
	"github.com/gaia-pipeline/gaia/services"
	"github.com/labstack/echo"
	"net/http"
)

func PermissionGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, gaia.PermissionsCategories)
}

func PermissionGroupGetAll(c echo.Context) error {
	store, _ := services.StorageService()
	pgs, err := store.PermissionGroupGetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pgs)
}

func PermissionGroupCreate(c echo.Context) error {
	var pg *gaia.PermissionGroup

	if err := c.Bind(&pg); err != nil {
		return c.String(http.StatusBadRequest, "Invalid Permission Group provided")
	}

	store, _ := services.PermissionManager()
	err := store.CreateGroup(pg, false)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Permission Group added")
}

func PermissionGroupDelete(c echo.Context) error {
	name := c.Param("name")

	store, _ := services.StorageService()
	err := store.PermissionGroupDelete(name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Permission Group %s deleted", name))
}