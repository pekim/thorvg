package thorvg

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// Creates a new LottieAnimation object.
//
// Return: Tvg_Animation A new Tvg_LottieAnimation object.
//
// Since: 0.15
func LottieAnimationNew() Animation {
	return Animation(tvg_lottie_animation_new())
}

// Generates a new slot from the given slot data.
//
//   - animation : The Lottie animation object.
//   - slot : The Lottie slot data in JSON format.
//
// Return: The generated slot ID when successful, 0 otherwise.
//
// Since: 1.0
func (animation Animation) GenSlot(slot string) uint32 {
	return tvg_lottie_animation_gen_slot(animation, slot)
}

// Applies a previously generated slot to the animation.
//
//   - animation : The Lottie animation object.
//   - id : The ID of the slot to apply, or 0 to reset all slots.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// Returns [RESULT_INVALID_ARGUMENT] : When the given id is invalid
//
// Returns [RESULT_NOT_SUPPORTED] : The Lottie Animation is not supported.
//
// Since: 1.0
func (animation Animation) ApplySlot(id uint) error {
	return tvg_lottie_animation_apply_slot(animation, uint32(id)).error()
}

// Deletes a previously generated slot.
//
//   - animation : The Lottie animation object.
//   - id : The ID of the slot to delete.
//
// Return: Tvg_Result enumeration.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded or the slot ID is invalid.
//
// Returns [RESULT_NOT_SUPPORTED] : The Lottie Animation is not supported.
//
// Note: This function should be paired with gen.
//
// see [Animation.GenSlot]
// Since: 1.0
func (animation Animation) DelSlot(id uint) error {
	return tvg_lottie_animation_del_slot(animation, uint32(id)).error()
}

// Specifies a segment by marker.
//
//   - animation : The Lottie animation object.
//   - marker : The name of the segment marker.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// Returns [RESULT_INVALID_ARGUMENT] : When the given marker is invalid.
//
// Returns [RESULT_NOT_SUPPORTED] : The Lottie Animation is not supported.
//
// Since: 1.0
func (animation Animation) SetMarker(marker string) error {
	return tvg_lottie_animation_set_marker(animation, marker).error()
}

// Gets the marker count of the animation.
//
//   - animation : The Lottie animation object.
//   - cnt : The count value of the markers.
//
// Returns [RESULT_INVALID_ARGUMENT] : In case a nullptr is passed as the argument.
//
// Since: 1.0
func (animation Animation) GetMarkersCnt() (uint, error) {
	var cnt uint32
	result := tvg_lottie_animation_get_markers_cnt(animation, &cnt)
	return uint(cnt), result.error()
}

// Deprecate: see tvg_lottie_animation_get_marker_info()
func (animation Animation) GetMarker(idx uint) (string, error) {
	var name *byte
	result := tvg_lottie_animation_get_marker(animation, uint32(idx), &name)
	if result != RESULT_SUCCESS {
		return "", result.error()
	}
	return goString(name), nil
}

// Retrieves marker information by index.
//
//   - animation : The Lottie animation object.
//   - idx : The zero-based index of the animation marker.
//   - name : Pointer to receive the marker name.                  Pass nullptr if the value is not required.
//   - begin : Pointer to receive the marker's starting frame.                   Pass nullptr if the value is not required.
//   - end : Pointer to receive the marker's ending frame.                 Pass nullptr if the value is not required.
//
// Returns [RESULT_INVALID_ARGUMENT] : if idx is out of range.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// see [Animation.GetMarkersCnt]
// Since: 1.1
func (animation Animation) GetMarkerInfo(idx uint) (string, float32, float32, error) {
	var name *byte
	var begin float32
	var end float32
	result := tvg_lottie_animation_get_marker_info(animation, uint32(idx), &name, &begin, &end)
	if result != RESULT_SUCCESS {
		return "", 0, 0, result.error()
	}
	return goString(name), begin, end, nil
}

