package thorvg

// #include "thorvg_capi.h"
import "C"

import (
	"unsafe"
)

var sizes = []struct {
	name string
	c    int
	go_  uintptr
}{
	// structs
	{"ColorStop", C.sizeof_Tvg_Color_Stop, unsafe.Sizeof(ColorStop{})},
	{"Point", C.sizeof_Tvg_Point, unsafe.Sizeof(Point{})},
	{"Matrix", C.sizeof_Tvg_Matrix, unsafe.Sizeof(Matrix{})},
}
