package handlers

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const (
	Id     = "id"
	Gid    = "gid"
	Name   = "name"
	Member = "member"
	Indent = "  "
)

type Handler interface {
	DoGet() error
	//DoInfo() error
	//DoPost() error
	//DoPut() error
	//DoPatch() error
	//DoDelete() error
}

//type handler struct {
//	context *Context
//	lock    sync.RWMutex
//}

//func (h *handler) PrepGet() (*models.ReadParam, error) {
//	c := h.context
//	err := c.Valid()
//	if err != nil {
//		return nil, err
//	}
//	//log.Printf("client accepts type %v, stream content type is %v", at, m.Content)
//	rp := parseQuery(c)
//	return &rp, nil
//}

// strParam returns string parameter from context
func strParam(c echo.Context, name string) string {
	val := c.Param(name)
	if len(val) == 0 {
		return ""
	}
	return val
}

// intParam returns int parameter from context, -1 if not found/error
func intParam(c echo.Context, name string) int {
	val := strParam(c, name)
	if len(val) == 0 {
		return -1
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return -1
	}
	return n
}

// strQueryParam returns string parameter from query
func strQueryParam(c echo.Context, name string) string {
	return c.QueryParam(name)
}

func strQueryArray(c echo.Context, name string) []string {
	m := c.QueryParams()
	if m == nil {
		return nil
	}
	return m[name]
}

// intQueryParam returns int parameter from query, -1 if not found/error
func intQueryParam(c echo.Context, name string) int {
	s := strQueryParam(c, name)
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}

//func check(err error) bool {
//	if err == nil {
//		return false
//	}
//	log.Print(err.Error())
//	return true
//}
