// +build integration

package test

import (
	"github.com/ivost/nixug/internal/services"
	"log"
	"testing"
)

func init() {
	// assumes 'make install' has been executed
	services.Run("pwd", true)
	services.Run("../../nixug", true)
}
//var fooAddr = flag.String(...)

func TestUsers(t *testing.T) {
	log.Printf("Integration test for /users")
	//services.Run("pgrep", true, "nixug")

	//f, err := foo.Connect(*fooAddr)
	// ...
}
