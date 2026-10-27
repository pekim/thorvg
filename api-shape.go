package thorvg

type Shape struct {
	paintCommon
}

// Creates a new Shape object.
//
// This function allocates and returns a new Shape instance. To properly destroy the Shape object, use [Paint.Rel]
//
// Return: The newly created Shape object.
//
// see [Paint.Rel]
func ShapeNew() Shape {
	return Shape{
		paintCommon: paintCommon{
			paint_: tvg_shape_new(),
		},
	}
}

// Resets the shape path properties.
//
// The color, the fill and the stroke properties are retained.
//
//   - paint : The shape object.
//
// Note: The memory, where the path data is stored, is not deallocated at this stage for caching effect.
func (shape Shape) Reset() error {
	return tvg_shape_reset(shape.paint_).error()
}

// Sets the initial point of the sub-path.
//
// The value of the current point is set to the given point.
//
//   - paint : The shape object.
//   - x : The horizontal coordinate of the initial point of the sub-path.
//   - y : The vertical coordinate of the initial point of the sub-path.
func (shape Shape) MoveTo(x float32, y float32) error {
	return tvg_shape_move_to(shape.paint_, x, y).error()
}

// Adds a new point to the sub-path, which results in drawing a line from the current point to the given end-point.
//
// The value of the current point is set to the given end-point.
//
//   - paint : The shape object.
//   - x : The horizontal coordinate of the end-point of the line.
//   - y : The vertical coordinate of the end-point of the line.
//
// Note: In case this is the first command in the path, it corresponds to the tvg_shape_move_to() call.
func (shape Shape) LineTo(x float32, y float32) error {
	return tvg_shape_line_to(shape.paint_, x, y).error()
}

// Adds new points to the sub-path, which results in drawing a cubic Bezier curve.
//
// The Bezier curve starts at the current point and ends at the given end-point (x, y). Two control points (cx1, cy1) and (cx2, cy2) are used to determine the shape of the curve. The value of the current point is set to the given end-point.
//
//   - paint : The shape object.
//   - cx1 : The horizontal coordinate of the 1st control point.
//   - cy1 : The vertical coordinate of the 1st control point.
//   - cx2 : The horizontal coordinate of the 2nd control point.
//   - cy2 : The vertical coordinate of the 2nd control point.
//   - x : The horizontal coordinate of the endpoint of the curve.
//   - y : The vertical coordinate of the endpoint of the curve.
//
// Note: In case this is the first command in the path, no data from the path are rendered.
func (shape Shape) CubicTo(cx1 float32, cy1 float32, cx2 float32, cy2 float32, x float32, y float32) error {
	return tvg_shape_cubic_to(shape.paint_, cx1, cy1, cx2, cy2, x, y).error()
}

// Closes the current sub-path by drawing a line from the current point to the initial point of the sub-path.
//
// The value of the current point is set to the initial point of the closed sub-path.
//
//   - paint : The shape object.
//
// Note: In case the sub-path does not contain any points, this function has no effect.
func (shape Shape) Close() error {
	return tvg_shape_close(shape.paint_).error()
}

// Appends a rectangle to the path.
//
// The rectangle with rounded corners can be achieved by setting non-zero values to rx and ry arguments. The rx and ry values specify the radii of the ellipse defining the rounding of the corners.
// The position of the rectangle is specified by the coordinates of its upper-left corner -  x and y arguments.
// The rectangle is treated as a new sub-path - it is not connected with the previous sub-path.
// The value of the current point is set to (x + w, y + ry) - in case ry is greater than h/2 the current point is set to (x + w, y + h/2).
//
//   - paint : The shape object.
//   - x : The horizontal coordinate of the upper-left corner of the rectangle.
//   - y : The vertical coordinate of the upper-left corner of the rectangle.
//   - w : The width of the rectangle.
//   - h : The height of the rectangle.
//   - rx : The x-axis radius of the ellipse defining the rounded corners of the rectangle.
//   - ry : The y-axis radius of the ellipse defining the rounded corners of the rectangle.
//   - cw : Specifies the path direction: true for clockwise, false for counterclockwise.
//
// Note: For rx and ry greater than or equal to the half of w and the half of h, respectively, the shape become an ellipse.
func (shape Shape) AppendRect(x float32, y float32, w float32, h float32, rx float32, ry float32, cw bool) error {
	return tvg_shape_append_rect(shape.paint_, x, y, w, h, rx, ry, cw).error()
}

