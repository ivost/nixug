package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

var sb = new(strings.Builder)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	log.Printf("HealthCheck")
	sb.Reset()
	sb.WriteString("OK")
	return c.String(http.StatusOK, sb.String())
}
