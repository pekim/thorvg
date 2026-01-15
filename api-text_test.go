package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	text := TextNew()
	_ = FontLoadData("DejaVuSans", data.DejaVuSans, "")
	_ = text.SetFont("DejaVuSans")
}

func TestFontLoad(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	_ = FontLoad("example/data/DejaVuSans.ttf")
	_ = FontUnload("example/data/DejaVuSans.ttf")
}

func TestFontLoadData(t *testing.T) {
	SetErrorHandler(func(err ResultError) { assert.Fail(t, err.Error()) })
	_ = EngineInit(2)
	defer func() { _ = EngineTerm() }()

	_ = FontLoadData("DejaVuSans", data.DejaVuSans, "")
	_ = FontUnloadData("DejaVuSans")
}
