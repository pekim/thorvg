package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, RESULT_SUCCESS, EngineInit(2))
	// EngineTerm is not tested, as it appears to not work reliably in tests.

	result, major, minor, micro, version, commit := Version()
	assert.Equal(t, RESULT_SUCCESS, result)
	assert.Equal(t, 1, major)
	assert.Equal(t, 0, minor)
	assert.Equal(t, 0, micro)
	assert.Equal(t, "1.0.0", version)
	assert.Equal(t, "0a680b13d1753afdb85b498f26a03e605efe7c2e", commit)
}
