package thorvg

// #include "thorvg_capi.h"
import "C"

import "unsafe"

/*
Shape represents two-dimensional figures and their properties.

A shape has three major properties: shape outline, stroking, filling.
The outline in the Shape is retained as the path.
Path can be composed by accumulating primitive commands such as moveTo(), lineTo(), cubicTo(),
or complete shape interfaces such as appendRect(), appendCircle(), etc.
Path can consists of sub-paths.
One sub-path is determined by a close command.

The stroke of Shape is an optional property in case the Shape needs to be represented
with/without the outline borders.
It's efficient since the shape path and the stroking path can be shared with each other.
It's also convenient when controlling both in one context.
*/
type Shape struct {
	paintCommon
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
		paintCommon: paintCommon{
			paint_: C.tvg_shape_new(),
		},
	}
}

/*
Reset resets the shape path properties.

The color, the fill and the stroke properties are retained.

	@param[in] paint A Tvg_Paint pointer to the shape object.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note The memory, where the path data is stored, is not deallocated at this stage for caching effect.
*/
func (shape Shape) Reset() error {
	result := C.tvg_shape_reset(shape.paint_)
	return resultError(result)
}

/*
MoveTo sets the initial point of the sub-path.

The value of the current point is set to the given point.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] x The horizontal coordinate of the initial point of the sub-path.
	@param[in] y The vertical coordinate of the initial point of the sub-path.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) MoveTo(x float32, y float32) error {
	result := C.tvg_shape_move_to(shape.paint_, C.float(x), C.float(y))
	return resultError(result)
}

/*
LineTo adds a new point to the sub-path, which results in drawing a line from the current point to the given end-point.

The value of the current point is set to the given end-point.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] x The horizontal coordinate of the end-point of the line.
	@param[in] y The vertical coordinate of the end-point of the line.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note In case this is the first command in the path, it corresponds to the tvg_shape_move_to() call.
*/
func (shape Shape) LineTo(x float32, y float32) error {
	result := C.tvg_shape_line_to(shape.paint_, C.float(x), C.float(y))
	return resultError(result)
}

/*
CubicTo Adds new points to the sub-path, which results in drawing a cubic Bezier curve.

The Bezier curve starts at the current point and ends at the given end-point (@p x, @p y). Two control points (@p cx1, @p cy1) and (@p cx2, @p cy2) are used to determine the shape of the curve.
The value of the current point is set to the given end-point.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] cx1 The horizontal coordinate of the 1st control point.
	@param[in] cy1 The vertical coordinate of the 1st control point.
	@param[in] cx2 The horizontal coordinate of the 2nd control point.
	@param[in] cy2 The vertical coordinate of the 2nd control point.
	@param[in] x The horizontal coordinate of the endpoint of the curve.
	@param[in] y The vertical coordinate of the endpoint of the curve.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note In case this is the first command in the path, no data from the path are rendered.
*/
func (shape Shape) CubicTo(cx1 float32, cy1 float32, cx2 float32, cy2 float32, x float32, y float32) error {
	result := C.tvg_shape_cubic_to(shape.paint_, C.float(cx1), C.float(cy1), C.float(cx2), C.float(cy2), C.float(x), C.float(y))
	return resultError(result)
}

/*
Close closes the current sub-path by drawing a line from the current point to the initial point of the sub-path.

The value of the current point is set to the initial point of the closed sub-path.

	@param[in] paint A Tvg_Paint pointer to the shape object.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note In case the sub-path does not contain any points, this function has no effect.
*/
func (shape Shape) Close() error {
	result := C.tvg_shape_close(shape.paint_)
	return resultError(result)
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
	result := C.tvg_shape_append_rect(shape.paint_,
		C.float(x), C.float(y), C.float(w), C.float(h), C.float(rx), C.float(ry), C.bool(cw))
	return resultError(result)
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
	result := C.tvg_shape_append_circle(shape.paint_, C.float(cx), C.float(cy), C.float(rx), C.float(ry), C.bool(cw))
	return resultError(result)
}

