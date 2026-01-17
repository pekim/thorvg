package api

/*
#cgo CFLAGS: -I${SRCDIR}/../thorvg-src/src/bindings/capi
#include "thorvg_capi.h"
*/
import "C"

import (
	"fmt"

	_ "github.com/pekim/thorvg/internal/cgo" // required for C/C++ object files in the package
)

func Temp() {
	fmt.Println(C.tvg_engine_init(2))
	canvas := C.tvg_swcanvas_create(C.TVG_ENGINE_OPTION_DEFAULT)
	fmt.Println("canvas", canvas)
}
