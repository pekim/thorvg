package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	defer testInitTerm(t)()

	major, minor, micro, version, commit, _ := Version()
	assert.Equal(t, 1, major)
	assert.Equal(t, 0, minor)
	assert.Equal(t, 0, micro)
	assert.Equal(t, "1.0.0", version)
	assert.Equal(t, "1a43240ec3ffdaa689412e7cd52e83cf8118e2b9", commit)
}
