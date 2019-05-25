// +build integration

package test

import (
	"encoding/json"
	"github.com/ivost/nixug/internal/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	BaseUrl = "http://localhost:8080"
)
func init() {
	//services.Run("pkill", false, "nixug")
	// assumes 'make install' has been executed
	// run in background
	//cmd := exec.Command("../../nixug")
	//cmd.Start()
}

func TestHealth(t *testing.T) {
	t.Log("Integration test for /health")
	d := getOK(t,"/health")
	assert.NotNil(t, d)
	assert.Equal(t, "OK", string(d))
}

func TestGroups(t *testing.T) {
	t.Logf("Integration test for /groups")
	assertGroups(t,"/groups")
	assertGroup (t,"/groups/0")
	get404(t,"/groups/123456")
	assertGroups(t,"/groups/query?name=sudo")
	assertGroups0(t,"/groups/query?name=foo-bar")
	assertGroups0(t,"/groups/query?name=adm&member=tap&member=syslog&member=foo")
	assertGroups(t,"/groups/query?name=adm&member=tap&member=syslog")
	assertGroups(t,"/groups/query?name=adm&member=tap")
	assertGroups(t,"/groups/query?name=adm")
}

func TestUsers(t *testing.T) {
	t.Log("Integration test for /users")
	assertUsers(t,"/users")
	assertUser (t,"/users/0")
	get404(t,"/users/123456")
	assertUsers(t,"/users/query?name=root")
	assertUsers0(t,"/users/query?name=root-beer")
	assertUsers(t,"/users/query?name=root&uid=0")
	assertUsers0(t,"/users/query?name=root&uid=123")
	assertUsers(t,"/users/query?name=tap&member=tap&member=ivos")
	assertUsers(t,"/users/query?name=tap&member=tap")
	assertUsers(t,"/users/1001/groups")

}

func assertGroup(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	j := getOK(t, q)
	assert.True(t, len(j) > 0)
	var x models.Group
	err := json.Unmarshal(j, &x)
	assert.NoError(t, err)
}

func assertGroups(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	g, err := models.NewGroupsFromJson(getOK(t, q))
	assert.NoError(t, err)
	assert.NotNil(t, g)
	assert.True(t, len(g) > 0)
}

func assertGroups0(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	g, err := models.NewGroupsFromJson(getOK(t, q))
	assert.NoError(t, err)
	assert.NotNil(t, g)
	assert.True(t, len(g) == 0)
}

func assertUser(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	j := getOK(t, q)
	assert.True(t, len(j) > 0)
	var x models.User
	err := json.Unmarshal(j, &x)
	assert.NoError(t, err)
}

func assertUsers(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	u, err := models.NewUsersFromJson(getOK(t,q))
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.True(t, len(u) > 0)
}

func assertUsers0(t *testing.T, q string) {
	t.Logf("Query: %v", q)
	u, err := models.NewUsersFromJson(getOK(t,q))
	//log.Printf("users %+v", u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.True(t, len(u) == 0)
}

func getOK(t *testing.T, url string) []byte {
	resp, err := http.Get(BaseUrl + url)
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.NotNil(t, body)
		//t.Log(string(body))
		return body
	}
	return nil
}

func get404(t *testing.T, url string) {
	resp, _ := http.Get(BaseUrl + url)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}
