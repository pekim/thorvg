package cvalidation

// #include "../lib/thorvg_capi.h"
import "C"

import (
	"unsafe"

	"github.com/pekim/thorvg"
)

var enum thorvg.Result

var sizes = []struct {
	name string
	c    int
	go_  uintptr
}{
	// structs
	{"ColorStop", C.sizeof_Tvg_Color_Stop, unsafe.Sizeof(thorvg.ColorStop{})},
	{"Point", C.sizeof_Tvg_Point, unsafe.Sizeof(thorvg.Point{})},
	{"Matrix", C.sizeof_Tvg_Matrix, unsafe.Sizeof(thorvg.Matrix{})},

	// enum type
	{"enum", C.sizeof_Tvg_Result, unsafe.Sizeof(enum)},
}
