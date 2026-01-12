package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeGradient(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()

	// assert that shape has no gradient
	gradient2, ok, err := shape.GetGradient()
	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Nil(t, gradient2)

	// set a gradient
	gradient3 := RadialGradientNew()
	assert.NoError(t, shape.SetGradient(gradient3))

	// assert that shape has a radial gradient
	gradient4, ok, err := shape.GetGradient()
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, gradient3, gradient4)
	assert.IsType(t, RadialGradient{}, gradient4)
}
