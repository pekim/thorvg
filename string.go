package thorvg

import (
	"math"
	"unsafe"

	"golang.org/x/sys/unix"
)

// func cString(s string) (uintptr, func()) {
// 	cStr := malloc(len(s) + 1)
// 	cStrSlice := unsafe.Slice(cStr, len(s)+1)
// 	copy(cStrSlice, []byte(s))
// 	cStrSlice[len(s)] = 0
// 	return uintptr(unsafe.Pointer(cStr)), func() { free(cStr) }
// }

// goString returns a Go string for a null terminated C string.
//
// The Go string is a copy of the C bytes.
func goString(cStr *byte) string {
	if cStr == nil {
		return ""
	}
	return unix.ByteSliceToString(unsafe.Slice(cStr, math.MaxInt))
}
