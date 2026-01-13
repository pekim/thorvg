package thorvg

import (
	"unsafe"
)

/*
Gradient is a structure representing a gradient fill of a Paint object.
*/
type Gradient interface {
	gradient() uintptr
}

type gradientCommon struct {
	gradient_ uintptr
}

func (gradient gradientCommon) gradient() uintptr {
	return gradient.gradient_
}

type LinearGradient struct {
	gradientCommon
}

type RadialGradient struct {
	gradientCommon
}

/*
LinearGradientNew creates a new linear gradient object.

	@return A new linear gradient object.
*/
func LinearGradientNew() LinearGradient {
	return LinearGradient{
		gradientCommon{gradient_: tvg_linear_gradient_new()},
	}
}

/*
RadialGradientNew Creates a new radial gradient object.

	@return A new radial gradient object.
*/
func RadialGradientNew() RadialGradient {
	return RadialGradient{
		gradientCommon{gradient_: tvg_radial_gradient_new()},
	}
}

/*
Set sets the linear gradient bounds.
The bounds of the linear gradient are defined as a surface constrained by two parallel lines crossing
the given points (@p x1, @p y1) and (@p x2, @p y2), respectively. Both lines are perpendicular to the line linking
(@p x1, @p y1) and (@p x2, @p y2).

	@param[in] grad The Tvg_Gradient object of which bounds are to be set.
	@param[in] x1 The horizontal coordinate of the first point used to determine the gradient bounds.
	@param[in] y1 The vertical coordinate of the first point used to determine the gradient bounds.
	@param[in] x2 The horizontal coordinate of the second point used to determine the gradient bounds.
	@param[in] y2 The vertical coordinate of the second point used to determine the gradient bounds.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
	@note In case the first and the second points are equal, an object is filled with a single color using the last color specified in the tvg_gradient_set_color_stops().
	@see tvg_gradient_set_color_stops()
*/
func (gradient LinearGradient) Set(x1 float32, y1 float32, x2 float32, y2 float32) error {
	return tvg_linear_gradient_set(gradient.gradient_, x1, y1, x2, y2).error()
}

/*
Get gets the linear gradient bounds.
The bounds of the linear gradient are defined as a surface constrained by two parallel lines crossing
the given points (@p x1, @p y1) and (@p x2, @p y2), respectively. Both lines are perpendicular to the line linking
(@p x1, @p y1) and (@p x2, @p y2).

	@param[in] grad The Tvg_Gradient object of which to get the bounds.
	@param[out] x1 The horizontal coordinate of the first point used to determine the gradient bounds.
	@param[out] y1 The vertical coordinate of the first point used to determine the gradient bounds.
	@param[out] x2 The horizontal coordinate of the second point used to determine the gradient bounds.
	@param[out] y2 The vertical coordinate of the second point used to determine the gradient bounds.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
*/
func (gradient LinearGradient) Get() (float32, float32, float32, float32, error) {
	var x1, y1 float32
	var x2, y2 float32
	result := tvg_linear_gradient_get(gradient.gradient_, &x1, &y1, &x2, &y2)
	return x1, y1, x2, y2, result.error()
}

/*
RadialSet sets the radial gradient attributes.

The radial gradient is defined by the end circle with a center (@p cx, @p cy) and a radius @p r and
the start circle with a center/focal point (@p fx, @p fy) and a radius @p fr.
The gradient will be rendered such that the gradient stop at an offset of 100% aligns with the edge of the end circle
and the stop at an offset of 0% aligns with the edge of the start circle.

	@param[in] grad The Tvg_Gradient object of which bounds are to be set.
	@param[in] cx The horizontal coordinate of the center of the end circle.
	@param[in] cy The vertical coordinate of the center of the end circle.
	@param[in] r The radius of the end circle.
	@param[in] fx The horizontal coordinate of the center of the start circle.
	@param[in] fy The vertical coordinate of the center of the start circle.
	@param[in] fr The radius of the start circle.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer or the radius @p r or @p fr value is negative.

	@note In case the radius @p r is zero, an object is filled with a single color using the last color specified in the specified in the tvg_gradient_set_color_stops().
	@note In case the focal point (@p fx and @p fy) lies outside the end circle, it is projected onto the edge of the end circle.
	@note If the start circle doesn't fully fit inside the end circle (after possible repositioning), the @p fr is reduced accordingly.
	@note By manipulating the position and size of the focal point, a wide range of visual effects can be achieved, such as directing
				the gradient focus towards a specific edge or enhancing the depth and complexity of shading patterns.
				If a focal effect is not desired, simply align the focal point (@p fx and @p fy) with the center of the end circle (@p cx and @p cy)
				and set the radius (@p fr) to zero. This will result in a uniform gradient without any focal variations.

	@see tvg_gradient_set_color_stops()
*/
func (gradient RadialGradient) RadialSet(cx float32, cy float32, r float32, fx float32, fy float32, fr float32) error {
	return tvg_radial_gradient_set(gradient.gradient_, cx, cy, r, fx, fy, fr).error()
}

/*
RadialGet gets radial gradient attributes.

	@param[in] grad The Tvg_Gradient object of which to get the gradient attributes.
	@param[out] cx The horizontal coordinate of the center of the end circle.
	@param[out] cy The vertical coordinate of the center of the end circle.
	@param[out] r The radius of the end circle.
	@param[out] fx The horizontal coordinate of the center of the start circle.
	@param[out] fy The vertical coordinate of the center of the start circle.
	@param[out] fr The radius of the start circle.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.

	@see tvg_radial_gradient_set()
*/
func (gradient RadialGradient) RadialGet() (float32, float32, float32, float32, float32, float32, error) {
	var cx, cy, r float32
	var fx, fy, fr float32
	result := tvg_radial_gradient_get(gradient.gradient_, &cx, &cy, &r, &fx, &fy, &fr)
	return cx, cy, r, fx, fy, fr, result.error()
}

