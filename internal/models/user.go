package models

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string `json:"name"`
	UID     int    `json:"uid"`
	GID     int    `json:"gid"`
	Comment string `json:"comment"`
	Home    string `json:"home"`
	Shell   string `json:"shell"`
}

/*
/users/query[?name=<nq>][&uid=<uq>][&gid=<gq>][&comment=<cq>][&home=< hq>][&shell=<sq>]
*/
func NewUser(line string) (*User, error) {
	f := strings.Split(line, ":")
	// ivos:x:1001:1001:Ivo Stoyanov,,,:/home/ivos:/bin/bash
	if len(f) < 7 {
		return nil, fmt.Errorf("invalid user rec: %v", line)
	}
	u := &User{
		Name:    safeStr(f[0]),
		UID:     safeInt(f[2]),
		GID:     safeInt(f[3]),
		Comment: safeStr(f[4]),
		Home:    safeStr(f[5]),
		Shell:   safeStr(f[6]),
	}
	return u, nil
}

func safeStr(s string) string {
	return strings.Trim(s, " \t\n")
}

func safeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
