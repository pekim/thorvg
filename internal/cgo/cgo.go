package cgo

/*
#cgo CXXFLAGS: -I${SRCDIR}/src
#cgo CXXFLAGS: -I${SRCDIR}/inc
#cgo CXXFLAGS: -I${SRCDIR}
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie
#cgo CXXFLAGS: -I${SRCDIR}/src/bindings/capi
#cgo CXXFLAGS: -I${SRCDIR}/src/renderer
#cgo CXXFLAGS: -I${SRCDIR}/src/renderer/sw_engine
#cgo CXXFLAGS: -I${SRCDIR}/src/renderer/gl_engine
#cgo CXXFLAGS: -I${SRCDIR}/src/common
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/svg
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/ttf
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/api
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/include
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/ecma/base
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/ecma/builtin-objects/typedarray
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/ecma/builtin-objects
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/ecma/operations
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/jcontext
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/jmem
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/jrt
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/lit
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/parser/js
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/parser/regexp
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/lottie/jerryscript/jerry-core/vm
#cgo CXXFLAGS: -I${SRCDIR}/src/loaders/raw
#cgo CXXFLAGS: -I${SRCDIR}/src/savers
#cgo CXXFLAGS: -fvisibility=hidden
#cgo CXXFLAGS: -D_GLIBCXX_ASSERTIONS=1
#cgo CXXFLAGS: -D_FILE_OFFSET_BITS=64
#cgo CXXFLAGS: -Wall
#cgo CXXFLAGS: -Winvalid-pch
#cgo CXXFLAGS: -std=c++14
#cgo CXXFLAGS: -O3
#cgo CXXFLAGS: -g
#cgo CXXFLAGS: -fPIC
#cgo CXXFLAGS: -fopenmp
#cgo CXXFLAGS: -DTVG_EXPORT
#cgo CXXFLAGS: -DTVG_BUILD
#cgo CXXFLAGS: -mavx
#cgo CXXFLAGS: -Wno-unknown-pragmas
#cgo CXXFLAGS: -Wdouble-promotion
#cgo CXXFLAGS: -fno-exceptions
#cgo CXXFLAGS: -fno-rtti
#cgo CXXFLAGS: -fno-stack-protector
#cgo CXXFLAGS: -fno-asynchronous-unwind-tables
#cgo CXXFLAGS: -Woverloaded-virtual
#cgo CXXFLAGS: -DTHORVG_GL_TARGET_GL=1

#cgo LDFLAGS: -Wl,--as-needed
#cgo LDFLAGS: -Wl,--no-undefined
#cgo LDFLAGS: -Wl,-O1
#cgo LDFLAGS: -shared
#cgo LDFLAGS: -fPIC
#cgo LDFLAGS: -fopenmp
#cgo LDFLAGS: -lpthread

*/
import "C"
