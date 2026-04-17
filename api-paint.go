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

/*
Rel safely releases a Tv_Paint object.

This is the counterpart to the `new()` API, and releases the given Paint object safely,
handling @c nullptr and managing ownership properly.

	@param[in] paint A Tvg_Paint object to release.
*/
func (paint paintCommon) Rel() error {
	return tvg_paint_rel(paint.paint_).error()
}

/*
Ref increments the reference count for the Tvg_Paint object.

This method increases the reference count of Tvg_Paint object, allowing shared ownership and control over its lifetime.

	@param[in] paint The Tvg_Paint object to increase the reference count.

	@return The updated reference count after the increment by 1.

	@warning Please ensure that each call to tvg_paint_ref() is paired with a corresponding call to tvg_paint_unref() to prevent a dangling instance.

	@see tvg_paint_unref()
	@see tvg_paint_get_ref()

	@since 1.0
*/
func (paint paintCommon) Ref() int {
	return int(tvg_paint_ref(paint.paint_))
}

/*
Unref decrements the reference count for the Tvg_Paint object.

This method decreases the reference count of the Tvg_Paint object by 1.
If the reference count reaches zero and the @p free flag is set to true, the instance is automatically deleted.

	@param[in] paint The Tvg_Paint object to decrease the reference count.
	@param[in] free Flag indicating whether to delete the Paint instance when the reference count reaches zero.

	@return The updated reference count after the decrement.

	@see tvg_paint_ref()
	@see tvg_paint_get_ref()

	@since 1.0
*/
func (paint paintCommon) Unref(free bool) int {
	return int(tvg_paint_unref(paint.paint_, free))
}

/*
GetRef retrieves the current reference count of the Tvg_Paint object.

This method provides the current reference count, allowing the user to check the shared ownership state of the Tvg_Paint object.

	@param[in] paint The Tvg_Paint object to return the reference count.

	@return The current reference count of the Tvg_Paint object.

	@see tvg_paint_ref()
	@see tvg_paint_unref()

	@since 1.0
*/
func (paint paintCommon) GetRef() int {
	return int(tvg_paint_get_ref(paint.paint_))
}

/*
SetVisible sets the visibility of the Paint object.

This is useful for selectively excluding paint objects during rendering.

	@param[in] paint The Tvg_Paint object to set the visibility status.
	@param[in] on A boolean flag indicating visibility. The default is @c true.
							@c true, the object will be rendered by the engine.
							@c false, the object will be excluded from the drawing process.

	@note An invisible object is not considered inactive—it may still participate
				in internal update processing if its prope

				rties are updated, but it will not
				be taken into account for the final drawing output. To completely deactivate
				a paint object, remove it from the canvas.

	@see tvg_paint_get_visible()
	@see tvg_canvas_remove()

	@since 1.0
*/
func (paint paintCommon) SetVisible(visible bool) error {
	return tvg_paint_set_visible(paint.paint_, visible).error()
}

/*
GetVisible gets the current visibility status of the Paint object.

	@param[in] paint The Tvg_Paint object to return the visibility status.

	@return true if the object is visible and will be rendered.
					false if the object is hidden and will not be rendered.

	@see tvg_paint_set_visible()

	@since 1.0
*/
func (paint paintCommon) GetVisible() bool {
	return tvg_paint_get_visible(paint.paint_)
}

/*
GetId gets the ID of the Paint object.

	@param[in] paint The paint object whose ID will be returned.

	@return The ID of the paint object, or 0 if the ID is not set.

	@see tvg_picture_get_paint()
	@see tvg_accessor_generate_id()
	@see tvg_paint_set_id()

	@note Experimental API
*/
func (paint paintCommon) GetId() uint32 {
	return tvg_paint_get_id(paint.paint_)
}

/*
SetId sets the ID of the Paint object.

The ID is used to specify a paint instance in a scene.

	@param[in] paint The paint object whose ID will be set.
	@param[in] id The ID to assign to the paint object.

	@see tvg_picture_get_paint()
	@see tvg_accessor_generate_id()
	@see tvg_paint_get_id()

	@note Experimental API
*/
func (paint paintCommon) SetId(id uint32) error {
	return tvg_paint_set_id(paint.paint_, id).error()
}

/*
Scale scales the given Tvg_Paint object by the given factor.

	@param[in] paint The Tvg_Paint object to be scaled.
	@param[in] factor The value of the scaling factor. The default value is 1.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION in case a custom transform is applied.

	@see tvg_paint_set_transform()
*/
func (paint paintCommon) Scale(factor float32) error {
	return tvg_paint_scale(paint.paint_, factor).error()
}

