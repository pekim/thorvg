package thorvg

/*
Paint is a structure representing a graphical element.

	@warning The TvgPaint objects cannot be shared between Canvases.
*/
type Paint interface {
	paint() uintptr
}

type paintCommon struct {
	paint_ uintptr
}

func (paint paintCommon) paint() uintptr {
	return paint.paint_
}

// Safely releases a Tv_Paint object.
//
// This is the counterpart to the `new()` API, and releases the given Paint object safely,  handling nullptr and managing ownership properly.
//
//   - paint : A paint object to release.
func (paint paintCommon) Rel() error {
	return tvg_paint_rel(paint.paint_).error()
}

// Increment the reference count for the Tvg_Paint object.
//
// This method increases the reference count of Tvg_Paint object, allowing shared ownership and control over its lifetime.
//
//   - paint : The paint object to increase the reference count.
//
// Return: The updated reference count after the increment by 1.
//
// Warning: Please ensure that each call to tvg_paint_ref() is paired with a corresponding call to tvg_paint_unref() to prevent a dangling instance.
//
// see [Paint.Unref]
//
// see [Paint.GetRef]
//
// Since: 1.0
func (paint paintCommon) Ref() int {
	return int(tvg_paint_ref(paint.paint_))
}

// Decrement the reference count for the Tvg_Paint object.
//
// This method decreases the reference count of the Tvg_Paint object by 1. If the reference count reaches zero and the free flag is set to true, the instance is automatically deleted.
//
//   - paint : The paint object to decrease the reference count.
//   - free : Flag indicating whether to delete the Paint instance when the reference count reaches zero.
//
// Return: The updated reference count after the decrement.
//
// see [Paint.Ref]
//
// see [Paint.GetRef]
//
// Since: 1.0
func (paint paintCommon) Unref(free bool) int {
	return int(tvg_paint_unref(paint.paint_, free))
}

// Retrieve the current reference count of the Tvg_Paint object.
//
// This method provides the current reference count, allowing the user to check the shared ownership state of the Tvg_Paint object.
//
//   - paint : The paint object to return the reference count.
//
// Return: The current reference count of the Tvg_Paint object.
//
// see [Paint.Ref]
//
// see [Paint.Unref]
//
// Since: 1.0
func (paint paintCommon) GetRef() int {
	return int(tvg_paint_get_ref(paint.paint_))
}

// Sets the visibility of the Paint object.
//
// This is useful for selectively excluding paint objects during rendering.
//
//   - paint : The paint object to set the visibility status.
//   - on : A boolean flag indicating visibility. The initial value is true.               true, the object will be rendered by the engine.               false, the object will be excluded from the drawing process.
//
// Note: An invisible object is not considered inactive—it may still participate       in internal update processing if its properties are updated, but it will not       be taken into account for the final drawing output. To completely deactivate       a paint object, remove it from the canvas.
//
// see [Paint.GetVisible]
//
// see [Canvas.Remove]
//
// Since: 1.0
func (paint paintCommon) SetVisible(visible bool) error {
	return tvg_paint_set_visible(paint.paint_, visible).error()
}

// Gets the current visibility status of the Paint object.
//
//   - paint : The paint object to return the visibility status.
//
// Return: true if the object is visible and will be rendered.         false if the object is hidden and will not be rendered.
//
// see [Paint.SetVisible]
//
// Since: 1.0
func (paint paintCommon) GetVisible() bool {
	return tvg_paint_get_visible(paint.paint_)
}

// Gets the ID of the Paint object.
//
//   - paint : The paint object whose ID will be returned.
//
// Return: The ID of the paint object, or 0 if the ID is not set.
//
// see [Picture.GetPaint]
//
// see [Accessor.GenerateId]
//
// see [Paint.SetId]
//
// Since: 1.1
func (paint paintCommon) GetId() uint32 {
	return tvg_paint_get_id(paint.paint_)
}

// Sets the ID of the Paint object.
//
// The ID is used to specify a paint instance in a scene.
//
//   - paint : The paint object whose ID will be set.
//   - id : The ID to assign to the paint object.
//
// see [Picture.GetPaint]
//
// see [Accessor.GenerateId]
//
// see [Paint.GetId]
//
// Since: 1.1
func (paint paintCommon) SetId(id uint32) error {
	return tvg_paint_set_id(paint.paint_, id).error()
}

