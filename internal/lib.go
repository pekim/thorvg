package internal

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

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

var tvg_engine_init func(threads int) Result
var tvg_swcanvas_create func(op int) Canvas

func Init() error {
	filepath, err := libraryFilepath()
	if err != nil {
		return err
	}

	lib, err := purego.Dlopen(filepath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&tvg_engine_init, lib, "tvg_engine_init")
	purego.RegisterLibFunc(&tvg_swcanvas_create, lib, "tvg_swcanvas_create")

	return nil
}
