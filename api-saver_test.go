package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaverSavePaint(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	background := ShapeNew()
	assert.NoError(t, background.AppendRect(0, 0, 200, 150, 0, 0, true))
	assert.NoError(t, background.SetFillColor(0xff, 0xff, 0xff, 0xff))

	rect := ShapeNew()
	assert.NoError(t, rect.AppendRect(50, 25, 100, 100, 0, 0, true))
	assert.NoError(t, rect.SetFillColor(0xff, 0x00, 0x00, 0xff))

	scene := SceneNew()
	assert.NoError(t, scene.Push(background))
	assert.NoError(t, scene.Push(rect))

	// saver := SaverNew()
	// assert.NoError(t, saver.SavePaint(scene, "saved.tvg", 100))
	// assert.NoError(t, saver.Sync())
}
