package c

// #include "../thorvg_capi.h"
import "C"

import (
	tvg "github.com/pekim/thorvg"
)

// This file is not in the main thorvg package,
// because that would force the use of cgo on a user of the library.

// This variable cannot be in the test file where it is used because
// the use of cgo is not permitted in test files.
var testStructs = []struct {
	c   any
	go_ any
}{
	{
		c:   C.Tvg_Color_Stop{},
		go_: tvg.ColorStop{},
	},
	{
		c:   C.Tvg_Point{},
		go_: tvg.Point{},
	},
	{
		c:   C.Tvg_Matrix{},
		go_: tvg.Matrix{},
	},
}
