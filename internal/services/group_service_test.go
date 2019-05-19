package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroupService(t *testing.T) {
	s, err := NewGroupService()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.groups)
	assert.True(t, len(s.groups) > 0)
}