/*
SetColorStops Sets the parameters of the colors of the gradient and their position.

	@param[in] grad The Tvg_Gradient object of which the color information is to be set.
	@param[in] color_stop An array of Tvg_Color_Stop data structure.
	@param[in] cnt The size of the @p color_stop array equal to the colors number used in the gradient.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
*/
func (gradient gradientCommon) SetColorStops(colorStops []ColorStop) error {
	if len(colorStops) == 0 {
		return tvg_gradient_set_color_stops(gradient.gradient_, nil, uint32(len(colorStops))).error()
	}
	return tvg_gradient_set_color_stops(gradient.gradient_, &colorStops[0], uint32(len(colorStops))).error()
}

/*
GetColorStops gets the parameters of the colors of the gradient, their position and number

The function does not allocate any memory.

	@param[in] grad The Tvg_Gradient object of which to get the color information.
	@param[out] color_stop An array of Tvg_Color_Stop data structure.
	@param[out] cnt The size of the @p color_stop array equal to the colors number used in the gradient.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
*/
func (gradient gradientCommon) GetColorStops() ([]ColorStop, error) {
	var cColorStops *ColorStop
	var cCount uint32
	result := tvg_gradient_get_color_stops(gradient.gradient_, &cColorStops, &cCount)
	colorStops := make([]ColorStop, cCount)
	copy(colorStops, unsafe.Slice(cColorStops, cCount))
	return colorStops, result.error()
}

/*
SetSpread sets the Tvg_Stroke_Fill value, which specifies how to fill the area outside the gradient bounds.

	@param[in] grad The Tvg_Gradient object.
	@param[in] spread The FillSpread value.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
*/
func (gradient gradientCommon) SetSpread(spread StrokeFill) error {
	return tvg_gradient_set_spread(gradient.gradient_, spread).error()
}

/*
GetSpread gets the FillSpread value of the gradient object.

	@param[in] grad The Tvg_Gradient object.
	@param[out] spread The FillSpread value.
	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
*/
func (gradient gradientCommon) GetSpread() (StrokeFill, error) {
	var spread StrokeFill
	result := tvg_gradient_get_spread(gradient.gradient_, &spread)
	return spread, result.error()
}

/*
SetTransform sets the matrix of the affine transformation for the gradient object.
The augmented matrix of the transformation is expected to be given.

	@param[in] grad The Tvg_Gradient object to be transformed.
	@param[in] m The 3x3 augmented matrix.
	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr is passed as the argument.
*/
func (gradient gradientCommon) SetTransform(matrix Matrix) error {
	return tvg_gradient_set_transform(gradient.gradient_, &matrix).error()
}

/*
GetTransform gets the matrix of the affine transformation of the gradient object.
In case no transformation was applied, the identity matrix is set.

	@param[in] grad The Tvg_Gradient object of which to get the transformation matrix.
	@param[out] m The 3x3 augmented matrix.
	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr is passed as the argument.
*/
func (gradient gradientCommon) GetTransform() (Matrix, error) {
	var matrix Matrix
	result := tvg_gradient_get_transform(gradient.gradient_, &matrix)
	return matrix, result.error()
}

/*
GetType gets the unique value of the gradient instance indicating the instance type.

	@param[in] grad The Tvg_Gradient object of which to get the type value.
	@param[out] type The unique type of the gradient instance type.
	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.
	@since 1.0
*/
func (gradient gradientCommon) GetType() (Type, error) {
	var typ Type
	result := tvg_gradient_get_type(gradient.gradient_, &typ)
	return typ, result.error()
}

func (gradient gradientCommon) duplicate() gradientCommon {
	return gradientCommon{
		gradient_: tvg_gradient_duplicate(gradient.gradient_),
	}
}

/*
Duplicate duplicates the given Tvg_Gradient object.
Creates a new object and sets its all properties as in the original object.

	@param[in] grad The Tvg_Gradient object to be copied.
	@return A copied Tvg_Gradient object if succeed, @c nullptr otherwise.
*/
func (gradient LinearGradient) Duplicate() LinearGradient {
	return LinearGradient{gradientCommon: gradient.duplicate()}
}

/*
Duplicate duplicates the given Tvg_Gradient object.
Creates a new object and sets its all properties as in the original object.

	@param[in] grad The Tvg_Gradient object to be copied.
	@return A copied Tvg_Gradient object if succeed, @c nullptr otherwise.
*/
func (gradient RadialGradient) Duplicate() RadialGradient {
	return RadialGradient{gradientCommon: gradient.duplicate()}
}

/*
Del deletes the given gradient object.

	@param[in] grad The gradient object to be deleted.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
*/
func (gradient gradientCommon) Del() error {
	return tvg_gradient_del(gradient.gradient_).error()
}

func newGradient(gradient uintptr) (Gradient, bool) {
	var typ Type
	result := tvg_gradient_get_type(gradient, &typ)
	if result != RESULT_SUCCESS {
		return nil, false
	}

	switch typ {

	case TYPE_LINEAR_GRAD:
		return LinearGradient{gradientCommon: gradientCommon{gradient_: gradient}}, true

	case TYPE_RADIAL_GRAD:
		return RadialGradient{gradientCommon: gradientCommon{gradient_: gradient}}, true

	default:
		return nil, false
	}
}
