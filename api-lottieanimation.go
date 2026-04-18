package thorvg

/*
LottieAnimationNew Creates a new LottieAnimation object.

	@return Tvg_Animation A new Tvg_LottieAnimation object.

	@since 0.15
*/
func LottieAnimationNew() Animation {
	return Animation(tvg_lottie_animation_new())
}

/*
AnimationGenSlot generates a new slot from the given slot data.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] slot The Lottie slot data in JSON format.

	@return The generated slot ID when successful, 0 otherwise.

	@since 1.0
*/
func (animation Animation) AnimationGenSlot(slot string) uint32 {
	return tvg_lottie_animation_gen_slot(uintptr(animation), slot)
}

/*
ApplySlot applies a previously generated slot to the animation.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] id The ID of the slot to apply, or 0 to reset all slots.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT When the given @p id is invalid
	@retval TVG_RESULT_NOT_SUPPORTED The Lottie Animation is not supported.

	@since 1.0
*/
func (animation Animation) ApplySlot(id uint) error {
	return tvg_lottie_animation_apply_slot(uintptr(animation), uint32(id)).error()
}

/*
DelSlot deletes a previously generated slot.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] id The ID of the slot to delete.

	@return Tvg_Result enumeration.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded or the slot ID is invalid.
	@retval TVG_RESULT_NOT_SUPPORTED The Lottie Animation is not supported.

	@note This function should be paired with gen.
	@see tvg_lottie_animation_gen_slot()
	@since 1.0
*/
func (animation Animation) DelSlot(id uint) error {
	return tvg_lottie_animation_del_slot(uintptr(animation), uint32(id)).error()
}

/*
SetMarker specifies a segment by marker.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] marker The name of the segment marker.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT When the given @p marker is invalid.
	@retval TVG_RESULT_NOT_SUPPORTED The Lottie Animation is not supported.

	@since 1.0
*/
func (animation Animation) SetMarker(marker string) error {
	return tvg_lottie_animation_set_marker(uintptr(animation), marker).error()
}

/*
GetMarkersCnt gets the marker count of the animation.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[out] cnt The count value of the markers.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.

	@since 1.0
*/
func (animation Animation) GetMarkersCnt() (uint, error) {
	var cnt uint32
	result := tvg_lottie_animation_get_markers_cnt(uintptr(animation), &cnt)
	return uint(cnt), result.error()
}

/*
GetMarker gets the marker name by a given index.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] idx The index of the animation marker, starts from 0.
	@param[out] name The name of marker when succeed.

	@retval TVG_RESULT_INVALID_ARGUMENT In case @c nullptr is passed as the argument or @c idx is out of range.

	@since 1.0
*/
func (animation Animation) GetMarker(idx uint) (string, error) {
	var name *byte
	result := tvg_lottie_animation_get_marker(uintptr(animation), uint32(idx), &name)
	if result != RESULT_SUCCESS {
		return "", result.error()
	}
	return goString(name), nil
}

/*
GetMarkerInfo retrieves marker information by index.

	@param[in] animation The Lottie animation object.
	@param[in] idx The zero-based index of the animation marker.
	@param[out] name Pointer to receive the marker name.
									Pass @c nullptr if the value is not required.
	@param[out] begin Pointer to receive the marker's starting frame.
										Pass @c nullptr if the value is not required.
	@param[out] end Pointer to receive the marker's ending frame.
									Pass @c nullptr if the value is not required.

	@retval TVG_RESULT_INVALID_ARGUMENT if @p idx is out of range.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.

	@see tvg_lottie_animation_get_markers_cnt()
	@note Experimental API
*/
func (animation Animation) GetMarkerInfo(idx uint) (string, float32, float32, error) {
	var name *byte
	var begin float32
	var end float32
	result := tvg_lottie_animation_get_marker_info(uintptr(animation), uint32(idx), &name, &begin, &end)
	if result != RESULT_SUCCESS {
		return "", 0, 0, result.error()
	}
	return goString(name), begin, end, nil
}

// TVG_API Tvg_Result tvg_lottie_animation_get_marker_info(Tvg_Animation animation, uint32_t idx, const char** name, float* begin, float* end);

/*
Tween interpolates between two frames over a specified duration.

This method performs tweening, a process of generating intermediate frame
between @p from and @p to based on the given @p progress.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] from The start frame number of the interpolation.
	@param[in] to The end frame number of the interpolation.
	@param[in] progress The current progress of the interpolation (range: 0.0 to 1.0).

	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.

	@since 1.0
*/
func (animation Animation) Tween(from float32, to float32, progress float32) error {
	return tvg_lottie_animation_tween(uintptr(animation), from, to, progress).error()
}

/*
Assign updates the value of an expression variable for a specific layer.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] layer The name of the layer containing the variable to be updated.
	@param[in] ix The property index of the variable within the layer.
	@param[in] var The name of the variable to be updated.
	@param[in] val The new value to assign to the variable.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT When the given parameter is invalid.
	@retval TVG_RESULT_NOT_SUPPORTED When neither the layer nor the property is found in the current animation.

	@note Experimental API
*/
func (animation Animation) Assign(layer string, ix uint, var_ string, val float32) error {
	return tvg_lottie_animation_assign(uintptr(animation), layer, uint32(ix), var_, val).error()
}

/*
SetQuality sets the quality level for Lottie effects.

This function controls the rendering quality of effects like blur, shadows, etc.
Lower values prioritize performance while higher values prioritize quality.

	@param[in] animation The Tvg_Animation pointer to the Lottie animation object.
	@param[in] value The quality level (0-100). 0 represents lowest quality/best performance,
								100 represents highest quality/lowest performance, default is 50.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer.

	@note This option is used as a hint; its behavior heavily depends on the render backend.

	@since 1.0
*/
func (animation Animation) SetQuality(value uint) error {
	return tvg_lottie_animation_set_quality(uintptr(animation), uint8(value)).error()
}
