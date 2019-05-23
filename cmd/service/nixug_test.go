package main

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialize(t *testing.T) {
	cfg, err := config.NewConfig(config.DefaultConfigFile)
	assert.NoError(t, err)
	e, err := initialize(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.NotNil(t, e)
}
