package handlers

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetUserById(t *testing.T) {
	// Setup
	//e := NewEcho(t, "")
	//req := httptest.NewRequest(http.MethodGet, "/users/0", nil)
	//rec := httptest.NewRecorder()
	//
	//c := e.NewContext(req, rec)
	//
	//if assert.NoError(t, GetUserById(c)) {
	//	assert.Equal(t, http.StatusOK, rec.Code)
	//}
}

func NewEcho(t *testing.T, conf string) *echo.Echo {
	cfg, err := config.NewConfig(conf)
	assert.NoError(t, err)
	gs, err := services.NewGroupService(cfg)
	assert.NoError(t, err)
	us, err := services.NewUserService(cfg)
	assert.NoError(t, err)

	// new echo instance
	e := echo.New()
	e.HideBanner = true

	// convert echo context to our context - make available in middleware
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{Context: c, GroupSvc: gs, UserSvc: us}
			return h(cc)
		}
	})
	return e

}

/*
req := httptest.NewRequest(echo.GET, "/", nil)
rec := httptest.NewRecorder()
c := e.NewContext(req, rec)
q := req.URL.Query()
q.Add("units", "value")
req.URL.RawQuery = q.Encode()
if assert.NoError(t, CreateWallet(c)) {
  assert.Equal(t, http.StatusOK, rec.Code)
  assert.Contains(t, rec.Body.String(), "<YOUR_JSON_STRING>")
}
*/
