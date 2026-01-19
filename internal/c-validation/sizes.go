package cvalidation

// #include "../lib/thorvg_capi.h"
import "C"

import (
	"unsafe"

	tvg "github.com/pekim/thorvg"
)

var sizes = []struct {
	name string
	c    int
	go_  uintptr
}{
	// structs
	{"ColorStop", C.sizeof_Tvg_Color_Stop, unsafe.Sizeof(tvg.ColorStop{})},
	{"Point", C.sizeof_Tvg_Point, unsafe.Sizeof(tvg.Point{})},
	{"Matrix", C.sizeof_Tvg_Matrix, unsafe.Sizeof(tvg.Matrix{})},
}
