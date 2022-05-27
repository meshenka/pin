package http_test

import (
	"testing"

	_ "github.com/meshenka/pin/transport/http"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
    assert.Equal(t, "severity", zerolog.LevelFieldName)
}
