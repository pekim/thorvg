package thorvg

import "github.com/ebitengine/purego"

/*
Accessor ia a structure representing an object that enables iterating through a scene's descendents.
*/
type Accessor uintptr

// Creates a new accessor object.
//
// Return: A new accessor object.
//
// Since: 1.0
func AccessorNew() Accessor {
	return Accessor(tvg_accessor_new())
}

// Deletes the given accessor object.
//
//   - accessor : The accessor object to be deleted.
//
// Since: 1.0
func (accessor Accessor) Del() error {
	return tvg_accessor_del(uintptr(accessor)).error()
}

// Sets the paint of the accessor then iterates through its descendents.
//
// Iterates through all descendents of the scene passed through the paint argument while calling func on each and passing the data pointer to this function. When func returns false iteration stops and the function returns.
//
//   - accessor : An accessor object.
//   - paint : The scene object.
//   - func : A function pointer to the function that will be execute for each child.
//   - data : A void pointer to data that will be passed to the func.
//
// Since: 1.0
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

// Generate a unique ID (hash key) from a given name.
//
// This function computes a unique identifier value based on the provided string. You can use this to assign a unique ID to the Paint object.
//
//   - name : The input string to generate the unique identifier from.
//
// Return: The generated unique identifier value.
//
// Since: 1.0
func AccessorGenerateId(name string) uint {
	return uint(tvg_accessor_generate_id(name))
}

// Retrieve the original name string from a given unique ID.
//
// Returns the name associated with the specified identifier.
// This method is only valid when [Picture.SetAccessible] is set to true for the Picture associated with the given paint in [Accessor.Set] Otherwise, the name information may not be available.
//
//   - accessor : An accessor object.
//   - id : The unique identifier.
//
// Return: The corresponding name string, or nullptr if not found or unavailable.
//
// see [Accessor.GenerateId]
//
// see [Accessor.Set]
//
// see [Picture.SetAccessible]
//
// Note: This function is only available within Accessor callbacks registered via [Accessor.Set]
// Since: 1.1
func (accessor Accessor) GetName(id uint) string {
	return tvg_accessor_get_name(uintptr(accessor), uint32(id))
}