// Interpolates between two frames over a specified duration.
//
// This method performs tweening, a process of generating intermediate frame between from and to based on the given progress.
//
//   - animation : The Lottie animation object.
//   - from : The start frame number of the interpolation.
//   - to : The end frame number of the interpolation.
//   - progress : The current progress of the interpolation (range: 0.0 to 1.0).
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : In case the animation is not loaded.
//
// Since: 1.0
func (animation Animation) Tween(from float32, to float32, progress float32) error {
	return tvg_lottie_animation_tween(animation, from, to, progress).error()
}

// Sets the target frame for dynamic tweening.
//
// This method starts a dynamic interpolation from the current animation frame toward to. Use tvg_lottie_animation_tween_go() to update the interpolation progress.
//
//   - animation : The Lottie animation object.
//   - to : The target frame number of the interpolation.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the animation is not loaded.
//
// Note: The dynamic tweening set by this method is discarded when [Animation.SetFrame]       or [Animation.Tween] is called.
//
// see [Animation.TweenGo]
//
// Note: Experimental API
func (animation Animation) TweenTo(to float32) error {
	return tvg_lottie_animation_tween_to(animation, to).error()
}

// Updates the current tween toward the target frame.
//
// This method advances the interpolation started by [Animation.TweenTo] using the given progress value.
//
//   - animation : The Lottie animation object.
//   - progress : The current progress of the interpolation (range: 0.0 to 1.0).
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the animation is not loaded.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If [Animation.TweenTo] has not been called.
//
// see [Animation.TweenTo]
//
// Note: Experimental API
func (animation Animation) TweenGo(progress float32) error {
	return tvg_lottie_animation_tween_go(animation, progress).error()
}

// Sets the quality level for Lottie effects.
//
// This function controls the rendering quality of effects like blur, shadows, etc. Lower values prioritize performance while higher values prioritize quality.
//
//   - animation : The Lottie animation object.
//   - value : The quality level (0-100). 0 represents lowest quality/best performance,                  100 represents highest quality/lowest performance, the initial value is 50.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the animation is not loaded.
//
// Note: This option is used as a hint; its behavior heavily depends on the render backend.
//
// Since: 1.0
func (animation Animation) SetQuality(value uint) error {
	return tvg_lottie_animation_set_quality(animation, uint8(value)).error()
}

// Checks whether Lottie expressions are supported.
//
// Return: true if expression support is enabled in this build, otherwise false.
//
// Note: Experimental API
func (animation Animation) ExpressionsSupported() bool {
	return tvg_lottie_animation_expressions_supported(animation)
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

// Sets the audio resolver callback for Lottie audio layers.
//
// The resolver is invoked whenever the playback state of an audio layer changes. It allows applications to synchronize audio playback with the animation timeline.
//
//   - animation : A Lottie animation object.
//   - resolver : A user-defined callback that receives audio playback state updates.
//   - data : User data passed to resolver.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : The animation has not been loaded.
//
// Note: To disable audio notifications, pass nullptr as resolver.
//
// Note: Experimental API.
//
// see Tvg_Audio_Resolver
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

	return tvg_lottie_animation_set_audio_resolver(animation, puregoResolver, 0).error()
}

// Sets the master playback volume of the Lottie animation.
//
// The specified volume is applied to all media assets in the animation. Each media asset preserves its own volume, and the effective playback volume is calculated by multiplying the master volume by the media asset's volume.
// A value of 1.0 preserves the original volume. A value of 0.0 mutes all media. Values greater than 1.0 amplify the playback volume.
// The recommended range is [0.0, 2.0].
//
//   - animation : A Lottie animation object.
//   - volume : Master playback volume.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the animation is not loaded.
//
// Note: Experimental API
//
// see [Animation.GetVolume]
func (animation Animation) SetVolume(volume float32) error {
	return tvg_lottie_animation_set_volume(animation, volume).error()
}

// Retrieves the master playback volume of the Lottie animation.
//
//   - animation : A Lottie animation object.
//
// Return: The current master playback volume, otherwise 0 if the animation is invalid.
//
// Note: Experimental API
//
// see [Animation.SetVolume]
func (animation Animation) GetVolume() float32 {
	return tvg_lottie_animation_get_volume(animation)
}
