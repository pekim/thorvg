package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadialGradient(t *testing.T) {
	defer testInitTerm(t)()

	gradient := RadialGradientNew()

	assert.NoError(t, gradient.Set(1, 2, 3, 4, 5, 6))

	cx, cy, r, fx, fy, fr, _ := gradient.Get()
	assert.Equal(t, float32(1), cx)
	assert.Equal(t, float32(2), cy)
	assert.Equal(t, float32(3), r)
	assert.Equal(t, float32(4), fx)
	assert.Equal(t, float32(5), fy)
	assert.Equal(t, float32(6), fr)

	assert.NoError(t, gradient.SetColorStops([]ColorStop{}))
	colorStops, err := gradient.GetColorStops()
	assert.NoError(t, err)
	assert.Zero(t, len(colorStops))

	colorStops = []ColorStop{
		{Offset: 0, R: 1, G: 2, B: 3, A: 4},
		{Offset: 1, R: 10, G: 20, B: 30, A: 40},
	}
	assert.NoError(t, gradient.SetColorStops(colorStops))
	colorStops2, _ := gradient.GetColorStops()
	assert.Equal(t, colorStops, colorStops2)
}

func TestConicGradient(t *testing.T) {
	defer testInitTerm(t)()

	gradient := ConicGradientNew()

	assert.NoError(t, gradient.Set(1, 2, 3))

	cx, cy, angle, err := gradient.Get()
	assert.NoError(t, err)
	assert.Equal(t, float32(1), cx)
	assert.Equal(t, float32(2), cy)
	assert.Equal(t, float32(3), angle)
}
