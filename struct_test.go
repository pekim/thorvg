package thorvg

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStructs asserts that the Go structs types are the same size and
// shape as their C counterparts.
func TestStructs(t *testing.T) {
	for _, structs := range testStructs {
		cType := reflect.ValueOf(structs.c).Type()
		goType := reflect.ValueOf(structs.go_).Type()

		t.Run(goType.Name(), func(t *testing.T) {
			assert.Equal(t, cType.Size(), goType.Size())

			cFields := reflect.VisibleFields(cType)
			goFields := reflect.VisibleFields(goType)
			if goFields[0].Type.Name() == "HostLayout" {
				// discard HostLayout field
				goFields = goFields[1:]
			}

			for i, goField := range goFields {
				t.Run(goField.Name, func(t *testing.T) {
					cField := cFields[i]
					assert.Equal(t, cField.Offset, goField.Offset)           // offset
					assert.Equal(t, cField.Type.Size(), goField.Type.Size()) // size
				})
			}
		})
	}
}
