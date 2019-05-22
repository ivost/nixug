// +build integration

package test

import (
	"log"
	"os/exec"
	"testing"
)

func init() {
	// assumes 'make install' has been executed
	//services.Run("pwd", true)
	//services.Run("pkill", true, "nixug")
	// run in background
	//services.Run("../../nixug", true)
	cmd := exec.Command("../../nixug")
	cmd.Start()
}
//var fooAddr = flag.String(...)

func TestUsers(t *testing.T) {
	log.Printf("Integration test for /users")
	//services.Run("pgrep", true, "nixug")

	//f, err := foo.Connect(*fooAddr)
	// ...
}
