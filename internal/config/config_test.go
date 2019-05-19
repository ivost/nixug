package config_test

import (
	"github.com/ivost/nix_users/internal/config"
	"github.com/ivost/nix_users/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestConfig(t *testing.T) {
	tests := map[string]struct {
		input string
		want  *config.Config
	}{
		"good": {
			input: config.DefaultConfigFile,
			want: test.NewConfig(),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := config.ReadConfig(tc.input)
			assert.Nil(t, err)
			assert.EqualValues(t, tc.want, c)
		})
	}
}
