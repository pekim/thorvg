package thorvg

import (
	"math"
	"unsafe"

	"golang.org/x/sys/unix"
)

// goString returns a Go string for a null terminated C string.
//
// The Go string is a copy of the C bytes.
func goString(cStr *byte) string {
	if cStr == nil {
		return ""
	}
	return unix.ByteSliceToString(unsafe.Slice(cStr, math.MaxInt))
}