// Appends an ellipse to the path.
//
// The position of the ellipse is specified by the coordinates of its center - cx and cy arguments.
// The ellipse is treated as a new sub-path - it is not connected with the previous sub-path.
// The value of the current point is set to (cx, cy - ry).
//
//   - paint : The shape object.
//   - cx : The horizontal coordinate of the center of the ellipse.
//   - cy : The vertical coordinate of the center of the ellipse.
//   - rx : The x-axis radius of the ellipse.
//   - ry : The y-axis radius of the ellipse.
//   - cw : Specifies the path direction: true for clockwise, false for counterclockwise.
func (shape Shape) AppendCircle(cx float32, cy float32, rx float32, ry float32, cw bool) error {
	return tvg_shape_append_circle(shape.paint_, cx, cy, rx, ry, cw).error()
}

// Appends a given sub-path to the path.
//
// The current point value is set to the last point from the sub-path. For each command from the cmds array, an appropriate number of points in pts array should be specified. If the number of points in the pts array is different than the number required by the cmds array, the shape with this sub-path will not be displayed on the screen.
//
//   - paint : The shape object.
//   - cmds : The array of the commands in the sub-path.
//   - cmdCnt : The length of the cmds array.
//   - pts : The array of the two-dimensional points.
//   - ptsCnt : The length of the pts array.
//
// Returns [RESULT_INVALID_ARGUMENT] : A nullptr passed as the argument or cmdCnt or ptsCnt equal to zero.
func (shape Shape) AppendPath(cmds []PathCommand, pts []Point) error {
	return tvg_shape_append_path(shape.paint_, &cmds[0], uint32(len(cmds)), &pts[0], uint32(len(pts))).error()
}

// Retrieves the current path data of the shape.
//
// This function provides access to the shape's path data, including the commands and points that define the path.
//
//   - cmds : Pointer to the array of commands representing the path.                  Can be nullptr if this information is not needed.
//   - cmdsCnt : Pointer to the variable that receives the number of commands in the cmds array.                     Can be nullptr if this information is not needed.
//   - pts : Pointer to the array of two-dimensional points that define the path.                 Can be nullptr if this information is not needed.
//   - ptsCnt : Pointer to the variable that receives the number of points in the pts array.                    Can be nullptr if this information is not needed.
//
// Note: If any of the arguments are nullptr, that value will be ignored.
func (shape Shape) GetPath() ([]PathCommand, []Point, error) {
	var cmds *PathCommand
	var cmdsCnt uint32
	var pts *Point
	var ptsCnt uint32
	result := tvg_shape_get_path(shape.paint_, &cmds, &cmdsCnt, &pts, &ptsCnt)
	return slice(cmds, cmdsCnt), slice(pts, ptsCnt), result.error()
}

// Sets the stroke width for the path.
//
// This function defines the thickness of the stroke applied to all figures in the path object. A stroke is the outline drawn along the edges of the path's geometry.
//
//   - paint : The shape object.
//   - width : The width of the stroke in pixels. Must be positive value. (The initial value is 0)
//
// Note: A value of width 0 disables the stroke.
//
// see [Shape.SetStrokeColor]
func (shape Shape) SetStrokeWidth(width float32) error {
	return tvg_shape_set_stroke_width(shape.paint_, width).error()
}

// Gets the shape's stroke width.
//
//   - paint : The shape object.
//   - width : The stroke width.
func (shape Shape) GetStrokeWidth() (float32, error) {
	var width float32
	result := tvg_shape_get_stroke_width(shape.paint_, &width)
	return width, result.error()
}

// Sets the shape's stroke color.
//
//   - paint : The shape object.
//   - r : The red color channel value in the range [0 ~ 255]. The initial value is 0.
//   - g : The green color channel value in the range [0 ~ 255]. The initial value is 0.
//   - b : The blue color channel value in the range [0 ~ 255]. The initial value is 0.
//   - a : The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque. The initial value is 0.
//
// Note: If the stroke width is 0 (initial value), the stroke will not be visible regardless of the color.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Shape.SetStrokeWidth]
//
// see [Shape.SetStrokeGradient]
func (shape Shape) SetStrokeColor(r uint, g uint, b uint, a uint) error {
	return tvg_shape_set_stroke_color(shape.paint_, uint8(r), uint8(g), uint8(b), uint8(a)).error()
}

// Gets the shape's stroke color.
//
//   - paint : The shape object.
//   - r : The red color channel value in the range [0 ~ 255].
//   - g : The green color channel value in the range [0 ~ 255].
//   - b : The blue color channel value in the range [0 ~ 255].
//   - a : The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : No stroke was set.
func (shape Shape) GetStrokeColor() (int, int, int, int, error) {
	var r uint8
	var g uint8
	var b uint8
	var a uint8
	result := tvg_shape_get_stroke_color(shape.paint_, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a), result.error()
}

