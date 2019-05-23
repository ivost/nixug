package config_test

import (
	"github.com/ivost/nixug/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	c, err := config.NewConfig("../../config.json")
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.True(t, len(c.GroupFile) > 0)
	assert.True(t, len(c.UserFile) > 0)
}
