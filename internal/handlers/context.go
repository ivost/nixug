// Package handlers contains REST handlers
// It deals with handling htpp requests
// business logic is in services
package handlers

import (
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	//mu       sync.RWMutex
	//err      error
	GroupSvc *services.GroupService
}

//func (c *Context) NewGroupService() *services.GroupService {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	if c.GroupSvc == nil {
//		log.Printf("NewGroupService")
//		c.GroupSvc, _ = services.NewGroupService()
//	}
//	return c.GroupSvc
//}