// Gets the user-defined data associated with the Paint object.
//
//   - paint : The paint object whose user-defined data will be returned.
//
// Return: The user-defined data, or NULL if none is set or paint is invalid.
//
// see [Paint.SetData]
//
// Note: Experimental API
func (paint paintCommon) GetData() uintptr {
	return tvg_paint_get_data(paint.paint_)
}

// Associates user-defined data with the Paint object.
//
// ThorVG does not interpret or manage the lifetime of this data.
//
//   - paint : The paint object whose user-defined data will be set.
//   - data : The user-defined data to associate with the paint object.
//
// see [Paint.GetData]
//
// Note: Experimental API
func (paint paintCommon) SetData(data uintptr) error {
	return tvg_paint_set_data(paint.paint_, data).error()
}

// Scales the given Tvg_Paint object by the given factor.
//
//   - paint : The paint object to be scaled.
//   - factor : The value of the scaling factor. The initial value is 1.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : in case a custom transform is applied.
//
// see [Paint.SetTransform]
func (paint paintCommon) Scale(factor float32) error {
	return tvg_paint_scale(paint.paint_, factor).error()
}

// Rotates the given Tvg_Paint by the given angle.
//
// The angle in measured clockwise from the horizontal axis. The rotational axis passes through the point on the object with zero coordinates.
//
//   - paint : The paint object to be rotated.
//   - degree : The value of the rotation angle in degrees.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : in case a custom transform is applied.
//
// see [Paint.SetTransform]
func (paint paintCommon) Rotate(degree float32) error {
	return tvg_paint_rotate(paint.paint_, degree).error()
}

// Moves the given Tvg_Paint in a two-dimensional space.
//
// The origin of the coordinate system is in the upper-left corner of the canvas. The horizontal and vertical axes point to the right and down, respectively.
//
//   - paint : The paint object to be shifted.
//   - x : The value of the horizontal shift.
//   - y : The value of the vertical shift.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : in case a custom transform is applied.
//
// see [Paint.SetTransform]
func (paint paintCommon) Translate(x float32, y float32) error {
	return tvg_paint_translate(paint.paint_, x, y).error()
}

// Transforms the given Tvg_Paint using the augmented transformation matrix.
//
// The augmented matrix of the transformation is expected to be given.
//
//   - paint : The paint object to be transformed.
//   - m : The 3x3 augmented matrix.
func (paint paintCommon) SetTransform(matrix Matrix) error {
	return tvg_paint_set_transform(paint.paint_, &matrix).error()
}

// Gets the matrix of the affine transformation of the given Tvg_Paint object.
//
// In case no transformation was applied, the identity matrix is returned.
//
//   - paint : The paint object of which to get the transformation matrix.
//   - m : The 3x3 augmented matrix.
func (paint paintCommon) GetTransform() (Matrix, error) {
	var matrix Matrix
	result := tvg_paint_get_transform(paint.paint_, &matrix)
	return matrix, result.error()
}

// Sets the opacity of the given Tvg_Paint.
//
//   - paint : The paint object of which the opacity value is to be set.
//   - opacity : The opacity value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
//
// Note: Setting the opacity with this API may require multiple renderings using a composition. It is recommended to avoid changing the opacity if possible.
func (paint paintCommon) SetOpacity(opacity uint8) error {
	return tvg_paint_set_opacity(paint.paint_, opacity).error()
}

// Gets the opacity of the given Tvg_Paint.
//
//   - paint : The paint object of which to get the opacity value.
//   - opacity : The opacity value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.
func (paint paintCommon) GetOpacity() (uint8, error) {
	var opacity uint8
	result := tvg_paint_get_opacity(paint.paint_, &opacity)
	return opacity, result.error()
}

// Duplicates the given Tvg_Paint object.
//
// Creates a new object and sets its all properties as in the original object.
//
//   - paint : The paint object to be copied.
//
// Return: A copied Tvg_Paint object if succeed, nullptr otherwise.
func (paint paintCommon) duplicate() paintCommon {
	return paintCommon{
		paint_: tvg_paint_duplicate(paint.paint_),
	}
}

