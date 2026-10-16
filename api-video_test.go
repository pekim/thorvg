//go:build linux && amd64

package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideo(t *testing.T) {
	defer testInitTerm(t)()

	video := VideoNew()
	assert.NotNil(t, video)

	assert.NoError(t, video.Del())
}
