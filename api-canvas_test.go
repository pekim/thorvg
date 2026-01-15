package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwCanvas(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
	_ = canvas
	assert.Nil(t, canvas.Buffer())

	_ = canvas.SwSetTarget(16, 16, 16, COLORSPACE_ARGB8888)
	assert.Equal(t, 4*16*16, len(canvas.Buffer()))

	_ = canvas.Destroy()
	assert.Nil(t, canvas.Buffer())
}
