package handlers

import (
	"github.com/ivost/nixug/internal/models"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

/*
GET /users/query[?name=<nq>][&uid=<uq>][&gid=<gq>][&comment=<cq>][&home=< hq>][&shell=<sq>]
Return a list of users matching all of the specified query fields. The bracket notation indicates that any of the following query parameters may be supplied:
- name
- uid
- gid
- comment
- home
- shell
Only exact matches need to be supported.
Example Query: ​GET /users/query?shell=%2Fbin%2Ffalse Example Response:
[
{“name”: “dwoodlins”, “uid”: 1001, “gid”: 1001, “comment”: “”, “home”: “/home/dwoodlins”, “shell”: “/bin/false”}
]
*/

func GetAllUsers(c echo.Context) error {
	v := userService(c).FindUsers(nil)
	return c.JSONPretty(http.StatusOK, v, Indent)
}

func GetUserById(c echo.Context) error {
	ex := &models.User{UID: intParam(c, Uid), GID: -1}
	users := userService(c).FindUsers(ex)
	if len(users) == 0 {
		return c.String(http.StatusNotFound, "not found")
	}
	return c.JSONPretty(http.StatusOK, users[0], Indent)
}

func SearchUsers(c echo.Context) error {
	ex := userFromQuery(c)
	v := userService(c).FindUsers(ex)
	return c.JSONPretty(http.StatusOK, v, Indent)
}

func userService(c echo.Context) *services.UserService {
	return c.(*Context).UserSvc
}

func userFromQuery(c echo.Context) *models.User {
	user := models.User{}
	user.GID = intQueryParam(c, Gid)
	user.UID = intQueryParam(c, Uid)
	user.Name = strQueryParam(c, Name)
	user.Home = strQueryParam(c, Home)
	user.Comment = strQueryParam(c, Comment)
	user.Shell = strQueryParam(c, Shell)
	log.Printf("user example %+v", user)
	return &user
}
