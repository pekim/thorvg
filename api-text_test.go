package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	text := TextNew()
	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, text.SetFont("DejaVuSans"))
}

func TestFontLoad(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	assert.NoError(t, FontLoad("example/data/DejaVuSans.ttf"))
	assert.NoError(t, FontUnload("example/data/DejaVuSans.ttf"))
}

func TestFontLoadData(t *testing.T) {
	assert.NoError(t, EngineInit(2))
	defer func() { assert.NoError(t, EngineTerm()) }()

	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, FontUnloadData("DejaVuSans"))
}