// Sets the gradient fill of the stroke for all of the figures from the path.
//
//   - paint : The shape object.
//   - grad : The gradient fill.
//
// Returns [RESULT_MEMORY_CORRUPTION] : An invalid gradient object or an error with accessing it.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Shape.SetStrokeColor]
func (shape Shape) SetStrokeGradient(grad Gradient) error {
	return tvg_shape_set_stroke_gradient(shape.paint_, grad.gradient()).error()
}

// Gets the gradient fill of the shape's stroke.
//
// The function does not allocate any memory.
//
//   - paint : The shape object.
//   - grad : The gradient fill.
func (shape Shape) GetStrokeGradient() (Gradient, error) {
	var grad uintptr
	result := tvg_shape_get_stroke_gradient(shape.paint_, &grad)
	strokeGradient, _ := newGradient(grad)
	return strokeGradient, result.error()
}

// Sets the shape's stroke dash pattern.
//
//   - paint : The shape object.
//   - dashPattern : An array of alternating dash and gap lengths.
//   - cnt : The size of the dashPattern array.
//   - offset : The shift of the starting point within the repeating dash pattern, from which the pattern begins to be applied.
//
// Returns [RESULT_INVALID_ARGUMENT] : In case dashPattern is nullptr and cnt > 0 or dashPattern is not nullptr and cnt is zero.
//
// Note: To reset the stroke dash pattern, pass nullptr to dashPattern and zero to cnt.
//
// Note: Values of dashPattern less than zero are treated as zero.
//
// Note: If all values in the dashPattern are equal to or less than 0, the dash is ignored.
//
// Note: If the dashPattern contains an odd number of elements, the sequence is repeated in the same order to form an even-length pattern, preserving the alternation of dashes and gaps.
// Since: 1.0
func (shape Shape) SetStrokeDash(pattern []float32, offset float32) error {
	return tvg_shape_set_stroke_dash(shape.paint_, &pattern[0], uint32(len(pattern)), offset).error()
}

// Gets the dash pattern of the stroke.
//
// The function does not allocate any memory.
//
//   - paint : The shape object.
//   - dashPattern : The array of consecutive pair values of the dash length and the gap length.
//   - cnt : The size of the dashPattern array.
//   - offset : The shift of the starting point within the repeating dash pattern.
//
// Since: 1.0
func (shape Shape) GetStrokeDash() ([]float32, float32, error) {
	var dashPattern *float32
	var cnt uint32
	var offset float32
	result := tvg_shape_get_stroke_dash(shape.paint_, &dashPattern, &cnt, &offset)
	return slice(dashPattern, cnt), offset, result.error()
}

// Sets the cap style used for stroking the path.
//
// The cap style specifies the shape to be used at the end of the open stroked sub-paths.
//
//   - paint : The shape object.
//   - cap : The cap style value. The initial value is TVG_STROKE_CAP_SQUARE.
func (shape Shape) SetStrokeCap(strokeCap StrokeCap) error {
	return tvg_shape_set_stroke_cap(shape.paint_, strokeCap).error()
}

// Gets the stroke cap style used for stroking the path.
//
//   - paint : The shape object.
//   - cap : The cap style value.
func (shape Shape) GetStrokeCap() (StrokeCap, error) {
	var strokeCap StrokeCap
	result := tvg_shape_get_stroke_cap(shape.paint_, &strokeCap)
	return strokeCap, result.error()
}

// Sets the join style for stroked path segments.
//
//   - paint : The shape object.
//   - join : The join style value. The initial value is TVG_STROKE_JOIN_BEVEL.
func (shape Shape) SetStrokeJoin(join StrokeJoin) error {
	return tvg_shape_set_stroke_join(shape.paint_, join).error()
}

// The function gets the stroke join method
//
//   - paint : The shape object.
//   - join : The join style value.
func (shape Shape) GetStrokeJoin() (StrokeJoin, error) {
	var join StrokeJoin
	result := tvg_shape_get_stroke_join(shape.paint_, &join)
	return join, result.error()
}

// Sets the stroke miterlimit.
//
//   - paint : The shape object.
//   - miterlimit : The miterlimit imposes a limit on the extent of the stroke join when the TVG_STROKE_JOIN_MITER join style is set. The initial value is 4.
//
// Returns [RESULT_INVALID_ARGUMENT] : Unsupported miterlimit values (less than zero).
//
// Since: 0.11
func (shape Shape) SetStrokeMiterlimit(miterlimit float32) error {
	return tvg_shape_set_stroke_miterlimit(shape.paint_, miterlimit).error()
}

