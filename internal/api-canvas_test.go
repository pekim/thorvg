package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwCanvas(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, RESULT_SUCCESS, EngineInit(2))

	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
	assert.NotZero(t, canvas)
	assert.Nil(t, canvas.Buffer())

	result := canvas.SwSetTarget(16, 16, 16, COLORSPACE_ARGB8888)
	assert.Equal(t, RESULT_SUCCESS, result)
	assert.Equal(t, 16*16, len(canvas.Buffer()))

	result = canvas.Destroy()
	assert.Equal(t, RESULT_SUCCESS, result)
	assert.Nil(t, canvas.Buffer())
}
