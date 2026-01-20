package thorvg

// #include "thorvg_capi.h"
import "C"

type Animation struct {
	animation C.Tvg_Animation
}

/*
AnimationNew creates a new Animation object.

	@return Tvg_Animation A new Tvg_Animation object.

	@since 0.13
*/
func AnimationNew() Animation {
	return Animation{C.tvg_animation_new()}
}

/*
SetFrame Specifies the current frame in the animation.

	@param[in] animation A Tvg_Animation pointer to the animation object.
	@param[in] no The index of the animation frame to be displayed. The index should be less than the tvg_animation_get_total_frame().

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION if the given @p no is the same as the current frame value.
	@retval TVG_RESULT_NOT_SUPPORTED The picture data does not support animations.

	@note For efficiency, ThorVG ignores updates to the new frame value if the difference from the current frame value
			is less than 0.001. In such cases, it returns @c Result::InsufficientCondition.
			Values less than 0.001 may be disregarded and may not be accurately retained by the Animation.
	@see tvg_animation_get_total_frame()

	@since 0.13
*/
func (animation Animation) SetFrame(no float32) error {
	result := C.tvg_animation_set_frame(animation.animation, C.float(no))
	return resultError(result)
}

/*
GetPicture retrieves a picture instance associated with this animation instance.

This function provides access to the picture instance that can be used to load animation formats, such as lot.
After setting up the picture, it can be added to the designated canvas, enabling control over animation frames
with this Animation instance.

	@param[in] animation A Tvg_Animation pointer to the animation object.

	@return A picture instance that is tied to this animation.

	@warning The picture instance is owned by Animation. It should not be deleted manually.

	@since 0.13
*/
func (animation Animation) GetPicture() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: C.tvg_animation_get_picture(animation.animation),
		},
	}
}

/*
GetFrame retrieves the current frame number of the animation.

	@param[in] animation A Tvg_Animation pointer to the animation object.
	@param[in] no The current frame number of the animation, between 0 and totalFrame() - 1.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer or @p no

	@see tvg_animation_get_total_frame()
	@see tvg_animation_set_frame()

	@since 0.13
*/
func (animation Animation) GetFrame() (float32, error) {
	var no C.float
	result := C.tvg_animation_get_frame(animation.animation, &no)
	return float32(no), resultError(result)
}

/*
GetTotalFrame retrieves the total number of frames in the animation.

	@param[in] animation A Tvg_Animation pointer to the animation object.
	@param[in] cnt The total number of frames in the animation.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer or @p cnt.

	@note Frame numbering starts from 0.
	@note If the Picture is not properly configured, this function will return 0.

	@since 0.13
*/
func (animation Animation) GetTotalFrame() (float32, error) {
	var cnt C.float
	result := C.tvg_animation_get_total_frame(animation.animation, &cnt)
	return float32(cnt), resultError(result)
}

/*
Duration Retrieves the duration of the animation in seconds.

	@param[in] animation A Tvg_Animation pointer to the animation object.
	@param[in] duration The duration of the animation in seconds.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer or @p duration.

	@note If the Picture is not properly configured, this function will return 0.

	@since 0.13
*/
func (animation Animation) GetDuration() (float32, error) {
	var duration C.float
	result := C.tvg_animation_get_duration(animation.animation, &duration)
	return float32(duration), resultError(result)
}

/*
SetSegment specifies the playback segment of the animation.

The set segment is designated as the play area of the animation.
This is useful for playing a specific segment within the entire animation.
After setting, the number of animation frames and the playback time are calculated
by mapping the playback segment as the entire range.

	@param[in] animation The Tvg_Animation pointer to the animation object.
	@param[in] begin segment begin frame.
	@param[in] end segment end frame.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT If the @p begin is higher than @p end.

	@note Animation allows a range from 0.0 to the total frame. @p end should not be higher than @p begin.
	@note If a marker has been specified, its range will be disregarded.

	@see tvg_lottie_animation_set_marker()
	@see tvg_animation_get_total_frame()

	@since 1.0
*/
func (animation Animation) SetSegment(begin float32, end float32) error {
	result := C.tvg_animation_set_segment(animation.animation, C.float(begin), C.float(end))
	return resultError(result)
}

/*
GetSegment gets the current segment range information.

	@param[in] animation The Tvg_Animation pointer to the animation object.
	@param[out] begin segment begin frame.
	@param[out] end segment end frame.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION In case the animation is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer.

	@since 1.0
*/
func (animation Animation) GetSegment() (float32, float32, error) {
	var begin C.float
	var end C.float
	result := C.tvg_animation_get_segment(animation.animation, &begin, &end)
	return float32(begin), float32(end), resultError(result)
}

/*
Del deletes the given Tvg_Animation object.

	@param[in] animation The Tvg_Animation object to be deleted.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Animation pointer.

	@since 0.13
*/
func (animation Animation) Del() error {
	result := C.tvg_animation_del(animation.animation)
	return resultError(result)
}
