package handlers

import (
	"fmt"
	"github.com/ivost/nix_users/internal/models"
	"github.com/ivost/nix_users/internal/services"
	"github.com/labstack/echo/v4"
	"sync"
)

type Context struct {
	echo.Context

	m *models.Group
	ct string
	err error

	mutex sync.RWMutex
	GroupSvc *services.GroupService
}

func (c *Context) GroupService() *services.GroupService {
	return c.GroupSvc
}

func (c * Context) Valid() error {
	t := c.Param("t")
	i := c.Param("id")
	id := fmt.Sprintf("%v/%v", t, i)

	c.m = c.GroupSvc.GetMeta(id)
	if c.m == nil {
		return fmt.Errorf("id %v not found", id)
	}
	if t != c.m.Type {
		return fmt.Errorf("type mismatch: %v - %v", t, c.m.Type)
	}

	//c.ct = GetAcceptType(c)
	//if len(c.ct) == 0 {
	//	return c.JSON(http.StatusPreconditionFailed, "content type")
	//}
	//rm := c.Request().Method
	//if rm == echo.POST || rm == echo.PUT || rm == echo.PATCH {
	//	c.Set(echo.HeaderContentType, c.ct)
	//}

	//if rm == echo.GET {
	//	c.Set(echo.HeaderAccept, c.ct)
	//}
	return nil
}
