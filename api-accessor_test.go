package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessorNoChildren(t *testing.T) {
	defer testInitTerm(t)()

	shape := ShapeNew()
	accessor := AccessorNew()
	_ = accessor.Set(shape, func(paint Paint) bool {
		assert.IsType(t, Shape{}, paint)
		assert.Equal(t, shape.paint_, paint.paint())
		return true
	})
}

func TestAccessor2Children(t *testing.T) {
	defer testInitTerm(t)()

	scene := SceneNew()
	child1 := ShapeNew()
	_ = scene.Push(child1)
	child2 := ShapeNew()
	_ = scene.Push(child2)

	expected := []struct {
		paint Paint
		typ   Paint
	}{
		{scene, Scene{}},
		{child1, Shape{}},
		{child2, Shape{}},
	}
	i := 0

	accessor := AccessorNew()
	_ = accessor.Set(scene, func(paint Paint) bool {
		assert.IsType(t, expected[i].typ, paint)
		assert.Equal(t, expected[i].paint.paint(), paint.paint())
		i++
		return true
	})
	assert.Equal(t, len(expected), i, "all expected visited")
}

func TestAccessorNestedChildren(t *testing.T) {
	defer testInitTerm(t)()

	scene := SceneNew()
	child1 := SceneNew()
	_ = scene.Push(child1)
	child11 := ShapeNew()
	_ = child1.Push(child11)
	child2 := ShapeNew()
	_ = scene.Push(child2)

	expected := []struct {
		paint Paint
		typ   Paint
	}{
		{scene, Scene{}},
		{child1, Scene{}},
		{child11, Shape{}},
		{child2, Shape{}},
	}
	i := 0

	accessor := AccessorNew()
	_ = accessor.Set(scene, func(paint Paint) bool {
		assert.IsType(t, expected[i].typ, paint)
		assert.Equal(t, expected[i].paint.paint(), paint.paint())
		i++
		return true
	})
	assert.Equal(t, len(expected), i, "all expected visited")
}
