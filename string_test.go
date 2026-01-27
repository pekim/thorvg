package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoString(t *testing.T) {
	assert.Equal(t, "", goString(nil))

	data := []byte{'a', 'b', 'c', 0, 0, 0}
	assert.Equal(t, "abc", goString(&data[0]))
}
