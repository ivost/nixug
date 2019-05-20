package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroup(t *testing.T) {

	tests := map[string]struct {
		line string
		want *Group
		good bool
	}{
		"good1": {
			line: "root:x:0:",
			want: &Group{Name: "root", GID: 0},
			good: true,
		},
		"good2": {
			line: "adm:x:4:syslog,tap",
			want: &Group{Name: "adm", GID: 4, Members: []string{"syslog", "tap"}},
			good: true,
		},
		"bad1": {
			line: "root:x:",
			want: nil,
			good: false,
		},
		"bad2": {
			line: "   ",
			want: nil,
			good: false,
		},
		"bad3": {
			line: "foo:x:bar:",
			want: nil,
			good: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g, err := NewGroup(tc.line)
			// negative test?
			if !tc.good {
				if err == nil {
					t.Fail()
				}
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tc.want, g)
		})
	}
}
