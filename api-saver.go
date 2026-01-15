package thorvg

type Saver uintptr

/*
SaverNew Creates a new Tvg_Saver object.

	@return A new Tvg_Saver object.
*/
func SaverNew() Saver {
	return Saver(tvg_saver_new())
}

/*
SavePaint exports the given @p paint data to the given @p path

If the saver module supports any compression mechanism, it will optimize the data size.
This might affect the encoding/decoding time in some cases. You can turn off the compression
if you wish to optimize for speed.

	@param[in] saver The Tvg_Saver object connected with the saving task.
	@param[in] paint The paint to be saved with all its associated properties.
	@param[in] path A path to the file, in which the paint data is to be saved.
	@param[in] quality The encoded quality level. @c 0 is the minimum, @c 100 is the maximum value(recommended).

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION Currently saving other resources.
	@retval TVG_RESULT_NOT_SUPPORTED Trying to save a file with an unknown extension or in an unsupported format.
	@retval TVG_RESULT_UNKNOWN An empty paint is to be saved.

	@note Saving can be asynchronous if the assigned thread number is greater than zero. To guarantee the saving is done, call tvg_saver_sync() afterwards.
	@see tvg_saver_sync()
*/
func (saver Saver) SavePaint(paint Paint, path string, quality uint) error {
	return tvg_saver_save_paint(uintptr(saver), paint.paint(), path, uint32(quality)).error()
}

/*
SaveAnimation Exports the given @p animation data to the given @p path

If the saver module supports any compression mechanism, it will optimize the data size.
This might affect the encoding/decoding time in some cases. You can turn off the compression
if you wish to optimize for speed.

	@param[in] saver The Tvg_Saver object connected with the saving task.
	@param[in] animation The animation to be saved with all its associated properties.
	@param[in] path A path to the file, in which the animation data is to be saved.
	@param[in] quality The encoded quality level. @c 0 is the minimum, @c 100 is the maximum value(recommended).
	@param[in] fps The frames per second for the animation. If @c 0, the default fps is used.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION Currently saving other resources or animation has no frames.
	@retval TVG_RESULT_NOT_SUPPORTED Trying to save a file with an unknown extension or in an unsupported format.
	@retval TVG_RESULT_UNKNOWN Unknown if attempting to save an empty paint.

	@note A higher frames per second (FPS) would result in a larger file size. It is recommended to use the default value.
	@note Saving can be asynchronous if the assigned thread number is greater than zero. To guarantee the saving is done, call tvg_saver_sync() afterwards.

	@see tvg_saver_sync()

	@since 1.0
*/
func (saver Saver) SaveAnimation(animation Animation, path string, quality uint, fps uint) error {
	return tvg_saver_save_animation(uintptr(saver), uintptr(animation), path, uint32(quality), uint32(fps)).error()
}

/*
Sync guarantees that the saving task is finished.

The behavior of the Saver module works on a sync/async basis, depending on the threading setting of the Initializer.
Thus, if you wish to have a benefit of it, you must call tvg_saver_sync() after the tvg_saver_save_paint() in the proper delayed time.
Otherwise, you can call tvg_saver_sync() immediately.

	@param[in] saver The Tvg_Saver object connected with the saving task.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION No saving task is running.

	@note The asynchronous tasking is dependent on the Saver module implementation.
	@see tvg_saver_save_paint()
*/
func (saver Saver) Sync() error {
	return tvg_saver_sync(uintptr(saver)).error()
}

/*
Del deletes the given Tvg_Saver object.

@param[in] saver The Tvg_Saver object to be deleted.

@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Saver pointer.
*/
func (saver Saver) Del() error {
	return tvg_saver_del(uintptr(saver)).error()

}
