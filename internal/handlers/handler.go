package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

const (
	Gid     = "gid"
	Uid     = "uid"
	Name    = "name"
	Member  = "member"
	Comment = "comment"
	Home    = "home"
	Shell   = "shell"
	Indent  = "  "
)

// strParam returns string parameter from context
func strParam(c echo.Context, name string) string {
	val := c.Param(name)
	if len(val) == 0 {
		return ""
	}
	return val
}

func read(t *testing.T, resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.NotNil(t, body)
	return body
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
		// -1 is "wildcard" indicator
		return -1
	}
	return n
}
