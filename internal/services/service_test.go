package services

import (
	"github.com/ivost/nixug/internal/test"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

var (
	testFile = "../../test/group"
)

func TestReadLines(t *testing.T) {
	lines, err := readLines(testFile)
	assert.NoError(t, err)
	assert.True(t, len(lines) > 0)
}

func TestWatch(t *testing.T) {
	// file change notifications
	note := make(chan string)
	file := "/tmp/foo.bar"
	var count int32

	write := func(n int) {
		err := test.AppendToFile(file, strconv.Itoa(n))
		if check(err) {
			t.Fail()
		}
		time.Sleep(10*time.Millisecond)
	}

	os.Remove(file)
	write(0)

	go watch(file, note)

	write(1)
	write(2)

	start := time.Now()
	maxDur := int64(2 * time.Second)
	maxCount := int32(3)

	readCount := func() int32 { return atomic.LoadInt32(&count) }
	dur := func() int64 { return time.Now().Sub(start).Nanoseconds() }

	// run for maxDur
	for dur() < maxDur {
		select {
		case event := <-note:
			//log.Printf("count %v, dur %v event %v ", readCount(), dur(), event)
			if strings.HasPrefix(event,"modified:") {
				if readCount() < maxCount {
					write(int(readCount()))
					atomic.AddInt32(&count,1)
				}
			}
		default:
		}
		if readCount() >= maxCount {
			break
		}
	}
	assert.Equal(t, maxCount, readCount())
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