/*
Rotate rotates the given Tvg_Paint by the given angle.

The angle in measured clockwise from the horizontal axis.
The rotational axis passes through the point on the object with zero coordinates.

	@param[in] paint The Tvg_Paint object to be rotated.
	@param[in] degree The value of the rotation angle in degrees.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION in case a custom transform is applied.

	@see tvg_paint_set_transform()
*/
func (paint paintCommon) Rotate(degree float32) error {
	return tvg_paint_rotate(paint.paint_, degree).error()
}

/*
Translate moves the given Tvg_Paint in a two-dimensional space.

The origin of the coordinate system is in the upper-left corner of the canvas.
The horizontal and vertical axes point to the right and down, respectively.

	@param[in] paint The Tvg_Paint object to be shifted.
	@param[in] x The value of the horizontal shift.
	@param[in] y The value of the vertical shift.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION in case a custom transform is applied.

	@see tvg_paint_set_transform()
*/
func (paint paintCommon) Translate(x float32, y float32) error {
	return tvg_paint_translate(paint.paint_, x, y).error()
}

/*
SetTransform transforms the given Tvg_Paint using the augmented transformation matrix.

The augmented matrix of the transformation is expected to be given.

	@param[in] paint The Tvg_Paint object to be transformed.
	@param[in] m The 3x3 augmented matrix.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr is passed as the argument.
*/
func (paint paintCommon) SetTransform(matrix Matrix) error {
	return tvg_paint_set_transform(paint.paint_, &matrix).error()
}

/*
GetTransform gets the matrix of the affine transformation of the
given Tvg_Paint object.

In case no transformation was applied, the identity matrix is returned.

	@param[in] paint The Tvg_Paint object of which to get the transformation matrix.
	@param[out] m The 3x3 augmented matrix.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr is passed as the argument.
*/
func (paint paintCommon) Ret_transform() (Matrix, error) {
	var matrix Matrix
	result := tvg_paint_get_transform(paint.paint_, &matrix)
	return matrix, result.error()
}

/*
SetOpacity sets the opacity of the given Tvg_Paint.

	@param[in] paint The Tvg_Paint object of which the opacity value is to be set.
	@param[in] opacity The opacity value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note Setting the opacity with this API may require multiple renderings using a composition. It is recommended to avoid changing the opacity if possible.
*/
func (paint paintCommon) SetOpacity(opacity uint8) error {
	return tvg_paint_set_opacity(paint.paint_, opacity).error()
}

/*
GetOpacity gets the opacity of the given Tvg_Paint.

	@param[in] paint The Tvg_Paint object of which to get the opacity value.
	@param[out] opacity The opacity value in the range [0 ~ 255], where 0 is completely transparent and 255 is opaque.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.
*/
func (paint paintCommon) GetOpacity() (uint8, error) {
	var opacity uint8
	result := tvg_paint_get_opacity(paint.paint_, &opacity)
	return opacity, result.error()
}

/*
duplicate duplicates the given Tvg_Paint object.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (paint paintCommon) duplicate() paintCommon {
	return paintCommon{
		paint_: tvg_paint_duplicate(paint.paint_),
	}
}

/*
Intersects checks whether a given region intersects the filled area of the paint.

This function determines whether the specified rectangular region—defined by (`x`, `y`, `w`, `h`)—
intersects the geometric fill region of the paint object.

This is useful for hit-testing purposes, such as detecting whether a user interaction (e.g., touch or click)
occurs within a painted region.

The paint must be updated in a Canvas beforehand—typically after the Canvas has been
drawn and synchronized.

	@param[in] paint A Tvg_Paint pointer to the shape object to be tested.
	@param[in] x The x-coordinate of the top-left corner of the test region.
	@param[in] y The y-coordinate of the top-left corner of the test region.
	@param[in] w The width of the region to test. Must be greater than 0; defaults to 1.
	@param[in] h The height of the region to test. Must be greater than 0; defaults to 1.

	@return @c true if any part of the region intersects the filled area; otherwise, @c false.

	@note To test a single point, set the region size to w = 1, h = 1.
	@note For efficiency, an AABB (axis-aligned bounding box) test is performed internally before precise hit detection.
	@note This test does not take into account the results of blending or masking.
	@note This test does take into account the the hidden paints as well. @see tvg_paint_set_visible().
	@since 1.0
*/
func (paint paintCommon) Intersects(x int, y int, width int, height int) bool {
	return tvg_paint_intersects(paint.paint_, int32(x), int32(y), int32(width), int32(height))
}

