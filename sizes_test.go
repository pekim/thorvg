package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStructs asserts that the Go structs types are the same size as their C counterparts.
//
// It's just a crude sanity check.
// It doesn't verify that all fields have the correct size and offset.
func TestStructSize(t *testing.T) {
	for _, size := range sizes {
		t.Run(size.name, func(t *testing.T) {
			assert.Equal(t, size.c, int(size.go_))
		})
	}
}
