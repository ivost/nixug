package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

//var filters = map[string]string{}

var sb = new(strings.Builder)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	log.Printf("HealthCheck")
	sb.Reset()
	sb.WriteString("OK")
	return c.String(http.StatusOK, sb.String())
}

//func filterOut(s string) {
//	a := strings.Split(s, "\r\n")
//	for _, l := range a {
//		//log.Printf(l)
//		ll := strings.Split(l, ":")
//		if len(ll) < 2 {
//			continue
//		}
//		if _, ok := filters[ll[0]]; ok {
//			sb.WriteString(l)
//			sb.WriteString("\n")
//		}
//	}
//}
