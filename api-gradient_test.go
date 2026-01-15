package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadialGradient(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	gradient := RadialGradientNew()

	_ = gradient.RadialSet(1, 2, 3, 4, 5, 6)

	cx, cy, r, fx, fy, fr, _ := gradient.RadialGet()
	assert.Equal(t, float32(1), cx)
	assert.Equal(t, float32(2), cy)
	assert.Equal(t, float32(3), r)
	assert.Equal(t, float32(4), fx)
	assert.Equal(t, float32(5), fy)
	assert.Equal(t, float32(6), fr)

	_ = gradient.SetColorStops([]ColorStop{})
	colorStops, _ := gradient.GetColorStops()
	assert.Zero(t, len(colorStops))

	colorStops = []ColorStop{
		{Offset: 0, R: 1, G: 2, B: 3, A: 4},
		{Offset: 1, R: 10, G: 20, B: 30, A: 40},
	}
	_ = gradient.SetColorStops(colorStops)
	colorStops2, _ := gradient.GetColorStops()
	assert.Equal(t, colorStops, colorStops2)
}
