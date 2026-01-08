package internal

import "unsafe"

/*
Canvas is a structure responsible for managing and drawing graphical elements.

It sets up the target buffer, which can be drawn on the screen. It stores the Paint objects (Shape, Scene, Picture).
*/
type Canvas struct {
	canvas     uintptr
	buffer     *byte // only used for a software canvas
	bufferSize int
}

/*
Buffer returns the buffer target of a software canvas.

If the canvas is not a software canvas, nil will be returned.
If the canvas is a software canvas but has never has a target set, nil will be returned.
*/
func (canvas Canvas) Buffer() []uint32 {
	if canvas.buffer == nil {
		return nil
	}

	return unsafe.Slice((*uint32)(unsafe.Pointer(canvas.buffer)), canvas.bufferSize/4)
}

/*
SwCanvasCreate creates a new SwCanvas object with optional rendering engine settings.

This method generates a software canvas instance that can be used for drawing vector graphics.
It accepts an optional parameter @p op to choose between different rendering engine behaviors.

@param[in] op The rendering engine option.

@return A new Tvg_Canvas object.

@see enum Tvg_Engine_Option
*/
func SwCanvasCreate(option EngineOption) Canvas {
	return Canvas{
		canvas: tvg_swcanvas_create(option),
	}
}

/*
*
@brief Sets the buffer used in the rasterization process and defines the used colorspace.

For optimisation reasons TVG does not allocate memory for the output buffer on its own.
The buffer of a desirable size should be allocated and owned by the caller.

@param[in] canvas The Tvg_Canvas object managing the @p buffer.
@param[in] buffer A pointer to the allocated memory block of the size @p stride x @p h.
@param[in] stride The stride of the raster image - in most cases same value as @p w.
@param[in] w The width of the raster image.
@param[in] h The height of the raster image.
@param[in] cs The colorspace value defining the way the 32-bits colors should be read/written.

@retval TVG_RESULT_INVALID_ARGUMENTS An invalid canvas or buffer pointer passed or one of the @p stride, @p w or @p h being zero.
@retval TVG_RESULT_INSUFFICIENT_CONDITION if the canvas is performing rendering. Please ensure the canvas is synced.
@retval TVG_RESULT_NOT_SUPPORTED The software engine is not supported.

@warning Do not access @p buffer during tvg_canvas_draw() - tvg_canvas_sync(). It should not be accessed while the engine is writing on it.

@see Tvg_Colorspace
*/
func (canvas *Canvas) SwCanvasSetTarget(stride uint, w uint, h uint, cs ColorSpace) Result {
	if canvas.buffer != nil {
		free(canvas.buffer)
	}
	canvas.bufferSize = int(4 * w * h)
	canvas.buffer = malloc(canvas.bufferSize)

	result := tvg_swcanvas_set_target(canvas.canvas, canvas.buffer, uint32(stride), uint32(w), uint32(h), cs)
	return result
}
