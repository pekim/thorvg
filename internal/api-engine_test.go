package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, RESULT_SUCCESS, EngineInit(2))
	// EngineTerm is not tested, as it appear to not work reliably in tests.

	result, major, minor, micro, version := Version()
	assert.Equal(t, RESULT_SUCCESS, result)
	assert.Equal(t, 1, major)
	assert.Equal(t, 0, minor)
	assert.Equal(t, 0, micro)
	assert.Equal(t, "1.0.0", version)
}
