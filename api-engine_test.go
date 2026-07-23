package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	defer testInitTerm(t)()

	major, minor, micro, version, commit, _ := Version()
	assert.Equal(t, 1, major)
	assert.Equal(t, 1, minor)
	assert.Equal(t, 0, micro)
	assert.Equal(t, "1.1.0", version)
	assert.Equal(t, libthorvgVersion, commit)
}
