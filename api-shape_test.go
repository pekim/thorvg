package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeGradient(t *testing.T) {
	defer testInitTerm(t)()

	shape := ShapeNew()

	// assert that shape has no gradient
	gradient2, ok, _ := shape.GetGradient()
	assert.False(t, ok)
	assert.Nil(t, gradient2)

	// set a gradient
	gradient3 := RadialGradientNew()
	_ = shape.SetGradient(gradient3)

	// assert that shape has a radial gradient
	gradient4, ok, _ := shape.GetGradient()
	assert.True(t, ok)
	assert.Equal(t, gradient3, gradient4)
	assert.IsType(t, RadialGradient{}, gradient4)
}

func TestShapePath(t *testing.T) {
	defer testInitTerm(t)()

	shape := ShapeNew()

	// assert that shape has no path points
	commands, points, _ := shape.GetPath()
	assert.Nil(t, commands)
	assert.Nil(t, points)

	// set a path
	commands1 := []PathCommand{PATH_COMMAND_MOVE_TO, PATH_COMMAND_LINE_TO, PATH_COMMAND_LINE_TO, PATH_COMMAND_CLOSE}
	points1 := []Point{{X: 0, Y: 1}, {X: 2, Y: 3}, {X: 4, Y: 5}, {X: 6, Y: 7}}
	_ = shape.AppendPath(commands1, points1)

	// assert that the path is retrieved
	commands2, points2, _ := shape.GetPath()
	assert.Equal(t, commands1, commands2)
	assert.Equal(t, points1, points2)
}
