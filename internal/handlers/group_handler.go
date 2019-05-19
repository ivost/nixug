package handlers

import (
	"github.com/ivost/nixug/internal/models"
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

func GetAllGroups(c echo.Context) error {
	srv := c.(*Context).GroupService()
	v := srv.FindGroups(nil)
	return c.JSONPretty(http.StatusOK, v, "    ")
}

func GetById(c echo.Context) error {
	//i := c.Param("id")
	//id := fmt.Sprintf("%v/%v", t, i)
	//log.Printf("GetById %v", id)
	srv := c.(*Context).GroupService()
	//todo - id int32 from path
	x := &models.Group{GID: 1}
	g := srv.FindGroups(x)
	if len(g) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSONPretty(http.StatusOK, g[0], "    ")
}