// The function gets the stroke miterlimit.
//
//   - paint : The shape object.
//   - miterlimit : The stroke miterlimit.
//
// Since: 0.11
func (shape Shape) GetStrokeMiterlimit() (float32, error) {
	var miterlimit float32
	result := tvg_shape_get_stroke_miterlimit(shape.paint_, &miterlimit)
	return miterlimit, result.error()
}

// Sets the trim of the shape along the defined path segment, allowing control over which part of the shape is visible.
//
// If the values of the arguments begin and end exceed the 0-1 range, they are wrapped around in a manner similar to angle wrapping, effectively treating the range as circular.
//
//   - paint : The shape object.
//   - begin : Specifies the start of the segment to display along the path.
//   - end : Specifies the end of the segment to display along the path.
//   - simultaneous : Determines how to trim multiple paths within a single shape. If set to true (default), trimming is applied simultaneously to all paths;                         Otherwise, all paths are treated as a single entity with a combined length equal to the sum of their individual lengths and are trimmed as such.
//
// Since: 1.0
func (shape Shape) SetTrimpath(begin float32, end float32, simultaneous bool) error {
	return tvg_shape_set_trimpath(shape.paint_, begin, end, simultaneous).error()
}

// Sets the shape's solid color.
//
// The parts of the shape defined as inner are colored.
//
//   - paint : The shape object.
//   - r : The red color channel value in the range [0 ~ 255]. The initial value is 0.
//   - g : The green color channel value in the range [0 ~ 255]. The initial value is 0.
//   - b : The blue color channel value in the range [0 ~ 255]. The initial value is 0.
//   - a : The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque. The initial value is 0.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Shape.SetFillRule]
func (shape Shape) SetFillColor(r uint8, g uint8, b uint8, a uint8) error {
	return tvg_shape_set_fill_color(shape.paint_, r, g, b, a).error()
}

// Gets the shape's solid color.
//
//   - paint : The shape object.
//   - r : The red color channel value in the range [0 ~ 255].
//   - g : The green color channel value in the range [0 ~ 255].
//   - b : The blue color channel value in the range [0 ~ 255].
//   - a : The alpha channel value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
func (shape Shape) GetFillColor() (uint8, uint8, uint8, uint8, error) {
	var r uint8
	var g uint8
	var b uint8
	var a uint8
	result := tvg_shape_get_fill_color(shape.paint_, &r, &g, &b, &a)
	return r, g, b, a, result.error()
}

// Sets the fill rule for the shape.
//
// Specifies how the interior of the shape is determined when its path intersects itself. The initial fill rule is TVG_FILL_RULE_NON_ZERO.
//
//   - paint : The shape object.
//   - rule : The fill rule to apply to the shape.
func (shape Shape) SetFillRule(rule FillRule) error {
	return tvg_shape_set_fill_rule(shape.paint_, rule).error()
}

// Retrieves the current fill rule used by the shape.
//
// This function returns the fill rule, which determines how the interior  regions of the shape are calculated when it overlaps itself.
//
//   - paint : The shape object.
//   - rule : The current Tvg_Fill_Rule value of the shape.
func (shape Shape) GetFillRule() (FillRule, error) {
	var rule FillRule
	result := tvg_shape_get_fill_rule(shape.paint_, &rule)
	return rule, result.error()
}

// Sets the rendering order of the stroke and the fill.
//
//   - paint : The shape object.
//   - strokeFirst : If true the stroke is rendered before the fill, otherwise the stroke is rendered as the second one (the initial option).
//
// Since: 0.10
func (shape Shape) SetPaintOrder(strokeFirst bool) error {
	return tvg_shape_set_paint_order(shape.paint_, strokeFirst).error()
}

// Sets the gradient fill for all of the figures from the path.
//
// The parts of the shape defined as inner are filled.
//
//   - paint : The shape object.
//   - grad : The gradient fill.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Shape.SetFillRule]
func (shape Shape) SetGradient(gradient Gradient) error {
	return tvg_shape_set_gradient(shape.paint_, gradient.gradient()).error()
}

// Gets the gradient fill of the shape.
//
// The function does not allocate any data.
//
//   - paint : The shape object.
//   - grad : The gradient fill.
func (shape Shape) GetGradient() (Gradient, bool, error) {
	var gradient uintptr
	result := tvg_shape_get_gradient(shape.paint_, &gradient)
	if result != RESULT_SUCCESS || gradient == 0 {
		return nil, false, result.error()
	}

	grad, ok := newGradient(gradient)
	return grad, ok, result.error()
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
