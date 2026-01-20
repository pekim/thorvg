package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructSizes(t *testing.T) {
	for _, size := range structSizes {
		t.Run(size.name, func(t *testing.T) {
			assert.Equal(t, size.c, int(size.go_))
		})
	}
}
