package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
)

func TestText(t *testing.T) {
	defer testInitTerm(t)()

	text := TextNew()
	_ = FontLoadData("DejaVuSans", data.DejaVuSans, "")
	_ = text.SetFont("DejaVuSans")

	if err := text.SetColor(0, 0, 0); err != nil {
		panic(err)
	}
	if err := text.SetSize(18); err != nil {
		panic(err)
	}

	gradient := RadialGradientNew()
	_ = gradient.RadialSet(1, 2, 3, 4, 5, 6)
	_ = text.SetGradient(gradient)
}

func TestFontLoad(t *testing.T) {
	defer testInitTerm(t)()

	_ = FontLoad("example/data/DejaVuSans.ttf")
	_ = FontUnload("example/data/DejaVuSans.ttf")
}

func TestFontLoadData(t *testing.T) {
	defer testInitTerm(t)()

	_ = FontLoadData("DejaVuSans", data.DejaVuSans, "")
	_ = FontUnloadData("DejaVuSans")
}
