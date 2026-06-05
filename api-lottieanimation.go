package thorvg

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

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

/*
AudioInfo describes the current state of a Lottie audio layer.

This structure is provided to the audio resolver callback and contains
the information required to synchronize audio playback with the animation
timeline. Applications are responsible for managing audio playback using
their own audio engine.

Example:

	@code
		void on_audio(const Tvg_Audio_Info* info, void* data)
		{
		    if (info->active) {
		        // Start or seek playback of info->src.
		    } else {
		        // Stop playback of info->src.
		    }
		}
		@endcode
		@see tvg_lottie_animation_set_audio_resolver()
		@note Experimental API
*/
type AudioInfo struct {
	Src      string  // Audio source: a file path/URL or embedded raw bytes.
	MimeType string  // MIME type string; valid when @c embedded; may be @c NULL.
	Size     uint32  // Embedded data size in bytes; valid when @c embedded.
	Offset   float32 // Position within the audio file in seconds; valid when @c active.
	Volume   float32 // Volume [0, 100]; valid when @c active.
	Active   bool    // @c true while the layer is within its playback range.
	Embedded bool    // @c true if @p src points to embedded audio data; @c false if it is a file path or URL.
}
type audioInfo struct {
	src      *byte   // Audio source: a file path/URL or embedded raw bytes.
	mimeType *byte   // MIME type string; valid when @c embedded; may be @c NULL.
	size     uint32  // Embedded data size in bytes; valid when @c embedded.
	offset   float32 // Position within the audio file in seconds; valid when @c active.
	volume   float32 // Volume [0, 100]; valid when @c active.
	active   bool    // @c true while the layer is within its playback range.
	embedded bool    // @c true if @p src points to embedded audio data; @c false if it is a file path or URL.
}

/*
AudioResolver is a callback invoked to provide audio playback information for a Lottie animation.

Applications can use this callback to synchronize external audio
playback with the animation timeline.

	@param[in] info Audio information for the current timeline state.
	@param[in] data User data specified when registering the callback.

	@see tvg_lottie_animation_set_audio_resolver()
	@note Experimental API.
*/
type AudioResolver func(info AudioInfo)

/*
SetAudioResolver sets the audio resolver callback for Lottie audio layers.

The resolver is invoked whenever the playback state of an audio layer changes.
It allows applications to synchronize audio playback with the animation timeline.

	@param[in] animation A Lottie animation object.
	@param[in] resolver A user-defined callback that receives audio playback state updates.
	@param[in] data User data passed to @p resolver.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION The animation has not been loaded.

	@note To disable audio notifications, pass @c nullptr as @p resolver.
	@note Experimental API.

	@see Tvg_Audio_Resolver
*/
func (animation Animation) SetAudioResolver(resolver AudioResolver) error {
	puregoResolver := purego.NewCallback(func(cAudioInfo uintptr, _data uintptr) {
		cInfo := (*audioInfo)(unsafe.Pointer(cAudioInfo))
		info := AudioInfo{
			Src:      goString(cInfo.src),
			MimeType: goString(cInfo.mimeType),
			Size:     cInfo.size,
			Offset:   cInfo.offset,
			Volume:   cInfo.volume,
			Active:   cInfo.active,
			Embedded: cInfo.embedded,
		}
		resolver(info)
	})

	return tvg_lottie_animation_set_audio_resolver(uintptr(animation), puregoResolver, 0).error()
}
