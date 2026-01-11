package thorvg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/ebitengine/purego"
)

func libraryFilepath() (string, error) {
	filepath := filepath.Join(os.TempDir(), fmt.Sprintf("libthorvg-%s.so", sharedObjectHash))

	// Check if the file exists. If it does, don't create it.
	//
	// The code would be simpler, and very nearly as quick, if the check were omitted and the
	// file written every time. However that causes problems if multiple applications are run
	// concurrently. The replaced file becomes unavailable, and an segment violation results.
	_, err := os.Stat(filepath)
	if err == nil {
		return filepath, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("failed to check existence of shared object file %q : %w", filepath, err)
	}

	err = os.WriteFile(filepath, sharedObject, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to write shared object file %q : %w", filepath, err)
	}

	return filepath, nil
}

// Engine API
var tvg_engine_init func(threads int) Result
var tvg_engine_term func() Result
var tvg_engine_version func(major *uint32, minor *uint32, micro *uint32, version **byte) Result

// Canvas API
var tvg_swcanvas_create func(option EngineOption) uintptr
var tvg_swcanvas_set_target func(canvas uintptr, buffer *byte, stride uint32, w uint32, h uint32, cs ColorSpace) Result
var tvg_glcanvas_create func() uintptr
var tvg_glcanvas_set_target func(canvas uintptr, display unsafe.Pointer, surface unsafe.Pointer, context unsafe.Pointer, id int32, w uint32, h uint32, cs ColorSpace) Result
var tvg_wgcanvas_create func() uintptr
var tvg_wgcanvas_set_target func(canvas uintptr, device unsafe.Pointer, instance unsafe.Pointer, target unsafe.Pointer, w uint32, h uint32, cs ColorSpace, typ int32) Result
var tvg_canvas_destroy func(canvas uintptr) Result
var tvg_canvas_push func(canvas uintptr, paint uintptr) Result
var tvg_canvas_push_at func(canvas uintptr, target uintptr, paint uintptr) Result
var tvg_canvas_remove func(canvas uintptr, paint uintptr) Result
var tvg_canvas_update func(canvas uintptr) Result
var tvg_canvas_draw func(canvas uintptr, clear_ bool) Result
var tvg_canvas_sync func(canvas uintptr) Result
var tvg_canvas_set_viewport func(canvas uintptr, x int32, y int32, w int32, h int32) Result

// Shape API
var tvg_shape_new func() uintptr
var tvg_shape_append_rect func(paint uintptr, x float32, y float32, w float32, h float32, rx float32, ry float32, cw bool) Result
var tvg_shape_append_circle func(paint uintptr, cx float32, cy float32, rx float32, ry float32, cw bool) Result
var tvg_shape_set_fill_color func(paint uintptr, r uint8, g uint8, b uint8, a uint8) Result
var tvg_shape_set_gradient func(paint uintptr, grad uintptr) Result
var tvg_shape_get_gradient func(paint uintptr, grad *uintptr) Result

// Gradient API
var tvg_radial_gradient_new func() uintptr
var tvg_radial_gradient_set func(grad uintptr, cx float32, cy float32, r float32, fx float32, fy float32, fr float32) Result
var tvg_radial_gradient_get func(grad uintptr, cx *float32, cy *float32, r *float32, fx *float32, fy *float32, fr *float32) Result
var tvg_gradient_set_color_stops func(grad uintptr, color_stop *ColorStop, cnt uint32) Result
var tvg_gradient_get_color_stops func(grad uintptr, color_stop **ColorStop, cnt *uint32) Result

// Text API
var tvg_text_new func() uintptr
var tvg_text_set_font func(text uintptr, name string) Result
var tvg_text_set_size func(text uintptr, size float32) Result
var tvg_text_set_text func(text uintptr, utf8 string) Result
var tvg_text_align func(text uintptr, x float32, y float32) Result
var tvg_text_layout func(text uintptr, w float32, h float32) Result
var tvg_text_wrap_mode func(text uintptr, mode TextWrap) Result
var tvg_text_spacing func(text uintptr, letter float32, line float32) Result
var tvg_text_set_italic func(text uintptr, shear float32) Result
var tvg_text_set_outline func(text uintptr, width float32, r uint8, g uint8, b uint8) Result
var tvg_text_set_color func(text uintptr, r uint8, g uint8, b uint8) Result

// var tvg_text_set_gradient func(text uintptr, gradient Gradient) Result
var tvg_font_load func(path string) Result
var tvg_font_load_data func(name string, data *byte, size uint32, mimetype string, copy_ bool) Result
var tvg_font_unload func(path string) Result

