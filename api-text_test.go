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
