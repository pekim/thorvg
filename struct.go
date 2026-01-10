package thorvg

import "structs"

/*
ColorStop is a data structure storing the information about the color and its relative position inside the gradient bounds.
*/
type ColorStop struct {
	_ structs.HostLayout

	Offset float32 // The relative position of the color.
	R      uint8   // The red color channel value in the range [0 ~ 255].
	G      uint8   // The green color channel value in the range [0 ~ 255].
	B      uint8   // The blue color channel value in the range [0 ~ 255].
	A      uint8   // The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
}

/*
Point is a data structure representing a point in two-dimensional space.
*/
type Point struct {
	_ structs.HostLayout

	X, Y float32
}

/*
Matrix is a data structure representing a three-dimensional matrix.

The elements e11, e12, e21 and e22 represent the rotation matrix, including the scaling factor.
The elements e13 and e23 determine the translation of the object along the x and y-axis, respectively.
The elements e31 and e32 are set to 0, e33 is set to 1.
*/
type Matrix struct {
	_ structs.HostLayout

	E11, E12, E13 float32
	E21, E22, E23 float32
	E31, E32, E33 float32
}