/*
AppendPath appends a given sub-path to the path.

The current point value is set to the last point from the sub-path.
For each command from the @p cmds array, an appropriate number of points in @p pts array should be specified.
If the number of points in the @p pts array is different than the number required by the @p cmds array, the shape with this sub-path will not be displayed on the screen.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] cmds The array of the commands in the sub-path.
	@param[in] cmdCnt The length of the @p cmds array.
	@param[in] pts The array of the two-dimensional points.
	@param[in] ptsCnt The length of the @p pts array.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument or @p cmdCnt or @p ptsCnt equal to zero.
*/
func (shape Shape) AppendPath(cmds []PathCommand, pts []Point) error {
	result := C.tvg_shape_append_path(shape.paint_,
		(*C.Tvg_Path_Command)(unsafe.Pointer(&cmds[0])), C.uint32_t(len(cmds)),
		(*C.Tvg_Point)(unsafe.Pointer(&pts[0])), C.uint32_t(len(pts)),
	)
	return resultError(result)
}

/*
GetPath retrieves the current path data of the shape.

This function provides access to the shape's path data, including the commands
and points that define the path.

	@param[out] cmds Pointer to the array of commands representing the path.
								Can be @c nullptr if this information is not needed.
	@param[out] cmdsCnt Pointer to the variable that receives the number of commands in the @p cmds array.
										Can be @c nullptr if this information is not needed.
	@param[out] pts Pointer to the array of two-dimensional points that define the path.
								Can be @c nullptr if this information is not needed.
	@param[out] ptsCnt Pointer to the variable that receives the number of points in the @p pts array.
									Can be @c nullptr if this information is not needed.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note If any of the arguments are @c nullptr, that value will be ignored.
*/
func (shape Shape) GetPath() ([]PathCommand, []Point, error) {
	var cmds *C.Tvg_Path_Command
	var cmdsCnt C.uint32_t
	var pts *C.Tvg_Point
	var ptsCnt C.uint32_t
	result := C.tvg_shape_get_path(shape.paint_, &cmds, &cmdsCnt, &pts, &ptsCnt)
	return slice((*PathCommand)(unsafe.Pointer(cmds)), cmdsCnt),
		slice((*Point)(unsafe.Pointer(pts)), ptsCnt),
		resultError(result)
}

/*
SetStrokeWidth sets the stroke width for the path.

This function defines the thickness of the stroke applied to all figures
in the path object. A stroke is the outline drawn along the edges of the
path's geometry.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] width The width of the stroke in pixels. Must be positive value. (The default is 0)

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note A value of @p width 0 disables the stroke.

	@see tvg_shape_set_stroke_color()
*/
func (shape Shape) SetStrokeWidth(width float32) error {
	result := C.tvg_shape_set_stroke_width(shape.paint_, C.float(width))
	return resultError(result)
}

/*
GetStrokeWidth gets the shape's stroke width.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] width The stroke width.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetStrokeWidth() (float32, error) {
	var width C.float
	result := C.tvg_shape_get_stroke_width(shape.paint_, &width)
	return float32(width), resultError(result)
}

/*
SetStrokeColor sets the shape's stroke color.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] r The red color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] g The green color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] b The blue color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] a The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note If the stroke width is 0 (default), the stroke will not be visible regardless of the color.
	@note Either a solid color or a gradient fill is applied, depending on what was set as last.

	@see tvg_shape_set_stroke_width()
	@see tvg_shape_set_stroke_gradient()
*/
func (shape Shape) SetStrokeColor(r uint8, g uint8, b uint8, a uint8) error {
	result := C.tvg_shape_set_stroke_color(shape.paint_, C.uint8_t(r), C.uint8_t(g), C.uint8_t(b), C.uint8_t(a))
	return resultError(result)
}

/*
GetStrokeColor gets the shape's stroke color.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] r The red color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] g The green color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] b The blue color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] a The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION No stroke was set.
*/
func (shape Shape) GetStrokeColor() (uint8, uint8, uint8, uint8, error) {
	var r C.uint8_t
	var g C.uint8_t
	var b C.uint8_t
	var a C.uint8_t
	result := C.tvg_shape_get_stroke_color(shape.paint_, &r, &g, &b, &a)
	return uint8(r), uint8(g), uint8(b), uint8(a), resultError(result)
}

/*
SetStrokeGradient sets the gradient fill of the stroke for all of the figures from the path.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] grad The gradient fill.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_MEMORY_CORRUPTION An invalid Tvg_Gradient pointer or an error with accessing it.

	@note Either a solid color or a gradient fill is applied, depending on what was set as last.

	@see tvg_shape_set_stroke_color()
*/
func (shape Shape) SetStrokeGradient(grad Gradient) error {
	result := C.tvg_shape_set_stroke_gradient(shape.paint_, grad.gradient())
	return resultError(result)
}

