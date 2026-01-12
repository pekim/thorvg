package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaintGetAABB(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()
	assert.NoError(t, shape.AppendRect(10, 20, 30, 40, 0, 0, true))
	x, y, width, height, err := shape.GetAABB()
	assert.NoError(t, err)
	assert.Equal(t, float32(10), x)
	assert.Equal(t, float32(20), y)
	assert.Equal(t, float32(30), width)
	assert.Equal(t, float32(40), height)
}

func TestPaintGetOBB(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()
	assert.NoError(t, shape.AppendRect(10, 20, 30, 40, 0, 0, true))
	points, err := shape.GetOBB()
	assert.NoError(t, err)
	assert.Equal(t, [4]Point{
		{X: 10, Y: 20}, // top left
		{X: 40, Y: 20}, // top right
		{X: 40, Y: 60}, // bottom right
		{X: 10, Y: 60}, // bottom left
	}, points)
}

func TestPaintClip(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	shape := ShapeNew()

	clipper, ok := shape.GetClip()
	assert.Nil(t, clipper)
	assert.False(t, ok)

	assert.NoError(t, shape.SetClip(ShapeNew()))
	clipper, ok = shape.GetClip()
	assert.NotNil(t, clipper)
	assert.True(t, ok)
	assert.IsType(t, Shape{}, clipper)

	assert.NoError(t, shape.SetClip(TextNew()))
	clipper, ok = shape.GetClip()
	assert.NotNil(t, clipper)
	assert.True(t, ok)
	assert.IsType(t, Text{}, clipper)
}
