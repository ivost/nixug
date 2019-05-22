package services

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/ivost/nixug/internal/models"
	"github.com/ivost/nixug/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserService(t *testing.T) {
	// should load /etc/passwd when created
	cfg, err := config.NewConfig("test/" + config.DefaultConfigFile)
	assert.NoError(t, err)
	s, err := NewUserService(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.users)
	assert.True(t, len(s.users) > 0)
}

func TestUserByIdName(t *testing.T) {
	// func usersByIdName(example *models.User, users []models.User ) []models.User {
	var match []models.User

	// 2 users to test - root and
	users := test.NewTestUsers()
	match = usersByIdName(nil, users)
	assert.Equal(t, len(match), len(users))

	ex := models.User{}
	//curry
	assertLen := func(l int) { assert.Equal(t, l, len(usersByIdName(&ex, users))) }

	assertLen(1)

	ex.Name = "root"
	assertLen(1)

	ex.UID = -1
	assertLen(1)

	ex.Name = "sshd"
	assertLen(1)

	ex.UID = 1234
	assertLen(0)

	ex.UID = 121
	ex.Name = "sshd"
	assertLen(1)

	ex.UID = 120
	assertLen(0)

}

// test users
//{Name: "root", UID: 0, GID: 0, Comment: "root", Home: "/root", Shell: "/bin/bash"},
//{Name: "sshd", UID: 121, GID: 65534, Home: "/var/run/sshd", Shell: "/usr/sbin/nologin"},
func TestFindUsers(t *testing.T) {
	cfg, err := config.NewConfig("../../" + config.DefaultConfigFile)
	assert.NoError(t, err)
	s, _ := NewUserService(cfg)
	s.users = test.NewTestUsers()

	match := s.FindUsers(nil)
	assert.Equal(t, len(s.users), len(match))

	ex := models.User{}

	// closure
	assertLen := func(l int) { assert.Equal(t, l, len(s.FindUsers(&ex))) }

	ex.UID = -1
	ex.Name = "root"
	assertLen(1)

	ex.UID = 0
	assertLen(1)

	ex.UID = 123456
	assertLen(0)

	ex.UID = -1
	ex.GID = -1
	ex.Name = ""
	assertLen(2)

	ex2 := models.User{
		UID:   -1,
		Name:  "sshd",
		GID:   65534,
		Home:  "/var/run/sshd",
		Shell: "/usr/sbin/nologin",
	}

	assert.Equal(t, 1, len(s.FindUsers(&ex2)))
}
