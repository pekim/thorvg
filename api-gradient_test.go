package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadialGradient(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	gradient := RadialGradientNew()

	err := gradient.RadialSet(1, 2, 3, 4, 5, 6)
	assert.NoError(t, err)

	cx, cy, r, fx, fy, fr, err := gradient.RadialGet()
	assert.NoError(t, err)
	assert.Equal(t, float32(1), cx)
	assert.Equal(t, float32(2), cy)
	assert.Equal(t, float32(3), r)
	assert.Equal(t, float32(4), fx)
	assert.Equal(t, float32(5), fy)
	assert.Equal(t, float32(6), fr)

	err = gradient.SetColorStops([]ColorStop{})
	assert.NoError(t, err)
	colorStops, err := gradient.GetColorStops()
	assert.NoError(t, err)
	assert.Zero(t, len(colorStops))

	colorStops = []ColorStop{
		{Offset: 0, R: 1, G: 2, B: 3, A: 4},
		{Offset: 1, R: 10, G: 20, B: 30, A: 40},
	}
	err = gradient.SetColorStops(colorStops)
	assert.NoError(t, err)
	colorStops2, err := gradient.GetColorStops()
	assert.NoError(t, err)
	assert.Equal(t, colorStops, colorStops2)
}