// Checks whether a given region intersects the filled area of the paint.
//
// This function determines whether the specified rectangular region—defined by (`x`, `y`, `w`, `h`)— intersects the geometric fill region of the paint object.
// This is useful for hit-testing purposes, such as detecting whether a user interaction (e.g., touch or click) occurs within a painted region.
// The paint must be updated in a Canvas beforehand—typically after the Canvas has been drawn and synchronized.
//
//   - paint : The paint object to be tested.
//   - x : The x-coordinate of the top-left corner of the test region.
//   - y : The y-coordinate of the top-left corner of the test region.
//   - w : The width of the region to test. Must be greater than 0.
//   - h : The height of the region to test. Must be greater than 0.
//
// Return: true if any part of the region intersects the filled area; otherwise, false.
//
// Note: To test a single point, set the region size to w = 1, h = 1.
//
// Note: For efficiency, an AABB (axis-aligned bounding box) test is performed internally before precise hit detection.
//
// Note: This test does not take into account the results of blending or masking.
//
// Note: This test does take into account hidden paints as well.
//
// see [Paint.SetVisible]
// Since: 1.0
func (paint paintCommon) Intersects(x int, y int, width int, height int) bool {
	return tvg_paint_intersects(paint.paint_, int32(x), int32(y), int32(width), int32(height))
}

// Checks whether a given region intersects the filled area of the paint.
//
// This function determines whether the specified rectangular region—defined by (`x`, `y`, `w`, `h`)— intersects the geometric fill region of the paint object.
// This is useful for hit-testing purposes, such as detecting whether a user interaction (e.g., touch or click) occurs within a painted region.
// The paint must be updated in a Canvas beforehand—typically after the Canvas has been drawn and synchronized.
//
//   - paint : The paint object to be tested.
//   - x : The x-coordinate of the top-left corner of the test region.
//   - y : The y-coordinate of the top-left corner of the test region.
//   - w : The width of the region to test. Must be greater than 0.
//   - h : The height of the region to test. Must be greater than 0.
//   - visibleOnly : If true, hidden paints are excluded from the intersection test.
//
// Return: true if any part of the region intersects the filled area; otherwise, false.
//
// Note: To test a single point, set the region size to w = 1, h = 1.
//
// Note: This test does not take into account the results of blending or masking.
//
// see [Paint.SetVisible]
// Since: Experimental API
func (paint paintCommon) IntersectsRegion(x int, y int, width int, height int, visibleOnly bool) bool {
	return tvg_paint_intersects_region(paint.paint_, int32(x), int32(y), int32(width), int32(height), visibleOnly)
}

// Retrieves the axis-aligned bounding box (AABB) of the paint object in canvas space.
//
// Returns the bounding box of the paint as an axis-aligned bounding box (AABB), with all relevant transformations applied. The returned values x, y, w, h, may have invalid if the operation fails. Thus, please check the retval.
// This bounding box can be used to determine the actual rendered area of the object on the canvas, for purposes such as hit-testing, culling, or layout calculations.
//
//   - paint : The paint object of which to get the bounds.
//   - x : The x-coordinate of the upper-left corner of the bounding box.
//   - y : The y-coordinate of the upper-left corner of the bounding box.
//   - w : The width of the bounding box.
//   - h : The height of the bounding box.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If it failed to compute the bounding box (mostly due to invalid path information).
//
// see [Paint.GetObb]
//
// see [Canvas.Update]
func (paint paintCommon) GetAABB() (float32, float32, float32, float32, error) {
	var x, y float32
	var width, height float32
	result := tvg_paint_get_aabb(paint.paint_, &x, &y, &width, &height)
	return x, y, width, height, result.error()
}

// Retrieves the object-oriented bounding box (OBB) of the paint object in canvas space.
//
// This function returns the bounding box of the paint, as an oriented bounding box (OBB) after transformations are applied. The returned values pt4 may have invalid if the operation fails. Thus, please check the retval.
// This bounding box can be used to obtain the transformed bounding region in canvas space by taking the geometry's axis-aligned bounding box (AABB) in the object's local coordinate space and applying the object's transformations.
//
//   - paint : The paint object of which to get the bounds.
//   - pt4 : An array of four points representing the bounding box. The array size must be 4.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If it failed to compute the bounding box (mostly due to invalid path information).
//
// see [Paint.GetAabb]
//
// see [Canvas.Update]
//
// Since: 1.0
func (paint paintCommon) GetOBB() ([4]Point, error) {
	var points [4]Point
	result := tvg_paint_get_obb(paint.paint_, &points[0])
	return points, result.error()
}

