package thorvg

// #include "thorvg_capi.h"
// #include "api-accessor.h"
// #include <stdlib.h>
import "C"

import "unsafe"

// import "github.com/ebitengine/purego"

/*
Accessor ia a structure representing an object that enables iterating through a scene's descendents.
*/
type Accessor struct {
	accessor C.Tvg_Accessor
}

/*
AccessorNew creates a new accessor object.

	@return A new accessor object.

	@since 1.0
*/
func AccessorNew() Accessor {
	return Accessor{C.tvg_accessor_new()}
}

/*
Del deletes the given accessor object.

	@param[in] accessor The accessor object to be deleted.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Accessor pointer.

	@since 1.0
*/
func (accessor Accessor) Del() error {
	result := C.tvg_accessor_del(accessor.accessor)
	return resultError(result)
}

/*
Set sets the paint of the accessor then iterates through its descendents.

Iterates through all descendents of the scene passed through the paint argument
while calling func on each and passing the data pointer to this function. When
func returns false iteration stops and the function returns.

	@param[in] accessor A Tvg_Accessor pointer to the accessor object.
	@param[in] paint A Tvg_Paint pointer to the scene object.
	@param[in] func A function pointer to the function that will be execute for each child.
	@param[in] data A void pointer to data that will be passed to the func.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Accessor, Tvg_Paint, or function pointer.

	@since 1.0
*/
func (accessor Accessor) Set(paint Paint, func_ func(paint Paint) bool) error {
	accessorCallback = func_
	result := C.tvg_accessor_set(accessor.accessor, paint.paint(), (*[0]byte)(C.c_accessor_callback), nil)
	return resultError(result)
}

//export go_accessor_callback
func go_accessor_callback(cPaint C.Tvg_Paint, _data unsafe.Pointer) C.bool {
	paint, ok := newPaint(cPaint)
	if !ok {
		// This should never occur, so provide a non-specific Paint type.
		paint = paintCommon{paint_: cPaint}
	}

	return C.bool(accessorCallback(paint))
}

var accessorCallback func(paint Paint) bool

/*
AccessorGenerateId generates a unique ID (hash key) from a given name.

This function computes a unique identifier value based on the provided string.
You can use this to assign a unique ID to the Paint object.

	@param[in] name The input string to generate the unique identifier from.

	@return The generated unique identifier value.

	@since 1.0
*/
func AccessorGenerateId(name string) uint {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return uint(C.tvg_accessor_generate_id(cName))
}
