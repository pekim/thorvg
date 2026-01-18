package thorvg

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestSwCanvas(t *testing.T) {
// 	defer testInitTerm(t)()

// 	canvas := SwCanvasCreate(ENGINE_OPTION_DEFAULT)
// 	_ = canvas
// 	assert.Nil(t, canvas.Buffer())

// 	_ = canvas.SwSetTarget(16, 16, 16, COLORSPACE_ARGB8888)
// 	assert.Equal(t, 4*16*16, len(canvas.Buffer()))

// 	_ = canvas.Destroy()
// 	assert.Nil(t, canvas.Buffer())
// }
