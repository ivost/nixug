package models

import (
	"fmt"
)

type Identity struct {
	ID   string
	Name string
	Hash []byte
}

var ID = "nix"

func GetIdentity(name string) (Identity, error) {
	var id = Identity{}
	// hardcoded nix/nix auth for now
	if name == "nix" {
		id.ID = ID
		// hardcoded crypt hash for password = nix
		id.Hash = []byte("$2y$12$qPUU4SYAqdjist.T.Icxh.gMu9H.ZjXJh2dDHnfeFmB4V4/flegkW")
		return id, nil
	}
	return id, fmt.Errorf("%v not found", name)
}
