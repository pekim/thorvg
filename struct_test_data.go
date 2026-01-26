package thorvg

// #include "thorvg_capi.h"
import "C"

// This cannot be in the test file where it is used because use of cgo
// is not permitted in test files.
//
// And it is not in the file where the struct types are defined, because
// that would force the use of cgo on a user of the library.
var testStructs = []struct {
	c   any
	go_ any
}{
	{
		c:   C.Tvg_Color_Stop{},
		go_: ColorStop{},
	},
	{
		c:   C.Tvg_Point{},
		go_: Point{},
	},
	{
		c:   C.Tvg_Matrix{},
		go_: Matrix{},
	},
}
