package thorvg

type Animation uintptr

// Creates a new Animation object.
//
// Return: Tvg_Animation A new Tvg_Animation object.
//
// Since: 0.13
func AnimationNew() Animation {
	return Animation(tvg_animation_new())
}

// Specifies the current frame in the animation.
//
//   - animation : An animation object.
//   - no : The index of the animation frame to be displayed. The index should be less than the tvg_animation_get_total_frame().
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : if the given no is the same as the current frame value.
//
// Returns [RESULT_NOT_SUPPORTED] : The picture data does not support animations.
//
// Note: For efficiency, ThorVG ignores updates to the new frame value if the difference from the current frame value       is less than 0.001. In such cases, it returns Result::InsufficientCondition.       Values less than 0.001 may be disregarded and may not be accurately retained by the Animation.
//
// see [Animation.GetTotalFrame]
//
// Since: 0.13
func (animation Animation) SetFrame(no float32) error {
	return tvg_animation_set_frame(animation, no).error()
}

// Retrieves a picture instance associated with this animation instance.
//
// This function provides access to the picture instance that can be used to load animation formats, such as lot. After setting up the picture, it can be added to the designated canvas, enabling control over animation frames with this Animation instance.
//
//   - animation : A animation object to the animation object.
//
// Return: A picture instance that is tied to this animation.
//
// Warning: The picture instance is owned by Animation. It should not be deleted manually.
//
// Since: 0.13
func (animation Animation) GetPicture() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: tvg_animation_get_picture(animation),
		},
	}
}

// Retrieves the current frame number of the animation.
//
//   - animation : An animation object.
//   - no : The current frame number of the animation, between 0 and totalFrame() - 1.
//
// Returns [RESULT_INVALID_ARGUMENT] : An invalid no pointer.
//
// see [Animation.GetTotalFrame]
//
// see [Animation.SetFrame]
//
// Since: 0.13
func (animation Animation) GetFrame() (float32, error) {
	var no float32
	result := tvg_animation_get_frame(animation, &no)
	return no, result.error()
}

// Retrieves the total number of frames in the animation.
//
//   - animation : An animation object.
//   - cnt : The total number of frames in the animation.
//
// Returns [RESULT_INVALID_ARGUMENT] : An invalid cnt pointer.
//
// Note: Frame numbering starts from 0.
//
// Note: If the Picture is not properly configured, this function will return 0.
//
// Since: 0.13
func (animation Animation) GetTotalFrame() (float32, error) {
	var cnt float32
	result := tvg_animation_get_total_frame(animation, &cnt)
	return cnt, result.error()
}

// Retrieves the duration of the animation in seconds.
//
//   - animation : An animation object.
//   - duration : The duration of the animation in seconds.
//
// Returns [RESULT_INVALID_ARGUMENT] : An invalid duration pointer.
//
// Note: If the Picture is not properly configured, this function will return 0.
//
// Since: 0.13
func (animation Animation) GetDuration() (float32, error) {
	var duration float32
	result := tvg_animation_get_duration(animation, &duration)
	return duration, result.error()
}

// Specifies the playback segment of the animation.
//
// The set segment is designated as the play area of the animation. This is useful for playing a specific segment within the entire animation. After setting, the number of animation frames and the playback time are calculated by mapping the playback segment as the entire range.
//
//   - animation : The animation object.
//   - begin : segment begin frame.
//   - end : segment end frame.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// Returns [RESULT_INVALID_ARGUMENT] : If the begin is higher than end.
//
// Note: Animation allows a range from 0.0 to the total frame. end should not be lower than begin.
//
// Note: If a marker has been specified, its range will be disregarded.
//
// see [Animation.SetMarker]
//
// see [Animation.GetTotalFrame]
//
// Since: 1.0
func (animation Animation) SetSegment(begin float32, end float32) error {
	return tvg_animation_set_segment(animation, begin, end).error()
}

// Gets the current segment range information.
//
//   - animation : The animation object.
//   - begin : segment begin frame.
//   - end : segment end frame.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// Since: 1.0
func (animation Animation) GetSegment() (float32, float32, error) {
	var begin float32
	var end float32
	result := tvg_animation_get_segment(animation, &begin, &end)
	return begin, end, result.error()
}

// Deletes the given Tvg_Animation object.
//
//   - animation : The Tvg_Animation object to be deleted.
//
// Since: 0.13
func (animation Animation) Del() error {
	return tvg_animation_del(animation).error()
}
