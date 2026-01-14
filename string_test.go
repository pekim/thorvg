package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoString(t *testing.T) {
	data := []byte{'a', 'b', 'c', 0, 0, 0}
	assert.Equal(t, "abc", goString(&data[0]))
}
