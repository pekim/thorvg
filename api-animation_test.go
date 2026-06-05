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

func TestLottieAnimationAudioLayer(t *testing.T) {
	defer testInitTerm(t)()

	animation := LottieAnimationNew()
	picture := animation.GetPicture()
	_ = picture.LoadData(data.LottieAudioLayer, "lottie+json", "")
	var audioInfo *AudioInfo
	err := animation.SetAudioResolver(func(info AudioInfo) {
		audioInfo = &info
	})
	assert.NoError(t, err)

	SetErrorHandler(func(err ResultError) { panic(err) })
	err = animation.SetFrame(1)
	assert.NoError(t, err)

	assert.NotNil(t, audioInfo)
	assert.Equal(t, AudioInfo{
		Src:      "/audio/aud_1.mp3",
		MimeType: "",
		Size:     0,
		Offset:   0.016666668,
		Volume:   100,
		Active:   true,
		Embedded: false,
	}, *audioInfo)
}
