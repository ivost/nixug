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

func TestContains(t *testing.T) {
	haystack := []string{"foo", "bar"}
	assert.True(t, contains(haystack, "foo"))
	assert.True(t, contains(haystack, "bar"))
	assert.False(t, contains(haystack, "baz"))
}

func TestContainsAll(t *testing.T) {
	haystack := []string{"foo", "bar", "baz"}
	needles := []string{"foo", "bar"}
	assert.True(t, containsAll(haystack, needles))
	needles = append(needles, "boo")
	assert.False(t, containsAll(haystack, needles))
}
