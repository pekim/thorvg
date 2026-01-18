package thorvg

// #cgo CFLAGS: -I${SRCDIR}/internal/cgo/src/bindings/capi
import "C"

import (
	_ "github.com/pekim/thorvg/internal/cgo" // required for C/C++ object files in the package
)
