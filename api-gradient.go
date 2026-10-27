package thorvg

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

type ConicGradient struct {
	gradientCommon
}

// Creates a new linear gradient object.
//
// Return: A new linear gradient object.
func LinearGradientNew() LinearGradient {
	return LinearGradient{
		gradientCommon{gradient_: tvg_linear_gradient_new()},
	}
}

// Creates a new radial gradient object.
//
// Return: A new radial gradient object.
func RadialGradientNew() RadialGradient {
	return RadialGradient{
		gradientCommon{gradient_: tvg_radial_gradient_new()},
	}
}

// Creates a new conic gradient object.
//
// Return: A new conic gradient object.
//
// Note: Experimental API
func ConicGradientNew() ConicGradient {
	return ConicGradient{
		gradientCommon{gradient_: tvg_conic_gradient_new()},
	}
}

// Sets the linear gradient bounds.
//
// The bounds of the linear gradient are defined as a surface constrained by two parallel lines crossing the given points (x1, y1) and (x2, y2), respectively. Both lines are perpendicular to the line linking (x1, y1) and (x2, y2).
//
//   - grad : The Tvg_Gradient object of which bounds are to be set.
//   - x1 : The horizontal coordinate of the first point used to determine the gradient bounds.
//   - y1 : The vertical coordinate of the first point used to determine the gradient bounds.
//   - x2 : The horizontal coordinate of the second point used to determine the gradient bounds.
//   - y2 : The vertical coordinate of the second point used to determine the gradient bounds.
//
// Note: In case the first and the second points are equal, an object is filled with a single color using the last color specified in the tvg_gradient_set_color_stops().
//
// see [GradientSetColorStops]
func (gradient LinearGradient) Set(x1 float32, y1 float32, x2 float32, y2 float32) error {
	return tvg_linear_gradient_set(gradient.gradient_, x1, y1, x2, y2).error()
}

// Gets the linear gradient bounds.
//
// The bounds of the linear gradient are defined as a surface constrained by two parallel lines crossing the given points (x1, y1) and (x2, y2), respectively. Both lines are perpendicular to the line linking (x1, y1) and (x2, y2).
//
//   - grad : The Tvg_Gradient object of which to get the bounds.
//   - x1 : The horizontal coordinate of the first point used to determine the gradient bounds.
//   - y1 : The vertical coordinate of the first point used to determine the gradient bounds.
//   - x2 : The horizontal coordinate of the second point used to determine the gradient bounds.
//   - y2 : The vertical coordinate of the second point used to determine the gradient bounds.
func (gradient LinearGradient) Get() (float32, float32, float32, float32, error) {
	var x1, y1 float32
	var x2, y2 float32
	result := tvg_linear_gradient_get(gradient.gradient_, &x1, &y1, &x2, &y2)
	return x1, y1, x2, y2, result.error()
}

// Sets the radial gradient attributes.
//
// The radial gradient is defined by the end circle with a center (cx, cy) and a radius r and the start circle with a center/focal point (fx, fy) and a radius fr. The gradient will be rendered such that the gradient stop at an offset of 100% aligns with the edge of the end circle and the stop at an offset of 0% aligns with the edge of the start circle.
//
//   - grad : The Tvg_Gradient object of which bounds are to be set.
//   - cx : The horizontal coordinate of the center of the end circle.
//   - cy : The vertical coordinate of the center of the end circle.
//   - r : The radius of the end circle.
//   - fx : The horizontal coordinate of the center of the start circle.
//   - fy : The vertical coordinate of the center of the start circle.
//   - fr : The radius of the start circle.
//
// Returns [RESULT_INVALID_ARGUMENT] : The radius r or fr value is negative.
//
// Note: In case the radius r is zero, an object is filled with a single color using the last color specified in the specified in the tvg_gradient_set_color_stops().
//
// Note: In case the focal point (fx and fy) lies outside the end circle, it is projected onto the edge of the end circle.
//
// Note: If the start circle doesn't fully fit inside the end circle (after possible repositioning), the fr is reduced accordingly.
//
// Note: By manipulating the position and size of the focal point, a wide range of visual effects can be achieved, such as directing       the gradient focus towards a specific edge or enhancing the depth and complexity of shading patterns.       If a focal effect is not desired, simply align the focal point (fx and fy) with the center of the end circle (cx and cy)       and set the radius (fr) to zero. This will result in a uniform gradient without any focal variations.
//
// see [GradientSetColorStops]
func (gradient RadialGradient) Set(cx float32, cy float32, r float32, fx float32, fy float32, fr float32) error {
	return tvg_radial_gradient_set(gradient.gradient_, cx, cy, r, fx, fy, fr).error()
}

// The function gets radial gradient attributes.
//
//   - grad : The Tvg_Gradient object of which to get the gradient attributes.
//   - cx : The horizontal coordinate of the center of the end circle.
//   - cy : The vertical coordinate of the center of the end circle.
//   - r : The radius of the end circle.
//   - fx : The horizontal coordinate of the center of the start circle.
//   - fy : The vertical coordinate of the center of the start circle.
//   - fr : The radius of the start circle.
//
// see [RadialGradient.Set]
func (gradient RadialGradient) Get() (float32, float32, float32, float32, float32, float32, error) {
	var cx, cy, r float32
	var fx, fy, fr float32
	result := tvg_radial_gradient_get(gradient.gradient_, &cx, &cy, &r, &fx, &fy, &fr)
	return cx, cy, r, fx, fy, fr, result.error()
}

