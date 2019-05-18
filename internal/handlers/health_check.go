package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var filters = map[string]string{
	"redis_version": "",
	"process_id": "",
	"tcp_port": "",
	"uptime_in_seconds": "",
	"uptime_in_days": "",
	"config_file": "",
	"used_memory_human": "",
	"used_memory_peak_human": "",
	"total_system_memory_human": "",
	"maxmemory_human": "",
	"total_connections_received": "",
	"total_commands_processed": "",
	"used_cpu_sys": "",
	"used_cpu_user": "",
}

var sb = new(strings.Builder)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	sb.Reset()
	return c.String(http.StatusOK, sb.String())
}


func filterOut(s string) {
	a := strings.Split(s,"\r\n")
	for _, l := range a {
		//log.Printf(l)
		ll := strings.Split(l,":")
		if len(ll) < 2 {
			continue
		}
		if _, ok := filters[ll[0]]; ok {
			sb.WriteString(l)
			sb.WriteString("\n")
		}
	}
}