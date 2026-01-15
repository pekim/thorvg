package thorvg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInitTerm(t *testing.T) func() {
	t.Helper()
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	return func() { _ = EngineTerm() }
}
