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

func TestShapePath(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()

	// assert that shape has no path points
	commands, points, err := shape.GetPath()
	assert.NoError(t, err)
	assert.Nil(t, commands)
	assert.Nil(t, points)

	// set a path
	commands1 := []PathCommand{PATH_COMMAND_MOVE_TO, PATH_COMMAND_LINE_TO, PATH_COMMAND_LINE_TO, PATH_COMMAND_CLOSE}
	points1 := []Point{{X: 0, Y: 1}, {X: 2, Y: 3}, {X: 4, Y: 5}, {X: 6, Y: 7}}
	assert.NoError(t, shape.AppendPath(commands1, points1))

	// assert that the path is retrieved
	commands2, points2, err := shape.GetPath()
	assert.NoError(t, err)
	assert.Equal(t, commands1, commands2)
	assert.Equal(t, points1, points2)
}
