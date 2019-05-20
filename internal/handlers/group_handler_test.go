package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//q := make(url.Values)
//q.Set("email", "jon@labstack.com")
//req := httptest.NewRequest(http.MethodPost, "/?"+q.Encode(), nil)

func TestGroupFromQuery(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("gid", "42")
	q.Set("name", "foo")
	q.Set("member", "a")
	q.Add("member", "b")
	//q.Set("members", "b" )
	req := httptest.NewRequest(http.MethodGet, "/groups?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// sut
	g := groupFromQuery(c)
	// Assertions
	is := assert.New(t)
	if is.NotNil(g) {
		is.Equal(42, g.GID)
		is.Equal("foo", g.Name)
		if is.Equal(2, len(g.Members)) {
			is.Equal("a", g.Members[0])
			is.Equal("b", g.Members[1])
		}
	}
}
