package handlers

import (
	"fmt"
	"github.com/ivost/nix_users/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

//func CreateMeta(c echo.Context) error {
//	r := new(models.Meta)
//	if err := c.Bind(r); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	// todo: logger to service
//	//c.Logger().Printf("CreateMeta - %+v", r)
//	// todo - custom validator
//	//if err := c.Validate(r); err != nil {
//	//	return c.JSON(http.StatusBadRequest, err.Error())
//	//}
//	srv := c.(*Context).GroupService()
//	return srv.Create(*r)
//}
//
//func DeleteMeta(c echo.Context) error {
//	v := GetById(c)
//	if v == nil {
//		return c.NoContent(http.StatusNotFound)
//	}
//	//c.Logger().Printf("DeleteMeta - %v", v.Id)
//	srv := c.(*Context).GroupService()
//	return srv.Delete(v.Id)
//}
//
//func UpdateMeta(c echo.Context) error {
//	v := GetById(c)
//	if v == nil {
//		return c.NoContent(http.StatusNotFound)
//	}
//	r := new(models.Meta)
//	if err := c.Bind(r); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	//c.Logger().Printf("UpdateMeta -  %+v", r)
//	if len(r.Id) > 0 && r.Id != v.Id {
//		return c.JSON(http.StatusBadRequest, "Id mismatch")
//	}
//	srv := c.(*Context).GroupService()
//	return srv.Update(v.Id, *r)
//}

func GetGroup(c echo.Context) error {
	v := GetById(c)
	if v == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSONPretty(http.StatusOK, v, "    ")
}

func GetMetaAll(c echo.Context) error {
	srv := c.(*Context).GroupService()
	v := srv.GetMetaAll()
	return c.JSONPretty(http.StatusOK, v, "    ")
}

func GetById(c echo.Context) *models.Group {
	t := c.Param("t")
	i := c.Param("id")
	id := fmt.Sprintf("%v/%v", t, i)
	//log.Printf("GetById %v", id)
	srv := c.(*Context).GroupService()
	return srv.GetMeta(id)
}
