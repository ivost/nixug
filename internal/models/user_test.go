package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {

	tests := map[string]struct {
		line string
		want *User
		good bool
	}{
		"good1": {
			line: "ivos:x:1001:1001:Ivo Stoyanov,,,:/home/ivos:/bin/bash",
			want: &User{Name: "ivos", GID: 1001, UID: 1001, Comment: "Ivo Stoyanov,,,", Home: "/home/ivos", Shell: "/bin/bash"},
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
			u, err := NewUser(tc.line)
			// negative test?
			if !tc.good {
				if err == nil {
					t.Fail()
				}
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tc.want, u)
		})
	}
}
