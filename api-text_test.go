package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	defer testInitTerm(t)()

	text := TextNew()
	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, text.SetFont("DejaVuSans"))

	assert.NoError(t, text.SetColor(0, 0, 0))
	assert.NoError(t, text.SetSize(18))

	gradient := RadialGradientNew()
	assert.NoError(t, gradient.RadialSet(1, 2, 3, 4, 5, 6))
	assert.NoError(t, text.SetGradient(gradient))
}

func TestFontLoad(t *testing.T) {
	defer testInitTerm(t)()

	assert.NoError(t, FontLoad("example/data/DejaVuSans.ttf"))
	assert.NoError(t, FontUnload("example/data/DejaVuSans.ttf"))
}

func TestFontLoadData(t *testing.T) {
	defer testInitTerm(t)()

	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, FontUnloadData("DejaVuSans"))
}

func TestTextMetrics(t *testing.T) {
	defer testInitTerm(t)()

	text := TextNew()
	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, text.SetFont("DejaVuSans"))
	assert.NoError(t, text.SetSize(12))

	metrics, err := text.GetTextMetrics()
	assert.NoError(t, err)
	assert.Equal(t, TextMetrics{
		Ascent:  14.8515625,
		Descent: -3.7734375,
		Linegap: 0,
		Advance: 18.625,
	}, metrics)
}

func TestGlyphMetrics(t *testing.T) {
	defer testInitTerm(t)()

	text := TextNew()
	assert.NoError(t, FontLoadData("DejaVuSans", data.DejaVuSans, ""))
	assert.NoError(t, text.SetFont("DejaVuSans"))
	assert.NoError(t, text.SetSize(12))

	metrics, err := text.GetGlyphMetrics('A')
	assert.NoError(t, err)
	assert.Equal(t, GlyphMetrics{
		Advance: 10.9453125,
		Bearing: 0.125,
		Min:     Point{X: 0.125, Y: 0},
		Max:     Point{X: 10.8125, Y: 11.6640625},
	}, metrics)
}
