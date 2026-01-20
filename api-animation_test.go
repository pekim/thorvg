package thorvg

import (
	"testing"

	"github.com/pekim/thorvg/example/data"
	"github.com/stretchr/testify/assert"
)

func TestAnimation(t *testing.T) {
	defer testInitTerm(t)()

	animation := AnimationNew()
	picture := animation.GetPicture()
	_ = picture.LoadData(data.LottieGearsAnimation, "lottie+json", "")

	duration, _ := animation.GetDuration()
	assert.InEpsilon(t, float32(30.0), duration, 0.1)

	frameCount, _ := animation.GetTotalFrame()
	assert.InEpsilon(t, float32(900.0), frameCount, 0.1)

	begin, end, _ := animation.GetSegment()
	assert.Equal(t, float32(0.0), begin)
	assert.InEpsilon(t, float32(900.0), end, 0.1)
}
