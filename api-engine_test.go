package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	major, minor, micro, version, commit, err := Version()
	assert.NoError(t, err)
	assert.Equal(t, 1, major)
	assert.Equal(t, 0, minor)
	assert.Equal(t, 0, micro)
	assert.Equal(t, "1.0.0", version)
	assert.Equal(t, "0a680b13d1753afdb85b498f26a03e605efe7c2e", commit)
}
