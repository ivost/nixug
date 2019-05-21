package services

import (
	"github.com/ivost/nixug/internal/test"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

// test file change notifications
func TestFileWatch(t *testing.T) {
	file := "/tmp/foo.bar"
	write := func(n int) {
		err := test.AppendToFile(file, strconv.Itoa(n))
		if check(err) {
			t.Fail()
		}
		time.Sleep(10 * time.Millisecond)
	}
	write(0)
	fw, err := NewFileWatcher(file)
	assert.NoError(t, err)
	assert.NotNil(t, fw)
	assert.False(t, fw.HasChanged())

	go fw.Watch()

	write(1)
	assert.True(t, fw.HasChanged())
	fw.SetDirty(false)
	assert.False(t, fw.HasChanged())
	write(2)
	assert.True(t, fw.HasChanged())
}
