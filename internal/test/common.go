package test

import (
	"github.com/ivost/nixug/internal/models"
)

func NewTestGroups() []models.Group {
	// must be sorted by name
	g := []models.Group{
		{Name: "adm", GID: 4, Members: []string{"syslog", "foo"}},
		{Name: "log", GID: 42, Members: []string{"foo", "bar"}},
		{Name: "root", GID: 0},
	}
	return g
}
