package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwCanvas(t *testing.T) {
	defer testInitTerm(t)()

	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
	assert.NotZero(t, canvas.canvas)
	assert.Nil(t, canvas.Buffer())

	assert.NoError(t, canvas.SwSetTarget(16, 16, 16, COLORSPACE_ARGB8888))
	assert.Equal(t, 4*16*16, len(canvas.Buffer()))

	assert.NoError(t, canvas.Destroy())
	assert.Nil(t, canvas.Buffer())
}

func TestCanvasDirtyRegions(t *testing.T) {
	defer testInitTerm(t)()

	width := 800
	height := 600
	assert.NoError(t, EngineInit(2))
	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
	assert.NoError(t, canvas.SwSetTarget(uint(width), uint(width), uint(height), COLORSPACE_ARGB8888))
	assert.NoError(t, canvas.Draw(true))

	count := canvas.GetDirtyRegionsCnt()
	assert.Equal(t, uint32(1), count)
	regions, err := canvas.GetDirtyRegions(count)
	assert.NoError(t, err)
	assert.Equal(t, []Region{{0, 0, int32(width), int32(height)}}, regions)

	assert.NoError(t, canvas.Sync())
}
