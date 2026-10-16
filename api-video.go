package thorvg

/*
Video is a structure representing a video playback controller object.

Experimental API
*/
type Video uintptr

/*
VideoNew creates a new video playback.

	@return Tvg_Video A new video object.

	@note Experimental API
*/
func VideoNew() Video {
	return Video(tvg_video_new())
}

/*
Del releases the given video object and its resources.

	@param[in] video The video object to release.

	@note Experimental API
*/
// TVG_API Tvg_Result tvg_video_del(Tvg_Video video);
func (video Video) Del() error {
	return tvg_video_del(video).error()
}

/*
Picture retrieves the picture instance associated with this video object.

This function provides access to the internal picture used by the video object
for rendering video frames. After obtaining the picture, it can be added to a canvas
enabling control over video frames with this Video object.

	@param[in] video The video object.

	@return A picture instance representing the visual content of the video.

	@warning The picture instance is owned by the video object and should not be deleted manually.

	@note Experimental API
*/
func (video Video) Picture() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: tvg_video_get_picture(video),
		},
	}
}

/*
Play starts or resumes playback.

	@param[in] video The video object.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.

	@note Experimental API
*/
func (video Video) Play() Result {
	return tvg_video_play(video)
}

/*
Pause pauses playback while keeping the current position.

	@param[in] video The video object.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not started playing.

	@note Experimental API
*/
func (video Video) Pause() Result {
	return tvg_video_pause(video)
}

/*
Stop stops playback and rewinds to the start.

	@param[in] video The video object.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.

	@note Experimental API
*/
func (video Video) Stop() Result {
	return tvg_video_stop(video)
}

/*
Seek sets the playback position.

	@param[in] video The video object.
	@param[in] seconds The target playback position in seconds.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.
	@retval TVG_RESULT_INVALID_ARGUMENT If @p seconds is out of range.

	@note Experimental API
*/
func (video Video) Seek(seconds float32) Result {
	return tvg_video_seek(video, seconds)
}

/*
SetLoop enables or disables repeated playback.

	@param[in] video The video object.
	@param[in] on @c true to repeat playback, @c false to play once.
								The initial value is @c false.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.

	@see tvg_video_get_loop()

	@note Experimental API
*/
func (video Video) SetLoop(on bool) Result {
	return tvg_video_set_loop(video, on)
}

/*
GetLoop retrieves whether repeated playback is enabled.

	@return @c true if repeated playback is enabled, @c false otherwise.

	@see tvg_video_set_loop()

	@note Experimental API
*/
func (video Video) GetLoop() bool {
	return tvg_video_get_loop(video)
}

/*
SetVolume sets the audio volume level.

	@param[in] video The video object.
	@param[in] volume The audio volume level in the range [0.0, 1.0].
										The initial value is @c 1.0.

	@retval TVG_RESULT_INVALID_ARGUMENT If @p volume is out of range.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.

	@see tvg_video_get_volume()

	@note Experimental API
*/
func (video Video) SetVolume(volume float32) Result {
	return tvg_video_set_volume(video, volume)
}

/*
GetVolume retrieves the current audio volume level.

	@param[in] video The video object.

	@return The current audio volume level in the range [0.0, 1.0].

	@see tvg_video_set_volume()

	@note Experimental API
*/
func (video Video) GetVolume() float32 {
	return tvg_video_get_volume(video)
}

/*
SetMute mutes or unmutes the audio without changing the volume level.

	@param[in] video The video object.
	@param[in] on @c true to mute the audio, @c false to unmute.
								The initial value is @c false.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the video is not loaded.

	@see tvg_video_get_muted()

	@note Experimental API
*/
func (video Video) SetMute(on bool) Result {
	return tvg_video_set_mute(video, on)
}

/*
GetMuted retrieves whether the audio is currently muted.

	@param[in] video The video object.

	@return @c true if the audio is muted, @c false otherwise.

	@see tvg_video_set_mute()

	@note Experimental API
*/
func (video Video) GetMuted() bool {
	return tvg_video_get_muted(video)
}

/*
GetTime retrieves the current playback position.

	@param[in] video The video object.

	@return The current playback position in seconds.

	@see tvg_video_seek()
	@see tvg_video_get_duration()

	@note Experimental API
*/
func (video Video) GetTime() float32 {
	return tvg_video_get_time(video)
}

/*
GetDuration retrieves the total duration of the video.

	@param[in] video The video object.

	@return The total duration in seconds.

	@note Experimental API
*/
func (video Video) GetDuration() float32 {
	return tvg_video_get_duration(video)
}
