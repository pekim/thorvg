package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NoError(t, Init())

	result := tvg_engine_init(4)
	assert.Equal(t, RESULT_SUCCESS, result)
	canvas := tvg_swcanvas_create(1)
	assert.NotZero(t, canvas)
}
