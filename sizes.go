package thorvg

// #include "thorvg_capi.h"
import "C"

import (
	"unsafe"
)

var result C.Tvg_Result
var colorspace C.Tvg_Colorspace
var engineOption C.Tvg_Engine_Option
var maskMethod C.Tvg_Mask_Method
var blendMethod C.Tvg_Blend_Method
var type_ C.Tvg_Type
var pathCommand C.Tvg_Path_Command
var strokeCap C.Tvg_Stroke_Cap
var strokeJoin C.Tvg_Stroke_Join
var strokeFill C.Tvg_Stroke_Fill
var fillRule C.Tvg_Fill_Rule
var textWrap C.Tvg_Text_Wrap

var sizes = []struct {
	name string
	c    int
	go_  uintptr
}{
	// structs
	{"ColorStop", C.sizeof_Tvg_Color_Stop, unsafe.Sizeof(ColorStop{})},
	{"Point", C.sizeof_Tvg_Point, unsafe.Sizeof(Point{})},
	{"Matrix", C.sizeof_Tvg_Matrix, unsafe.Sizeof(Matrix{})},

	// enum types
	{"Result", int(unsafe.Sizeof(result)), unsafe.Sizeof(RESULT_SUCCESS)},
	{"ColorSpace", int(unsafe.Sizeof(colorspace)), unsafe.Sizeof(COLORSPACE_ABGR8888)},
	{"EngineOption", int(unsafe.Sizeof(engineOption)), unsafe.Sizeof(ENGINE_OPTION_NONE)},
	{"MaskMethod", int(unsafe.Sizeof(maskMethod)), unsafe.Sizeof(MASK_METHOD_NONE)},
	{"BlendMethod", int(unsafe.Sizeof(blendMethod)), unsafe.Sizeof(BLEND_METHOD_NORMAL)},
	{"Type", int(unsafe.Sizeof(type_)), unsafe.Sizeof(TYPE_UNDEF)},
	{"PathCommand", int(unsafe.Sizeof(pathCommand)), unsafe.Sizeof(PATH_COMMAND_CLOSE)},
	{"StrokeCap", int(unsafe.Sizeof(strokeCap)), unsafe.Sizeof(STROKE_CAP_BUTT)},
	{"StrokeJoin", int(unsafe.Sizeof(strokeJoin)), unsafe.Sizeof(STROKE_JOIN_MITER)},
	{"StrokeFill", int(unsafe.Sizeof(strokeFill)), unsafe.Sizeof(STROKE_FILL_PAD)},
	{"FillRule", int(unsafe.Sizeof(fillRule)), unsafe.Sizeof(FILL_RULE_NON_ZERO)},
	{"TextWrap", int(unsafe.Sizeof(textWrap)), unsafe.Sizeof(TEXT_WRAP_NONE)},
}
