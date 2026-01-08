package cvalidation

// #include "../lib/thorvg_capi.h"
import "C"

import (
	"unsafe"

	"github.com/pekim/thorvg/internal"
)

var enum internal.Result

var sizes = []struct {
	name string
	c    int
	go_  uintptr
}{
	// structs
	{"ColorStop", C.sizeof_Tvg_Color_Stop, unsafe.Sizeof(internal.ColorStop{})},
	{"Point", C.sizeof_Tvg_Point, unsafe.Sizeof(internal.Point{})},
	{"Matrix", C.sizeof_Tvg_Matrix, unsafe.Sizeof(internal.Matrix{})},

	// enum type
	{"enum", C.sizeof_Tvg_Result, unsafe.Sizeof(enum)},
}
