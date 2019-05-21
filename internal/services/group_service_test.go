package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"github.com/ivost/nixug/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroupService(t *testing.T) {
	// should load /etc/group when created
	cfg, err := config.NewConfig("test/" + config.DefaultConfigFile)
	assert.NoError(t, err)
	s, err := NewGroupService(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.groups)
	assert.True(t, len(s.groups) > 0)
}

func TestGroupByIdName(t *testing.T) {
	// func groupsByIdName(example *models.Group, groups []models.Group ) []models.Group {
	var match []models.Group

	// 2 groups to test - adm and root
	groups := test.NewTestGroups()
	match = groupsByIdName(nil, groups)
	assert.Equal(t, len(match), len(groups))

	ex := models.Group{GID: 0}
	//curry
	assertLen := func(l int) { assert.Equal(t, l, len(groupsByIdName(&ex, groups))) }

	assertLen(1)

	ex.Name = "root"
	assertLen(1)

	ex.GID = -1
	assertLen(1)

	ex.Name = "adm"
	assertLen(1)

	ex.GID = 4
	assertLen(1)

	ex.GID = 0
	assertLen(0)

	ex.GID = 4
	ex.Name = "foo"
	assertLen(0)
}

func TestFindGroups(t *testing.T) {
	cfg, err := config.NewConfig("../../" + config.DefaultConfigFile)
	assert.NoError(t, err)
	s, _ := NewGroupService(cfg)
	s.groups = test.NewTestGroups()

	match := s.FindGroups(nil)
	assert.Equal(t, len(s.groups), len(match))

	ex := models.Group{GID: 0}

	// closure
	assertLen := func(l int) { assert.Equal(t, l, len(s.FindGroups(&ex))) }

	ex.GID = -1
	ex.Name = "adm"
	ex.Members = []string{"foo"}
	assertLen(1)

	ex.Members = []string{"foo", "bar"}
	assertLen(0)

	ex.Name = ""
	ex.Members = []string{"foo"}
	assertLen(2)

	ex.Members = []string{"a"}
	assertLen(0)

	ex.Name = "log"
	ex.Members = []string{"foo"}
	assertLen(1)

}
