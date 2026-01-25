package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInitTerm(t *testing.T) func() {
	t.Helper()

	SetErrorHandler(func(err ResultError) {
		assert.Fail(t, err.Error())
	})

	err := EngineInit(0)
	assert.NoError(t, err)

	return func() {
		err := EngineTerm()
		assert.NoError(t, err)
	}
}
