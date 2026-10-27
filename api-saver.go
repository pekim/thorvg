package thorvg

type Saver uintptr

// Creates a new Tvg_Saver object.
//
// Return: A new Tvg_Saver object.
func SaverNew() Saver {
	return Saver(tvg_saver_new())
}

// Exports the given paint data to the given path
//
// If the saver module supports any compression mechanism, it will optimize the data size. This might affect the encoding/decoding time in some cases. You can turn off the compression if you wish to optimize for speed.
//
//   - saver : The Tvg_Saver object connected with the saving task.
//   - paint : The paint to be saved with all its associated properties.
//   - path : A path to the file, in which the paint data is to be saved.
//   - quality : The encoded quality level. 0 is the minimum, 100 is the maximum value(recommended).
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : Currently saving other resources.
//
// Returns [RESULT_NOT_SUPPORTED] : Trying to save a file with an unknown extension or in an unsupported format.
//
// Returns [RESULT_UNKNOWN] : An empty paint is to be saved.
//
// Note: Saving can be asynchronous if the assigned thread number is greater than zero. To guarantee the saving is done, call tvg_saver_sync() afterwards.
//
// see [Saver.Sync]
func (saver Saver) SavePaint(paint Paint, path string, quality uint) error {
	return tvg_saver_save_paint(uintptr(saver), paint.paint(), path, uint32(quality)).error()
}

// Exports the given animation data to the given path
//
// If the saver module supports any compression mechanism, it will optimize the data size. This might affect the encoding/decoding time in some cases. You can turn off the compression if you wish to optimize for speed.
//
//   - saver : The Tvg_Saver object connected with the saving task.
//   - animation : The animation to be saved with all its associated properties.
//   - path : A path to the file, in which the animation data is to be saved.
//   - quality : The encoded quality level. 0 is the minimum, 100 is the maximum value(recommended).
//   - fps : The frames per second for the animation. If 0, the default fps is used.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : Currently saving other resources or animation has no frames.
//
// Returns [RESULT_NOT_SUPPORTED] : Trying to save a file with an unknown extension or in an unsupported format.
//
// Returns [RESULT_UNKNOWN] : Unknown if attempting to save an empty paint.
//
// Note: A higher frames per second (FPS) would result in a larger file size. It is recommended to use the default value.
//
// Note: Saving can be asynchronous if the assigned thread number is greater than zero. To guarantee the saving is done, call tvg_saver_sync() afterwards.
//
// see [Saver.Sync]
//
// Since: 1.0
func (saver Saver) SaveAnimation(animation Animation, path string, quality uint, fps uint) error {
	return tvg_saver_save_animation(uintptr(saver), uintptr(animation), path, uint32(quality), uint32(fps)).error()
}

// Guarantees that the saving task is finished.
//
// The behavior of the Saver module works on a sync/async basis, depending on the threading setting of the Initializer. Thus, if you wish to have a benefit of it, you must call tvg_saver_sync() after the tvg_saver_save_paint() in the proper delayed time. Otherwise, you can call tvg_saver_sync() immediately.
//
//   - saver : The Tvg_Saver object connected with the saving task.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : No saving task is running.
//
// Note: The asynchronous tasking is dependent on the Saver module implementation.
//
// see [Saver.SavePaint]
func (saver Saver) Sync() error {
	return tvg_saver_sync(uintptr(saver)).error()
}

// Deletes the given Tvg_Saver object.
//
//   - saver : The Tvg_Saver object to be deleted.
func (saver Saver) Del() error {
	return tvg_saver_del(uintptr(saver)).error()

}
