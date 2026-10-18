package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScene(t *testing.T) {
	defer testInitTerm(t)()

	scene := SceneNew()
	assert.NotZero(t, scene.paint_)
	assert.NoError(t, scene.AddEffectMotionBlur(2, 90, 50))
}
