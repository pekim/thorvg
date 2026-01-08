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
func (canvas *Canvas) SwSetTarget(stride uint, width uint, height uint, cs ColorSpace) error {
	if canvas.buffer != nil {
		free(canvas.buffer)
	}
	canvas.bufferSize = int(4 * width * height)
	canvas.buffer = malloc(canvas.bufferSize)

	result := tvg_swcanvas_set_target(canvas.canvas, canvas.buffer, uint32(stride), uint32(width), uint32(height), cs)
	return result.error()
}

/*
GlCanvasCreate creates an OpenGL rasterizer Canvas object.

@return A new Tvg_Canvas object.

@since 1.0.0
*/
func GlCanvasCreate() Canvas {
	return Canvas{
		canvas: tvg_glcanvas_create(),
	}
}

/*
GlSetTarget sets the drawing target for rasterization.

This function specifies the drawing target where the rasterization will occur. It can target
a specific framebuffer object (FBO) or the main surface.

@param[in] display The platform-specific display handle (EGLDisplay for EGL). Set @c nullptr for other systems.
@param[in] surface The platform-specific surface handle (EGLSurface for EGL, HDC for WGL). Set @c nullptr for other systems.
@param[in] context The OpenGL context to be used for rendering on this canvas.
@param[in] id The GL target ID, usually indicating the FBO ID. A value of @c 0 specifies the main surface.
@param[in] w The width (in pixels) of the raster image.
@param[in] h The height (in pixels) of the raster image.
@param[in] cs Specifies how the pixel values should be interpreted. Currently, it only allows @c TVG_COLORSPACE_ABGR8888S as @c GL_RGBA8.

@note If @p display and @p surface are not provided, the ThorVG GL engine assumes that
the appropriate OpenGL context is already current and will not attempt to bind a new one.

@retval TVG_RESULT_INSUFFICIENT_CONDITION If the canvas is currently rendering.

Ensure that @ref tvg_canvas_sync() has been called before setting a new target.

@retval TVG_RESULT_NOT_SUPPORTED In case the gl engine is not supported.

@see tvg_canvas_sync()

@since 1.0
*/
func (canvas *Canvas) GlSetTarget(
	display unsafe.Pointer, surface unsafe.Pointer, context unsafe.Pointer, id int,
	width uint, height uint, colorSpace ColorSpace,
) error {
	result := tvg_glcanvas_set_target(canvas.canvas, display, surface, context, int32(id),
		uint32(width), uint32(height), colorSpace)
	return result.error()
}

/*
Destroy clears the canvas internal data, releases all paints stored by the canvas and destroys the canvas object itself.

@param[in] canvas The Tvg_Canvas object to be destroyed.

@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer to the Tvg_Canvas object is passed.
*/
func (canvas *Canvas) Destroy() error {
	if canvas.buffer != nil {
		free(canvas.buffer)
		canvas.buffer = nil
	}

	return tvg_canvas_destroy(canvas.canvas).error()
}