/*
GetStrokeGradient gets the gradient fill of the shape's stroke.

The function does not allocate any memory.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] grad The gradient fill.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetStrokeGradient() (Gradient, error) {
	var grad C.Tvg_Gradient
	result := C.tvg_shape_get_stroke_gradient(shape.paint_, &grad)
	strokeGradient, _ := newGradient(grad)
	return strokeGradient, resultError(result)
}

/*
SetStrokeDash sets the shape's stroke dash pattern.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] dashPattern An array of alternating dash and gap lengths.
	@param[in] cnt The size of the @p dashPattern array.
	@param[in] offset The shift of the starting point within the repeating dash pattern, from which the pattern begins to be applied.

	@retval TVG_RESULT_INVALID_ARGUMENT In case @p dashPattern is @c nullptr and @p cnt > 0 or @p dashPattern is not @c nullptr and @p cnt is zero.

	@note To reset the stroke dash pattern, pass @c nullptr to @p dashPattern and zero to @p cnt.
	@note Values of @p dashPattern less than zero are treated as zero.
	@note If all values in the @p dashPattern are equal to or less than 0, the dash is ignored.
	@note If the @p dashPattern contains an odd number of elements, the sequence is repeated in the same

order to form an even-length pattern, preserving the alternation of dashes and gaps.

	@since 1.0
*/
func (shape Shape) SetStrokeDash(pattern []float32, offset float32) error {
	result := C.tvg_shape_set_stroke_dash(shape.paint_,
		(*C.float)(unsafe.Pointer(&pattern[0])), C.uint32_t(len(pattern)),
		C.float(offset),
	)
	return resultError(result)
}

/*
GetStrokeDash gets the dash pattern of the stroke.

The function does not allocate any memory.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] dashPattern The array of consecutive pair values of the dash length and the gap length.
	@param[out] cnt The size of the @p dashPattern array.
	@param[out] offset The shift of the starting point within the repeating dash pattern.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
	@since 1.0
*/
func (shape Shape) GetStrokeDash() ([]float32, float32, error) {
	var dashPattern *C.float
	var cnt C.uint32_t
	var offset C.float
	result := C.tvg_shape_get_stroke_dash(shape.paint_, &dashPattern, &cnt, &offset)
	return slice((*float32)(unsafe.Pointer(dashPattern)), cnt), float32(offset),
		resultError(result)
}

/*
SetStrokeCap sets the cap style used for stroking the path.

The cap style specifies the shape to be used at the end of the open stroked sub-paths.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] cap The cap style value. The default value is @c TVG_STROKE_CAP_SQUARE.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) SetStrokeCap(strokeCap StrokeCap) error {
	result := C.tvg_shape_set_stroke_cap(shape.paint_, C.Tvg_Stroke_Cap(strokeCap))
	return resultError(result)
}

/*
GetStrokeCap gets the stroke cap style used for stroking the path.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] cap The cap style value.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetStrokeCap() (StrokeCap, error) {
	var strokeCap C.Tvg_Stroke_Cap
	result := C.tvg_shape_get_stroke_cap(shape.paint_, &strokeCap)
	return *(*StrokeCap)(unsafe.Pointer(&strokeCap)), resultError(result)
}

/*
SetStrokeJoin sets the join style for stroked path segments.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] join The join style value. The default value is @c TVG_STROKE_JOIN_BEVEL.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) SetStrokeJoin(join StrokeJoin) error {
	result := C.tvg_shape_set_stroke_join(shape.paint_, C.Tvg_Stroke_Join(join))
	return resultError(result)
}

/*
GetStrokeJoin gets the stroke join method

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] join The join style value.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetStrokeJoin() (StrokeJoin, error) {
	var join C.Tvg_Stroke_Join
	result := C.tvg_shape_get_stroke_join(shape.paint_, &join)
	return *(*StrokeJoin)(unsafe.Pointer(&join)), resultError(result)
}

/*
SetStrokeMiterlimit sets the stroke miterlimit.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] miterlimit The miterlimit imposes a limit on the extent of the stroke join when the @c TVG_STROKE_JOIN_MITER join style is set. The default value is 4.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer or Unsupported @p miterlimit values (less than zero).

	@since 0.11
*/
func (shape Shape) SetStrokeMiterlimit(miterlimit float32) error {
	result := C.tvg_shape_set_stroke_miterlimit(shape.paint_, C.float(miterlimit))
	return resultError(result)
}

