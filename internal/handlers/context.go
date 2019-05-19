package handlers

import (
	"github.com/ivost/nixug/internal/models"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"sync"
)

type Context struct {
	echo.Context
	mu sync.RWMutex

	m   *models.Group
	ct  string
	err error

	s *services.GroupService
}

func (c *Context) GroupService() *services.GroupService {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.s == nil {
		c.s, _ = services.NewGroupService()
	}
	return c.s
}

func (c *Context) Valid() error {
	// todo
	//t := c.Param("t")
	//i := c.Param("id")
	//id := fmt.Sprintf("%v/%v", t, i)
	//
	//c.m = c.GroupSvc.GetGroup(id)
	//if c.m == nil {
	//	return fmt.Errorf("id %v not found", id)
	//}
	return nil
}