// Sets the masking target object and the masking method.
//
//   - paint : The source object of the masking.
//   - target : The target object of the masking.
//   - method : The method used to mask the source object with the target.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : if the target has already belonged to another paint.
func (paint paintCommon) SetMaskMethod(target Paint, method MaskMethod) error {
	return tvg_paint_set_mask_method(paint.paint_, target.paint(), method).error()
}

// Gets the masking target object and the masking method.
//
//   - paint : The source object of the masking.
//   - target : The target object of the masking.
//   - method : The method used to mask the source object with the target.
func (paint paintCommon) GetMaskMethod(target Paint) (MaskMethod, error) {
	var method MaskMethod
	result := tvg_paint_get_mask_method(paint.paint_, target.paint(), &method)
	return method, result.error()
}

// Clip the drawing region of the paint object.
//
// This function restricts the drawing area of the paint object to the specified shape's paths.
//
//   - paint : The target object of the clipping.
//   - clipper : The shape object as the clipper.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : if the target has already belonged to another paint.
//
// Returns [RESULT_NOT_SUPPORTED] : If the clipper type is not Shape.
//
// see [Paint.GetClip]
//
// Since: 1.0
func (paint paintCommon) SetClip(clipper Paint) error {
	return tvg_paint_set_clip(paint.paint_, clipper.paint()).error()
}

// Get the clipper shape of the paint object.
//
// This function returns the clipper that has been previously set to this paint object.
//
// Return: The shape object used as the clipper, or nullptr if no clipper is set.
//
// see [Paint.SetClip]
//
// Since: 1.0
func (paint paintCommon) GetClip() (any, bool) {
	clipper := tvg_paint_get_clip(paint.paint_)
	if clipper == 0 {
		return nil, false
	}
	return newPaint(clipper)
}

/*
GetParent retrieves the parent paint object.

This function returns a pointer to the parent object if the current paint
belongs to one. Otherwise, it returns @c nullptr.

	@param[in] paint The Paint object of which to get the scene.

	@return A pointer to the parent object if available, otherwise @c nullptr.

	@see tvg_scene_add()
	@see tvg_canvas_add()

	@since 1.0
*/
func (paint paintCommon) GetParent() (any, bool) {
	return newPaint(paint.paint_)
}

// Gets the unique value of the paint instance indicating the instance type.
//
//   - paint : The paint object of which to get the type value.
//   - type : The unique type of the paint instance type.
//
// Since: 1.0
func (paint paintCommon) GetType() (Type, error) {
	var typ Type
	result := tvg_paint_get_type(paint.paint_, &typ)
	return typ, result.error()
}

// Sets the blending method for the paint object.
//
// The blending feature allows you to combine colors to create visually appealing effects, including transparency, lighting, shading, and color mixing, among others. its process involves the combination of colors or images from the source paint object with the destination (the lower layer image) using blending operations. The blending operation is determined by the chosen BlendMethod, which specifies how the colors or images are combined.
//
//   - paint : The paint object of which to set the blend method.
//   - method : The blending method to be set.
//
// Since: 0.15
func (paint paintCommon) SetBlendMethod(method BlendMethod) error {
	return tvg_paint_set_blend_method(paint.paint_, method).error()
}

// Gets the unique value of the paint instance indicating the instance type.
//
//   - paint : The paint object of which to get the type value.
//   - type : The unique type of the paint instance type.
//
// Since: 1.0
func newPaint(paint uintptr) (Paint, bool) {
	var typ Type
	result := tvg_paint_get_type(paint, &typ)
	if result != RESULT_SUCCESS {
		return nil, false
	}

	switch typ {
	case TYPE_TEXT:
		return Text{paintCommon: paintCommon{paint_: paint}}, true

	case TYPE_SHAPE:
		return Shape{paintCommon: paintCommon{paint_: paint}}, true

	case TYPE_SCENE:
		return Scene{paintCommon: paintCommon{paint_: paint}}, true

	case TYPE_PICTURE:
		return Picture{paintCommon: paintCommon{paint_: paint}}, true

	default:
		return nil, false
	}
}
