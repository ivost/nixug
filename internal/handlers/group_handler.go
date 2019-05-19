package handlers

import (
	"fmt"
	"github.com/ivost/nix_users/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)


func GetGroup(c echo.Context) error {
	v := GetById(c)
	if v == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSONPretty(http.StatusOK, v, "    ")
}

func GetGroupsAll(c echo.Context) error {
	srv := c.(*Context).GroupService()
	v := srv.GetGroupsAll()
	return c.JSONPretty(http.StatusOK, v, "    ")
}

func GetById(c echo.Context) *models.Group {
	t := c.Param("t")
	i := c.Param("id")
	id := fmt.Sprintf("%v/%v", t, i)
	//log.Printf("GetById %v", id)
	srv := c.(*Context).GroupService()
	return srv.GetGroup(id)
}