/*
GetAABB retrieves the axis-aligned bounding box (AABB) of the paint object in canvas space.

Returns the bounding box of the paint as an axis-aligned bounding box (AABB), with all relevant transformations applied.
The returned values @p x, @p y, @p w, @p h, may have invalid if the operation fails. Thus, please check the retval.

This bounding box can be used to determine the actual rendered area of the object on the canvas,
for purposes such as hit-testing, culling, or layout calculations.

	@param[in] paint The Tvg_Paint object of which to get the bounds.
	@param[out] x The x-coordinate of the upper-left corner of the bounding box.
	@param[out] y The y-coordinate of the upper-left corner of the bounding box.
	@param[out] w The width of the bounding box.
	@param[out] h The height of the bounding box.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid @p paint.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION If it failed to compute the bounding box (mostly due to invalid path information).

	@see tvg_paint_get_obb()
	@see tvg_canvas_update()
*/
func (paint paintCommon) GetAABB() (float32, float32, float32, float32, error) {
	var x, y float32
	var width, height float32
	result := tvg_paint_get_aabb(paint.paint_, &x, &y, &width, &height)
	return x, y, width, height, result.error()
}

/*
GetOBB retrieves the object-oriented bounding box (OBB) of the paint object in canvas space.

This function returns the bounding box of the paint, as an oriented bounding box (OBB) after transformations are applied.
The returned values @p pt4 may have invalid if the operation fails. Thus, please check the retval.

This bounding box can be used to obtain the transformed bounding region in canvas space
by taking the geometry's axis-aligned bounding box (AABB) in the object's local coordinate space
and applying the object's transformations.

	@param[in] paint The Tvg_Paint object of which to get the bounds.
	@param[out] pt4 An array of four points representing the bounding box. The array size must be 4.

	@retval TVG_RESULT_INVALID_ARGUMENT @p paint or @p pt4 is invalid.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION If it failed to compute the bounding box (mostly due to invalid path information).

	@see tvg_paint_get_aabb()
	@see tvg_canvas_update()

	@since 1.0
*/
func (paint paintCommon) GetOBB() ([4]Point, error) {
	var points [4]Point
	result := tvg_paint_get_obb(paint.paint_, &points[0])
	return points, result.error()
}

/*
SetMaskMethod sets the masking target object and the masking method.

	@param[in] paint The source object of the masking.
	@param[in] target The target object of the masking.
	@param[in] method The method used to mask the source object with the target.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION if the target has already belonged to another paint.
*/
func (paint paintCommon) SetMaskMethod(target Paint, method MaskMethod) error {
	return tvg_paint_set_mask_method(paint.paint_, target.paint(), method).error()
}

/*
GetMaskMethod gets the masking target object and the masking method.

	@param[in] paint The source object of the masking.
	@param[out] target The target object of the masking.
	@param[out] method The method used to mask the source object with the target.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr is passed as the argument.
*/
func (paint paintCommon) GetMaskMethod(target Paint) (MaskMethod, error) {
	var method MaskMethod
	result := tvg_paint_get_mask_method(paint.paint_, target.paint(), &method)
	return method, result.error()
}

/*
SetClip clips the drawing region of the paint object.

This function restricts the drawing area of the paint object to the specified shape's paths.

	@param[in] paint The target object of the clipping.
	@param[in] clipper The shape object as the clipper.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION if the target has already belonged to another paint.
	@retval TVG_RESULT_NOT_SUPPORTED If the @p clipper type is not Shape.

	@see tvg_paint_get_clip()

	@since 1.0
*/
func (paint paintCommon) SetClip(clipper Paint) error {
	return tvg_paint_set_clip(paint.paint_, clipper.paint()).error()
}

/*
GetClip gets the clipper shape of the paint object.

This function returns the clipper that has been previously set to this paint object.

	@return The shape object used as the clipper, or @c nullptr if no clipper is set.

	@see tvg_paint_set_clip()

	@since 1.0
*/
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

	@param[in] paint The Tvg_Paint object of which to get the scene.

	@return A pointer to the parent object if available, otherwise @c nullptr.

	@see tvg_scene_add()
	@see tvg_canvas_add()

	@since 1.0
*/
func (paint paintCommon) GetParent() (any, bool) {
	return newPaint(paint.paint_)
}

/*
GetType gets the unique value of the paint instance indicating the instance type.

	@param[in] paint The Tvg_Paint object of which to get the type value.
	@param[out] type The unique type of the paint instance type.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.

	@since 1.0
*/
func (paint paintCommon) GetType() (Type, error) {
	var typ Type
	result := tvg_paint_get_type(paint.paint_, &typ)
	return typ, result.error()
}

/*
SetBlendMethod sets the blending method for the paint object.

The blending feature allows you to combine colors to create visually appealing effects, including transparency, lighting, shading, and color mixing, among others.
its process involves the combination of colors or images from the source paint object with the destination (the lower layer image) using blending operations.
The blending operation is determined by the chosen @p BlendMethod, which
specifies how the colors or images are combined.

	@param[in] paint The Tvg_Paint object of which to set the blend method.
	@param[in] method The blending method to be set.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.

	@since 0.15
*/
func (paint paintCommon) SetBlendMethod(method BlendMethod) error {
	return tvg_paint_set_blend_method(paint.paint_, method).error()
}

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