func initLibThorvg() error {
	filepath, err := libraryFilepath()
	if err != nil {
		return err
	}

	lib, err := purego.Dlopen(filepath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return err
	}

	// Engine API
	purego.RegisterLibFunc(&tvg_engine_init, lib, "tvg_engine_init")
	purego.RegisterLibFunc(&tvg_engine_term, lib, "tvg_engine_term")
	purego.RegisterLibFunc(&tvg_engine_version, lib, "tvg_engine_version")

	// Canvas API
	purego.RegisterLibFunc(&tvg_swcanvas_create, lib, "tvg_swcanvas_create")
	purego.RegisterLibFunc(&tvg_swcanvas_set_target, lib, "tvg_swcanvas_set_target")
	purego.RegisterLibFunc(&tvg_glcanvas_create, lib, "tvg_glcanvas_create")
	purego.RegisterLibFunc(&tvg_glcanvas_set_target, lib, "tvg_glcanvas_set_target")
	purego.RegisterLibFunc(&tvg_wgcanvas_create, lib, "tvg_wgcanvas_create")
	purego.RegisterLibFunc(&tvg_wgcanvas_set_target, lib, "tvg_wgcanvas_set_target")
	purego.RegisterLibFunc(&tvg_canvas_destroy, lib, "tvg_canvas_destroy")
	purego.RegisterLibFunc(&tvg_canvas_push, lib, "tvg_canvas_push")
	purego.RegisterLibFunc(&tvg_canvas_push_at, lib, "tvg_canvas_push_at")
	purego.RegisterLibFunc(&tvg_canvas_remove, lib, "tvg_canvas_remove")
	purego.RegisterLibFunc(&tvg_canvas_update, lib, "tvg_canvas_update")
	purego.RegisterLibFunc(&tvg_canvas_draw, lib, "tvg_canvas_draw")
	purego.RegisterLibFunc(&tvg_canvas_sync, lib, "tvg_canvas_sync")
	purego.RegisterLibFunc(&tvg_canvas_set_viewport, lib, "tvg_canvas_set_viewport")

	// Shape API
	purego.RegisterLibFunc(&tvg_shape_new, lib, "tvg_shape_new")
	purego.RegisterLibFunc(&tvg_shape_append_rect, lib, "tvg_shape_append_rect")
	purego.RegisterLibFunc(&tvg_shape_append_circle, lib, "tvg_shape_append_circle")
	purego.RegisterLibFunc(&tvg_shape_set_fill_color, lib, "tvg_shape_set_fill_color")
	purego.RegisterLibFunc(&tvg_shape_set_gradient, lib, "tvg_shape_set_gradient")
	purego.RegisterLibFunc(&tvg_shape_get_gradient, lib, "tvg_shape_get_gradient")

	// Gradient API
	purego.RegisterLibFunc(&tvg_radial_gradient_new, lib, "tvg_radial_gradient_new")
	purego.RegisterLibFunc(&tvg_radial_gradient_set, lib, "tvg_radial_gradient_set")
	purego.RegisterLibFunc(&tvg_radial_gradient_get, lib, "tvg_radial_gradient_get")
	purego.RegisterLibFunc(&tvg_gradient_set_color_stops, lib, "tvg_gradient_set_color_stops")
	purego.RegisterLibFunc(&tvg_gradient_get_color_stops, lib, "tvg_gradient_get_color_stops")

	// Text API
	purego.RegisterLibFunc(&tvg_text_new, lib, "tvg_text_new")
	purego.RegisterLibFunc(&tvg_text_set_font, lib, "tvg_text_set_font")
	purego.RegisterLibFunc(&tvg_text_set_size, lib, "tvg_text_set_size")
	purego.RegisterLibFunc(&tvg_text_set_text, lib, "tvg_text_set_text")
	purego.RegisterLibFunc(&tvg_text_align, lib, "tvg_text_align")
	purego.RegisterLibFunc(&tvg_text_layout, lib, "tvg_text_layout")
	purego.RegisterLibFunc(&tvg_text_wrap_mode, lib, "tvg_text_wrap_mode")
	purego.RegisterLibFunc(&tvg_text_spacing, lib, "tvg_text_spacing")
	purego.RegisterLibFunc(&tvg_text_set_italic, lib, "tvg_text_set_italic")
	purego.RegisterLibFunc(&tvg_text_set_outline, lib, "tvg_text_set_outline")
	purego.RegisterLibFunc(&tvg_text_set_color, lib, "tvg_text_set_color")
	// tvg_text_set_gradient is dependent on the linux struct argument support
	// in https://github.com/ebitengine/purego/pull/361.
	// purego.RegisterLibFunc(&tvg_text_set_gradient, lib, "tvg_text_set_gradient")
	purego.RegisterLibFunc(&tvg_font_load, lib, "tvg_font_load")
	purego.RegisterLibFunc(&tvg_font_load_data, lib, "tvg_font_load_data")
	purego.RegisterLibFunc(&tvg_font_unload, lib, "tvg_font_unload")

	return nil
}
