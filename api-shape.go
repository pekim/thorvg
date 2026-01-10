package thorvg

type Shape struct {
	paint_ uintptr
}

func (shape Shape) paint() uintptr {
	return shape.paint_
}

/*
ShapeNew creates a new Shape object.

This function allocates and returns a new Shape instance.
To properly destroy the Shape object, use @ref tvg_paint_rel().

@return A pointer to the newly created Shape object.

@see tvg_paint_rel()
*/
func ShapeNew() Shape {
	return Shape{
		paint_: tvg_shape_new(),
	}
}

/*
AppendRect appends a rectangle to the path.

The rectangle with rounded corners can be achieved by setting non-zero values to @p rx and @p ry arguments.
The @p rx and @p ry values specify the radii of the ellipse defining the rounding of the corners.

The position of the rectangle is specified by the coordinates of its upper-left corner -  @p x and @p y arguments.

The rectangle is treated as a new sub-path - it is not connected with the previous sub-path.

The value of the current point is set to (@p x + @p rx, @p y) - in case @p rx is greater
than @p w/2 the current point is set to (@p x + @p w/2, @p y)

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] x The horizontal coordinate of the upper-left corner of the rectangle.
	@param[in] y The vertical coordinate of the upper-left corner of the rectangle.
	@param[in] w The width of the rectangle.
	@param[in] h The height of the rectangle.
	@param[in] rx The x-axis radius of the ellipse defining the rounded corners of the rectangle.
	@param[in] ry The y-axis radius of the ellipse defining the rounded corners of the rectangle.
	@param[in] cw Specifies the path direction: @c true for clockwise, @c false for counterclockwise.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

@note For @p rx and @p ry greater than or equal to the half of @p w and the half of @p h, respectively, the shape become an ellipse.
*/
func (shape Shape) AppendRect(x float32, y float32, w float32, h float32, rx float32, ry float32, cw bool) error {
	return tvg_shape_append_rect(shape.paint_, x, y, w, h, rx, ry, cw).error()
}

/*
AppendCircle Appends an ellipse to the path.

The position of the ellipse is specified by the coordinates of its center - @p cx and @p cy arguments.

The ellipse is treated as a new sub-path - it is not connected with the previous sub-path.

The value of the current point is set to (@p cx, @p cy - @p ry).

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] cx The horizontal coordinate of the center of the ellipse.
	@param[in] cy The vertical coordinate of the center of the ellipse.
	@param[in] rx The x-axis radius of the ellipse.
	@param[in] ry The y-axis radius of the ellipse.
	@param[in] cw Specifies the path direction: @c true for clockwise, @c false for counterclockwise.

@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) AppendCircle(cx float32, cy float32, rx float32, ry float32, cw bool) error {
	return tvg_shape_append_circle(shape.paint_, cx, cy, rx, ry, cw).error()
}

// TVG_API Tvg_Result tvg_shape_append_circle(Tvg_Paint paint, float cx, float cy, float rx, float ry, bool cw);

/*
SetFillColor sets the shape's solid color.

The parts of the shape defined as inner are colored.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] r The red color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] g The green color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] b The blue color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] a The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque. The default value is 0.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

@note Either a solid color or a gradient fill is applied, depending on what was set as last.
@see tvg_shape_set_fill_rule()
*/
func (shape Shape) SetFillColor(r uint8, g uint8, b uint8, a uint8) error {
	return tvg_shape_set_fill_color(shape.paint_, r, g, b, a).error()
}

/*
SetGradient sets the gradient fill for all of the figures from the path.

The parts of the shape defined as inner are filled.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] grad The gradient fill.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_MEMORY_CORRUPTION An invalid Tvg_Gradient pointer.

@note Either a solid color or a gradient fill is applied, depending on what was set as last.
@see tvg_shape_set_fill_rule()
*/
func (shape Shape) SetGradient(gradient Gradient) error {
	return tvg_shape_set_gradient(shape.paint_, gradient.gradient).error()
}

/*
 * @brief Gets the gradient fill of the shape.
 *
 * The function does not allocate any data.
 *
 * @param[in] paint A Tvg_Paint pointer to the shape object.
 * @param[out] grad The gradient fill.
 *
 * @retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
 */
func (shape Shape) GetGradient() (Gradient, error) {
	var gradient uintptr
	result := tvg_shape_get_gradient(shape.paint_, &gradient)
	return Gradient{gradient: gradient}, result.error()
}
