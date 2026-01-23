package thorvg

import "github.com/ebitengine/purego"

/*
Accessor ia a structure representing an object that enables iterating through a scene's descendents.
*/
type Accessor uintptr

/*
AccessorNew creates a new accessor object.

	@return A new accessor object.

	@since 1.0
*/
func AccessorNew() Accessor {
	return Accessor(tvg_accessor_new())
}

/*
Del deletes the given accessor object.

	@param[in] accessor The accessor object to be deleted.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Accessor pointer.

	@since 1.0
*/
func (accessor Accessor) Del() error {
	return tvg_accessor_del(uintptr(accessor)).error()
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
	puregoFunc := purego.NewCallback(func(cPaint uintptr, _data uintptr) bool {
		paint, ok := newPaint(cPaint)
		if !ok {
			// This should never occur, so provide a non-specific Paint type.
			paint = paintCommon{paint_: cPaint}
		}

		return func_(paint)
	})

	return tvg_accessor_set(uintptr(accessor), paint.paint(), puregoFunc, 0).error()
}

/*
AccessorGenerateId generates a unique ID (hash key) from a given name.

This function computes a unique identifier value based on the provided string.
You can use this to assign a unique ID to the Paint object.

	@param[in] name The input string to generate the unique identifier from.

	@return The generated unique identifier value.

	@since 1.0
*/
func AccessorGenerateId(name string) uint {
	return uint(tvg_accessor_generate_id(name))
}
