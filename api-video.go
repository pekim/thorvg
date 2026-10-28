package thorvg

/*
Video is a structure representing a video playback controller object.

Experimental API
*/
type Video uintptr

// Creates a new video playback.
//
// Return: Tvg_Video A new video object.
//
// Note: Experimental API
func VideoNew() Video {
	return Video(tvg_video_new())
}

// Releases the given video object and its resources.
//
//   - video : The video object to release.
//
// Note: Experimental API
func (video Video) Del() error {
	return tvg_video_del(video).error()
}

// Retrieves the picture instance associated with this video object.
//
// This function provides access to the internal picture used by the video object for rendering video frames. After obtaining the picture, it can be added to a canvas enabling control over video frames with this Video object.
//
//   - video : The video object.
//
// Return: A picture instance representing the visual content of the video.
//
// Warning: The picture instance is owned by the video object and should not be deleted manually.
//
// Note: Experimental API
func (video Video) Picture() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: tvg_video_get_picture(video),
		},
	}
}

// Starts or resumes playback.
//
//   - video : The video object.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// Note: Experimental API
func (video Video) Play() error {
	return tvg_video_play(video).error()
}

// Pauses playback while keeping the current position.
//
//   - video : The video object.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not started playing.
//
// Note: Experimental API
func (video Video) Pause() error {
	return tvg_video_pause(video).error()
}

// Stops playback and rewinds to the start.
//
//   - video : The video object.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// Note: Experimental API
func (video Video) Stop() error {
	return tvg_video_stop(video).error()
}

// Sets the playback position.
//
//   - video : The video object.
//   - seconds : The target playback position in seconds.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// Returns [RESULT_INVALID_ARGUMENT] : If seconds is out of range.
//
// Note: Experimental API
func (video Video) Seek(seconds float32) error {
	return tvg_video_seek(video, seconds).error()
}

// Enables or disables repeated playback.
//
//   - video : The video object.
//   - on : true to repeat playback, false to play once.               The initial value is false.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// see [Video.GetLoop]
//
// Note: Experimental API
func (video Video) SetLoop(on bool) error {
	return tvg_video_set_loop(video, on).error()
}

// Retrieves whether repeated playback is enabled.
//
// Return: true if repeated playback is enabled, false otherwise.
//
// see [Video.SetLoop]
//
// Note: Experimental API
func (video Video) GetLoop() bool {
	return tvg_video_get_loop(video)
}

// Sets the audio volume level.
//
//   - video : The video object.
//   - volume : The audio volume level in the range [0.0, 1.0].                   The initial value is 1.0.
//
// Returns [RESULT_INVALID_ARGUMENT] : If volume is out of range.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// see [Video.GetVolume]
//
// Note: Experimental API
func (video Video) SetVolume(volume float32) error {
	return tvg_video_set_volume(video, volume).error()
}

// Retrieves the current audio volume level.
//
//   - video : The video object.
//
// Return: The current audio volume level in the range [0.0, 1.0].
//
// see [Video.SetVolume]
//
// Note: Experimental API
func (video Video) GetVolume() float32 {
	return tvg_video_get_volume(video)
}

// Mutes or unmutes the audio without changing the volume level.
//
//   - video : The video object.
//   - on : true to mute the audio, false to unmute.               The initial value is false.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the video is not loaded.
//
// see [Video.GetMuted]
//
// Note: Experimental API
func (video Video) SetMute(on bool) error {
	return tvg_video_set_mute(video, on).error()
}

// Retrieves whether the audio is currently muted.
//
//   - video : The video object.
//
// Return: true if the audio is muted, false otherwise.
//
// see [Video.SetMute]
//
// Note: Experimental API
func (video Video) GetMuted() bool {
	return tvg_video_get_muted(video)
}

// Retrieves the current playback position.
//
//   - video : The video object.
//
// Return: The current playback position in seconds.
//
// see [Video.Seek]
//
// see [Video.GetDuration]
//
// Note: Experimental API
func (video Video) GetTime() float32 {
	return tvg_video_get_time(video)
}

// Retrieves the total duration of the video.
//
//   - video : The video object.
//
// Return: The total duration in seconds.
//
// Note: Experimental API
func (video Video) GetDuration() float32 {
	return tvg_video_get_duration(video)
}
