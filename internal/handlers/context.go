// Package handlers contains REST handlers
// It processes http requests
// business logic is in services
package handlers

import (
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	GroupSvc *services.GroupService
	UserSvc  *services.UserService
}
