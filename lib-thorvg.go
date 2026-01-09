package thorvg

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/ebitengine/purego"
)

//go:embed internal/lib/libthorvg-1.so.1.0.0
var sharedObject []byte

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
var tvg_canvas_destroy func(canvas uintptr) Result
var tvg_canvas_push func(canvas uintptr, paint uintptr) Result
var tvg_canvas_draw func(canvas uintptr, clear_ bool) Result
var tvg_canvas_sync func(canvas uintptr) Result

// Shape API
var tvg_shape_new func() uintptr
var tvg_shape_append_rect func(paint uintptr, x float32, y float32, w float32, h float32, rx float32, ry float32, cw bool) Result
var tvg_shape_set_fill_color func(paint uintptr, r uint8, g uint8, b uint8, a uint8) Result

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
	purego.RegisterLibFunc(&tvg_canvas_destroy, lib, "tvg_canvas_destroy")
	purego.RegisterLibFunc(&tvg_canvas_push, lib, "tvg_canvas_push")
	purego.RegisterLibFunc(&tvg_canvas_draw, lib, "tvg_canvas_draw")
	purego.RegisterLibFunc(&tvg_canvas_sync, lib, "tvg_canvas_sync")

	// Shape API
	purego.RegisterLibFunc(&tvg_shape_new, lib, "tvg_shape_new")
	purego.RegisterLibFunc(&tvg_shape_append_rect, lib, "tvg_shape_append_rect")
	purego.RegisterLibFunc(&tvg_shape_set_fill_color, lib, "tvg_shape_set_fill_color")

	return nil
}
