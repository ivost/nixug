package models

import (
	"fmt"
	"strings"
)

type Group struct {
	GID     int      `json:"gid"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func NewGroup(line string) (*Group, error) {
	f := strings.Split(line, ":")
	//adm:x:4:syslog,tap
	if len(f) < 4 {
		return nil, fmt.Errorf("invalid group rec: %v", line)
	}
	g := &Group{Name: f[0], GID: safeInt(f[2])}
	if len(f[3]) > 0 {
		g.Members = strings.Split(f[3], ",")
	}
	return g, nil
}
