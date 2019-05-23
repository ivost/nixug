// +build integration

package test

import (
	"github.com/ivost/nixug/internal/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os/exec"
	"testing"
)

const (
	BaseUrl = "http://localhost:8080"
)
func init() {
	//services.Run("pkill", false, "nixug")
	// assumes 'make install' has been executed
	// run in background
	cmd := exec.Command("../../nixug")
	cmd.Start()
}

func TestHealth(t *testing.T) {
	t.Log("Integration test for /health")
	d := read(t,"/health")
	assert.Equal(t, "OK", string(d))
}

func TestGroups(t *testing.T) {
	t.Log("Integration test for /groups")
	g, err := models.NewGroupsFromJson(read(t,"/groups"))
	assert.NoError(t, err)
	assert.True(t, len(g) > 0)
}

func TestUsers(t *testing.T) {
	t.Log("Integration test for /users")
}

func read(t *testing.T, url string) []byte {
	resp, err := http.Get(BaseUrl + url)
	assert.NoError(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)
	//t.Log(string(body))
	return body
}