// Sets the conic gradient attributes.
//
// The gradient sweeps a full turn clockwise around the center (cx, cy), beginning at angle. Color stop offsets from 0% to 100% map clockwise over one full turn from angle.
//
//   - grad : The Tvg_Gradient object of which attributes are to be set.
//   - cx : The horizontal coordinate of the center of the gradient.
//   - cy : The vertical coordinate of the center of the gradient.
//   - angle : The initial angle in degrees, measured clockwise from the 3 o'clock direction (positive x-axis).
//
// Note: FillSpread options are ignored because the gradient always sweeps a full turn.
//
// Note: Experimental API
//
// see [ConicGradient.Get]
func (gradient ConicGradient) Set(cx float32, cy float32, angle float32) error {
	return tvg_conic_gradient_set(gradient.gradient_, cx, cy, angle).error()
}

// Gets the conic gradient attributes.
//
//   - grad : The Tvg_Gradient object of which to get the gradient attributes.
//   - cx : The horizontal coordinate of the center of the gradient.
//   - cy : The vertical coordinate of the center of the gradient.
//   - angle : The initial angle in degrees.
//
// Note: Experimental API
//
// see [ConicGradient.Set]
func (gradient ConicGradient) Get() (float32, float32, float32, error) {
	var cx, cy, angle float32
	result := tvg_conic_gradient_get(gradient.gradient_, &cx, &cy, &angle)
	return cx, cy, angle, result.error()
}

// Sets the parameters of the colors of the gradient and their position.
//
//   - grad : The Tvg_Gradient object of which the color information is to be set.
//   - color_stop : An array of Tvg_Color_Stop data structure.
//   - cnt : The size of the color_stop array equal to the colors number used in the gradient.
func (gradient gradientCommon) SetColorStops(colorStops []ColorStop) error {
	if len(colorStops) == 0 {
		return tvg_gradient_set_color_stops(gradient.gradient_, nil, uint32(len(colorStops))).error()
	}
	return tvg_gradient_set_color_stops(gradient.gradient_, &colorStops[0], uint32(len(colorStops))).error()
}

// Gets the parameters of the colors of the gradient, their position and number
//
// The function does not allocate any memory.
//
//   - grad : The Tvg_Gradient object of which to get the color information.
//   - color_stop : An array of Tvg_Color_Stop data structure.
//   - cnt : The size of the color_stop array equal to the colors number used in the gradient.
func (gradient gradientCommon) GetColorStops() ([]ColorStop, error) {
	var cColorStops *ColorStop
	var cCount uint32
	result := tvg_gradient_get_color_stops(gradient.gradient_, &cColorStops, &cCount)
	colorStops := make([]ColorStop, cCount)
	copy(colorStops, slice(cColorStops, cCount))
	return colorStops, result.error()
}

// Sets the Tvg_Stroke_Fill value, which specifies how to fill the area outside the gradient bounds.
//
//   - grad : The Tvg_Gradient object.
//   - spread : The FillSpread value.
func (gradient gradientCommon) SetSpread(spread StrokeFill) error {
	return tvg_gradient_set_spread(gradient.gradient_, spread).error()
}

// Gets the FillSpread value of the gradient object.
//
//   - grad : The Tvg_Gradient object.
//   - spread : The FillSpread value.
func (gradient gradientCommon) GetSpread() (StrokeFill, error) {
	var spread StrokeFill
	result := tvg_gradient_get_spread(gradient.gradient_, &spread)
	return spread, result.error()
}

// Sets the matrix of the affine transformation for the gradient object.
//
// The augmented matrix of the transformation is expected to be given.
//
//   - grad : The Tvg_Gradient object to be transformed.
//   - m : The 3x3 augmented matrix.
func (gradient gradientCommon) SetTransform(matrix Matrix) error {
	return tvg_gradient_set_transform(gradient.gradient_, &matrix).error()
}

// Gets the matrix of the affine transformation of the gradient object.
//
// In case no transformation was applied, the identity matrix is set.
//
//   - grad : The Tvg_Gradient object of which to get the transformation matrix.
//   - m : The 3x3 augmented matrix.
func (gradient gradientCommon) GetTransform() (Matrix, error) {
	var matrix Matrix
	result := tvg_gradient_get_transform(gradient.gradient_, &matrix)
	return matrix, result.error()
}

// Gets the unique value of the gradient instance indicating the instance type.
//
//   - grad : The Tvg_Gradient object of which to get the type value.
//   - type : The unique type of the gradient instance type.
//
// Since: 1.0
func (gradient gradientCommon) GetType() (Type, error) {
	var typ Type
	result := tvg_gradient_get_type(gradient.gradient_, &typ)
	return typ, result.error()
}

// Duplicates the given Tvg_Gradient object.
//
// Creates a new object and sets its all properties as in the original object.
//
//   - grad : The Tvg_Gradient object to be copied.
//
// Return: A copied Tvg_Gradient object if succeed, nullptr otherwise.
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

// Deletes the given gradient object.
//
//   - grad : The gradient object to be deleted.
func (gradient gradientCommon) Del() error {
	return tvg_gradient_del(gradient.gradient_).error()
}

// Gets the unique value of the gradient instance indicating the instance type.
//
//   - grad : The Tvg_Gradient object of which to get the type value.
//   - type : The unique type of the gradient instance type.
//
// Since: 1.0
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

	case TYPE_CONIC_GRAD:
		return ConicGradient{gradientCommon: gradientCommon{gradient_: gradient}}, true

	default:
		return nil, false
	}
}
