package thorvg

import "unsafe"

/*
slice is like unsafe.Slice, except that it also copies the underlying data.

It is used with arrays returned from thorvg.
*/
func slice[T any, I ~int | ~uint32](data *T, length I) []T {
	if data == nil && length == 0 {
		return nil
	}

	slice := make([]T, length)
	copy(slice, unsafe.Slice(data, length))
	return slice
}
