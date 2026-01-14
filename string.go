package thorvg

import (
	"math"
	"unsafe"
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

	// Look for terminating 0, to find the length of the string.
	length := 0
	str := unsafe.Slice(cStr, math.MaxInt)
	for i, b := range str {
		if b == 0 {
			length = i
			break
		}
	}

	// goBytes := make([]byte, length)
	// copy(goBytes, str)
	return string(slice(cStr, length))
}
