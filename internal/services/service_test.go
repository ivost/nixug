package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testFile = "../../test/group"
)

func TestReadLines(t *testing.T) {
	lines, err := readLines(testFile)
	assert.NoError(t, err)
	assert.True(t, len(lines) > 0)
}
