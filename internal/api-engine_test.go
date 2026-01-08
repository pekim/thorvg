package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, RESULT_SUCCESS, tvg_engine_init(2))
	assert.Equal(t, RESULT_SUCCESS, tvg_engine_term())
}
