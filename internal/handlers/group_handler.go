package handlers

import (
	"github.com/ivost/nixug/internal/models"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*

GET /groups
Return a list of all groups on the system, a defined by /etc/group.


GET /groups/<gid>
Return a single group with <gid>. Return 404 if <gid> is not found.

GET /groups/query[?name=<nq>][&gid=<gq>][&member=<mq1>[&member=<mq2>][&. ..]]
Return a list of groups matching all of the specified query fields.
The bracket notation indicates that any of the following query parameters may be supplied:
- name
- gid
- member (repeated)
Any group containing all the specified members should be returned,
i.e. when query members are a subset of group members.
Example Query: ​GET /groups/query?member=_analyticsd&member=_networkd
Example Response:
[
{“name”: “_analyticsusers”, “gid”: 250, “members”: [“_analyticsd’,”_networkd”,”_timed”]}
]
*/

func GetAllGroups(c echo.Context) error {
	v := groupService(c).FindGroups(nil)
	return c.JSONPretty(http.StatusOK, v, Indent)
}

// GetGroupById returns user(s) for given GID
// We can't assume unique group ids
func GetGroupById(c echo.Context) error {
	ex := &models.Group{GID: intParam(c, Gid)}
	g := groupService(c).FindGroups(ex)
	if len(g) == 0 {
		return c.JSON(http.StatusNotFound, "not found")
	}
	return c.JSONPretty(http.StatusOK, g[0], Indent)
}

func SearchGroups(c echo.Context) error {
	ex := groupFromQuery(c)
	v := groupService(c).FindGroups(ex)
	return c.JSONPretty(http.StatusOK, v, Indent)
}

func groupService(c echo.Context) *services.GroupService {
	return c.(*Context).GroupSvc
}

func groupFromQuery(c echo.Context) *models.Group {
	g := models.Group{}
	g.GID = intQueryParam(c, Gid)
	g.Name = strQueryParam(c, Name)
	g.Members = strQueryArray(c, Member)
	//log.Printf("group example %+v", g)
	if g.GID == -1 && len(g.Name) == 0 && len(g.Members) == 0 {
		return nil
	}
	return &g
}
