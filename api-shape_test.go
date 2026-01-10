package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeGradient(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()
	gradient := RadialGradientNew()
	err := shape.SetGradient(gradient)
	assert.NoError(t, err)
	gradient2, err := shape.GetGradient()
	assert.NoError(t, err)
	assert.Equal(t, gradient, gradient2)
}
