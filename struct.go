package thorvg

import "structs"

/*
ColorStop is a data structure storing the information about the color and its relative position inside the gradient bounds.
*/
type ColorStop struct {
	_ structs.HostLayout

	offset float32 // The relative position of the color.
	r      uint8   // The red color channel value in the range [0 ~ 255].
	g      uint8   // The green color channel value in the range [0 ~ 255].
	b      uint8   // The blue color channel value in the range [0 ~ 255].
	a      uint8   // The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
}

/*
Point is a data structure representing a point in two-dimensional space.
*/
type Point struct {
	_ structs.HostLayout

	x, y float32
}

/*
Matrix is a data structure representing a three-dimensional matrix.

The elements e11, e12, e21 and e22 represent the rotation matrix, including the scaling factor.
The elements e13 and e23 determine the translation of the object along the x and y-axis, respectively.
The elements e31 and e32 are set to 0, e33 is set to 1.
*/
type Matrix struct {
	_ structs.HostLayout

	e11, e12, e13 float32
	e21, e22, e23 float32
	e31, e32, e33 float32
}
