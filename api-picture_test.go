package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
	"github.com/stretchr/testify/assert"
)

func TestPictureLoad(t *testing.T) {
	defer testInitTerm(t)()

	picture := PictureNew()
	_ = picture.LoadData(data.LottieEllipse, "lottie+json", "")
	width, height, _ := picture.GetSize()
	assert.Equal(t, float32(100), width)
	assert.Equal(t, float32(50), height)
}
