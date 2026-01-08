package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwCanvas(t *testing.T) {
	assert.NoError(t, Init())
	assert.NoError(t, EngineInit(2))

	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
	assert.NotZero(t, canvas)
	assert.Nil(t, canvas.Buffer())

	err := canvas.SwSetTarget(16, 16, 16, COLORSPACE_ARGB8888)
	assert.NoError(t, err)
	assert.Equal(t, 16*16, len(canvas.Buffer()))

	err = canvas.Destroy()
	assert.NoError(t, err)
	assert.Nil(t, canvas.Buffer())
}
