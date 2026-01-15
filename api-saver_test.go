package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaverSavePaint(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	background := ShapeNew()
	_ = background.AppendRect(0, 0, 200, 150, 0, 0, true)
	_ = background.SetFillColor(0xff, 0xff, 0xff, 0xff)

	rect := ShapeNew()
	_ = rect.AppendRect(50, 25, 100, 100, 0, 0, true)
	_ = rect.SetFillColor(0xff, 0x00, 0x00, 0xff)

	scene := SceneNew()
	_ = scene.Push(background)
	_ = scene.Push(rect)

	// saver := SaverNew()
	// _= saver.SavePaint(scene, "saved.tvg", 100)
	// _= saver.Sync()
}
