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

	//tests := map[string]struct {
	//	input string
	//	want  *config.Config
	//}{
	//	"good": {
	//		input: config.DefaultConfigFile,
	//		want: test.NewConfig(),
	//	},
	//}
	//
	//for name, tc := range tests {
	//	t.Run(name, func(t *testing.T) {
	//		c, err := config.ReadConfig(tc.input)
	//		assert.Nil(t, err)
	//		assert.EqualValues(t, tc.want, c)
	//	})
	//}
}
