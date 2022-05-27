package cmd_test

import (
	"os"
	"testing"

	"github.com/meshenka/pin/cmd"
	"github.com/stretchr/testify/assert"
)

const env string = "env_unit_test"

func TestEnv(t *testing.T) {
	have := cmd.Env(env, "fallback")
	want := "fallback"

	assert.Equal(t, want, have)

	os.Setenv(env, "value")
	defer os.Unsetenv(env)

	have = cmd.Env(env, "fallback")
	want = "value"

	assert.Equal(t, want, have)

}