/*
GetStrokeMiterlimit gets the stroke miterlimit.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] miterlimit The stroke miterlimit.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.

	@since 0.11
*/
func (shape Shape) GetStrokeMiterlimit() (float32, error) {
	var miterlimit C.float
	result := C.tvg_shape_get_stroke_miterlimit(shape.paint_, &miterlimit)
	return float32(miterlimit), resultError(result)
}

/*
SetTrimpath sets the trim of the shape along the defined path segment, allowing control over which part of the shape is visible.

If the values of the arguments @p begin and @p end exceed the 0-1 range, they are wrapped around in a manner similar to angle wrapping, effectively treating the range as circular.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] begin Specifies the start of the segment to display along the path.
	@param[in] end Specifies the end of the segment to display along the path.
	@param[in] simultaneous Determines how to trim multiple paths within a single shape. If set to @c true (default), trimming is applied simultaneously to all paths;
												Otherwise, all paths are treated as a single entity with a combined length equal to the sum of their individual lengths and are trimmed as such.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@since 1.0
*/
func (shape Shape) SetTrimpath(begin float32, end float32, simultaneous bool) error {
	result := C.tvg_shape_set_trimpath(shape.paint_, C.float(begin), C.float(end), C.bool(simultaneous))
	return resultError(result)
}

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
	result := C.tvg_shape_set_fill_color(shape.paint_, C.uint8_t(r), C.uint8_t(g), C.uint8_t(b), C.uint8_t(a))
	return resultError(result)
}

/*
GetFillColor gets the shape's solid color.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] r The red color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] g The green color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] b The blue color channel value in the range [0 ~ 255]. The default value is 0.
	@param[out] a The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque. The default value is 0.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) GetFillColor() (uint8, uint8, uint8, uint8, error) {
	var r C.uint8_t
	var g C.uint8_t
	var b C.uint8_t
	var a C.uint8_t
	result := C.tvg_shape_get_fill_color(shape.paint_, &r, &g, &b, &a)
	return uint8(r), uint8(g), uint8(b), uint8(a), resultError(result)
}

/*
SetFillRule sets the fill rule for the shape.

Specifies how the interior of the shape is determined when its path intersects itself.
The default fill rule is @c TVG_FILL_RULE_NON_ZERO.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] rule The fill rule to apply to the shape.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (shape Shape) SetFillRule(rule FillRule) error {
	result := C.tvg_shape_set_fill_rule(shape.paint_, C.Tvg_Fill_Rule(rule))
	return resultError(result)
}

/*
GetFillRule retrieves the current fill rule used by the shape.

This function returns the fill rule, which determines how the interior
regions of the shape are calculated when it overlaps itself.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] rule The current Tvg_Fill_Rule value of the shape.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetFillRule() (FillRule, error) {
	var rule C.Tvg_Fill_Rule
	result := C.tvg_shape_get_fill_rule(shape.paint_, &rule)
	return FillRule(rule), resultError(result)
}

/*
SetPaintOrder sets the rendering order of the stroke and the fill.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[in] strokeFirst If @c true the stroke is rendered before the fill, otherwise the stroke is rendered as the second one (the default option).

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@since 0.10
*/
func (shape Shape) SetPaintOrder(strokeFirst bool) error {
	result := C.tvg_shape_set_paint_order(shape.paint_, C.bool(strokeFirst))
	return resultError(result)
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
	result := C.tvg_shape_set_gradient(shape.paint_, gradient.gradient())
	return resultError(result)
}

/*
GetGradient gets the gradient fill of the shape.

The function does not allocate any data.

	@param[in] paint A Tvg_Paint pointer to the shape object.
	@param[out] grad The gradient fill.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer passed as an argument.
*/
func (shape Shape) GetGradient() (Gradient, bool, error) {
	var gradient C.Tvg_Gradient
	result := C.tvg_shape_get_gradient(shape.paint_, &gradient)
	if result != C.TVG_RESULT_SUCCESS || gradient == nil {
		return nil, false, resultError(result)
	}

	grad, ok := newGradient(gradient)
	return grad, ok, resultError(result)
}

/*
Duplicate duplicates a Shape.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (shape Shape) Duplicate() Shape {
	return Shape{
		paintCommon: shape.duplicate(),
	}
}
