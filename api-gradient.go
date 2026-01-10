package thorvg

import "unsafe"

/*
Gradient is a structure representing a gradient fill of a Paint object.
*/
type Gradient struct {
	gradient uintptr
}

/**
 * RadialGradientNew Creates a new radial gradient object.
 *
 * @return A new radial gradient object.
 */
func RadialGradientNew() Gradient {
	return Gradient{
		gradient: tvg_radial_gradient_new(),
	}
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
func (gradient Gradient) RadialSet(cx float32, cy float32, r float32, fx float32, fy float32, fr float32) error {
	return tvg_radial_gradient_set(gradient.gradient, cx, cy, r, fx, fy, fr).error()
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
func (gradient Gradient) RadialGet() (float32, float32, float32, float32, float32, float32, error) {
	var cx, cy, r float32
	var fx, fy, fr float32
	result := tvg_radial_gradient_get(gradient.gradient, &cx, &cy, &r, &fx, &fy, &fr)
	return cx, cy, r, fx, fy, fr, result.error()
}

/*
SetColorStops Sets the parameters of the colors of the gradient and their position.

	@param[in] grad The Tvg_Gradient object of which the color information is to be set.
	@param[in] color_stop An array of Tvg_Color_Stop data structure.
	@param[in] cnt The size of the @p color_stop array equal to the colors number used in the gradient.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Gradient pointer.
*/
func (gradient Gradient) SetColorStops(colorStops []ColorStop) error {
	if len(colorStops) == 0 {
		return tvg_gradient_set_color_stops(gradient.gradient, nil, uint32(len(colorStops))).error()
	}
	return tvg_gradient_set_color_stops(gradient.gradient, &colorStops[0], uint32(len(colorStops))).error()
}

/*
GetColorStops gets the parameters of the colors of the gradient, their position and number

The function does not allocate any memory.

	@param[in] grad The Tvg_Gradient object of which to get the color information.
	@param[out] color_stop An array of Tvg_Color_Stop data structure.
	@param[out] cnt The size of the @p color_stop array equal to the colors number used in the gradient.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
*/
func (gradient Gradient) GetColorStops() ([]ColorStop, error) {
	var cColorStops *ColorStop
	var cCount uint32
	result := tvg_gradient_get_color_stops(gradient.gradient, &cColorStops, &cCount)
	colorStops := make([]ColorStop, cCount)
	copy(colorStops, unsafe.Slice(cColorStops, cCount))
	return colorStops, result.error()
}
