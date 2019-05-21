package services

import (
	"github.com/ivost/nixug/internal/test"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var (
	testFile = "test/group"
)

func TestReadLines(t *testing.T) {
	lines, err := readLines(testFile)
	assert.NoError(t, err)
	assert.True(t, len(lines) > 0)
}

// test file change notifications
func TestFileWatch(t *testing.T) {
	changed := false
	file := "/tmp/foo.bar"
	write := func(n int) {
		err := test.AppendToFile(file, strconv.Itoa(n))
		if check(err) {
			t.Fail()
		}
		time.Sleep(10 * time.Millisecond)
	}

	write(0)
	// start watching
	go watch(file, &changed)

	write(1)
	write(2)

	assert.True(t, changed)
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
