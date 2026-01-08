package internal

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/ebitengine/purego"
)

//go:embed lib/libthorvg-1.so.1.0.0
var sharedObject []byte

func libraryFilepath() (string, error) {
	filepath := filepath.Join(os.TempDir(), fmt.Sprintf("libthorvg-%s.so", sharedObjectHash))

	err := os.WriteFile(filepath, sharedObject, 0755)
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

	return nil
}
