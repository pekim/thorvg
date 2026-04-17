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

/*
TextMetrics describes the font metrics of a text object.

Provides the basic vertical layout metrics used for text rendering,
such as ascent, descent, and line spacing (linegap).

	@see tvg_text_get_text_metrics()
	@note Experimental API
*/
type TextMetrics struct {
	_ structs.HostLayout

	Ascent  float32 ///< Distance from the baseline to the top of the highest glyph (usually positive).
	Descent float32 ///< Distance from the baseline to the bottom of the lowest glyph (usually negative, as in TTF).
	Linegap float32 ///< Additional spacing recommended between lines (leading).
	Advance float32 ///< The total vertical advance between lines of text: ascent - descent + linegap (i.e., ascent + |descent| + linegap when descent is negative).
}

/*
GlyphMetrics describes the layout metrics of a glyph.

Provides the basic layout metrics used for positioning an individual glyph,
including its advance along the baseline direction, bearing relative to the
inline axis origin, and its bounding box in local glyph space.

The advance value represents the distance the pen position moves along the
baseline (inline direction), regardless of whether the text is laid out
horizontally or vertically.

The bounding box is defined in the glyph’s local coordinate space and is
independent of any layout direction or transformation.

	@see tvg_text_get_glyph_metrics()
	@note Experimental API
*/
type GlyphMetrics struct {
	_ structs.HostLayout

	Advance float32 ///< The advance distance along the baseline (inline) direction.
	Bearing float32 ///< The bearing from the origin to the glyph’s visible bound along the inline-start direction.
	Min     Point   ///< The minimum point of the glyph bounding box in local space.
	Max     Point   ///< The maximum point of the glyph bounding box in local space.
}
