package thorvg

// import (
// 	"runtime"
// 	"unsafe"
// )

// /*
// Canvas is a structure responsible for managing and drawing graphical elements.

// It sets up the target buffer, which can be drawn on the screen. It stores the Paint objects (Shape, Scene, Picture).
// */
// type Canvas struct {
// 	canvas       uintptr
// 	buffer       []byte // only used for a software canvas
// 	bufferPinner runtime.Pinner
// }

// /*
// Buffer returns the buffer target of a software canvas.

// The buffer remains the canvas's target, so its contents should not be altered.

// If the canvas is not a software canvas, nil will be returned.
// If the canvas is a software canvas but has never has a target set, nil will be returned.
// */
// func (canvas Canvas) Buffer() []byte {
// 	return canvas.buffer
// }

// /*
// SwCanvasCreate creates a new SwCanvas object with optional rendering engine settings.

// This method generates a software canvas instance that can be used for drawing vector graphics.
// It accepts an optional parameter @p op to choose between different rendering engine behaviors.

// 	@param[in] op The rendering engine option.

// 	@return A new Tvg_Canvas object.

// 	@see enum Tvg_Engine_Option
// */
// func SwCanvasCreate(option EngineOption) Canvas {
// 	return Canvas{
// 		canvas: tvg_swcanvas_create(option),
// 	}
// }

// /*
// SwSetTarget Sets the buffer used in the rasterization process and defines the used colorspace.

// For optimisation reasons TVG does not allocate memory for the output buffer on its own.
// The buffer of a desirable size should be allocated and owned by the caller.

// 	@param[in] canvas The Tvg_Canvas object managing the @p buffer.
// 	@param[in] buffer A pointer to the allocated memory block of the size @p stride x @p h.
// 	@param[in] stride The stride of the raster image - in most cases same value as @p w.
// 	@param[in] w The width of the raster image.
// 	@param[in] h The height of the raster image.
// 	@param[in] cs The colorspace value defining the way the 32-bits colors should be read/written.

// 	@retval TVG_RESULT_INVALID_ARGUMENTS An invalid canvas or buffer pointer passed or one of the @p stride, @p w or @p h being zero.
// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION if the canvas is performing rendering. Please ensure the canvas is synced.
// 	@retval TVG_RESULT_NOT_SUPPORTED The software engine is not supported.

// 	  @warning Do not access @p buffer during tvg_canvas_draw() - tvg_canvas_sync(). It should not be accessed while the engine is writing on it.

// 	  @see Tvg_Colorspace
// */
// func (canvas *Canvas) SwSetTarget(stride uint, width uint, height uint, cs ColorSpace) error {
// 	canvas.bufferPinner.Unpin()
// 	canvas.buffer = make([]byte, 4*width*height)
// 	canvas.bufferPinner.Pin(&canvas.buffer)

// 	result := tvg_swcanvas_set_target(canvas.canvas, &canvas.buffer[0], uint32(stride), uint32(width), uint32(height), cs)
// 	return result.error()
// }

// /*
// GlCanvasCreate creates an OpenGL rasterizer Canvas object.

// 	@return A new Tvg_Canvas object.

// 	@since 1.0.0
// */
// func GlCanvasCreate() Canvas {
// 	return Canvas{
// 		canvas: tvg_glcanvas_create(),
// 	}
// }

// /*
// GlSetTarget sets the drawing target for rasterization.

// This function specifies the drawing target where the rasterization will occur. It can target
// a specific framebuffer object (FBO) or the main surface.

// 	@param[in] display The platform-specific display handle (EGLDisplay for EGL). Set @c nullptr for other systems.
// 	@param[in] surface The platform-specific surface handle (EGLSurface for EGL, HDC for WGL). Set @c nullptr for other systems.
// 	@param[in] context The OpenGL context to be used for rendering on this canvas.
// 	@param[in] id The GL target ID, usually indicating the FBO ID. A value of @c 0 specifies the main surface.
// 	@param[in] w The width (in pixels) of the raster image.
// 	@param[in] h The height (in pixels) of the raster image.
// 	@param[in] cs Specifies how the pixel values should be interpreted. Currently, it only allows @c TVG_COLORSPACE_ABGR8888S as @c GL_RGBA8.

// 	@note If @p display and @p surface are not provided, the ThorVG GL engine assumes that the appropriate OpenGL context is already current and will not attempt to bind a new one.

// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the canvas is currently rendering.
// 	Ensure that @ref tvg_canvas_sync() has been called before setting a new target.

// 	@retval TVG_RESULT_NOT_SUPPORTED In case the gl engine is not supported.

// 	@see tvg_canvas_sync()

// 	@since 1.0
// */
// func (canvas *Canvas) GlSetTarget(
// 	display unsafe.Pointer, surface unsafe.Pointer, context unsafe.Pointer, id int,
// 	width uint, height uint, colorSpace ColorSpace,
// ) error {
// 	result := tvg_glcanvas_set_target(canvas.canvas, display, surface, context, int32(id),
// 		uint32(width), uint32(height), colorSpace)
// 	return result.error()
// }

// /*
// WgCanvasCreate creates a WebGPU rasterizer Canvas object.

// 	@return A new Tvg_Canvas object.

// 	@since 1.0.0
// */
// func WgCanvasCreate() Canvas {
// 	return Canvas{
// 		canvas: tvg_wgcanvas_create(),
// 	}
// }

// /*
// WgSetTarget sets the drawing target for the rasterization.

// 	@param[in] device WGPUDevice, a desired handle for the wgpu device. If it is @c nullptr, ThorVG will assign an appropriate device internally.
// 	@param[in] instance WGPUInstance, context for all other wgpu objects.
// 	@param[in] target Either WGPUSurface or WGPUTexture, serving as handles to a presentable surface or texture.
// 	@param[in] w The width of the target.
// 	@param[in] h The height of the target.
// 	@param[in] cs Specifies how the pixel values should be interpreted. Currently, it only allows @c TVG_COLORSPACE_ABGR8888S as @c WGPUTextureFormat_RGBA8Unorm.
// 	@param[in] type @c 0: surface, @c 1: texture are used as pesentable target.

// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION if the canvas is performing rendering. Please ensure the canvas is synced.
// 	@retval TVG_RESULT_NOT_SUPPORTED In case the wg engine is not supported.

// 	  @since 1.0
// */
// func (canvas *Canvas) WgSetTarget(
// 	device unsafe.Pointer, instance unsafe.Pointer, target unsafe.Pointer,
// 	width uint32, height uint32, colorSpace ColorSpace, typ int,
// ) error {
// 	result := tvg_wgcanvas_set_target(canvas.canvas, device, instance, target,
// 		uint32(width), uint32(height), colorSpace, int32(typ))
// 	return result.error()
// }

// // TVG_API Tvg_Result tvg_wgcanvas_set_target(
// // Tvg_Canvas canvas,
// //  void* device, void* instance, void* target,
// //  uint32_t w, uint32_t h, Tvg_Colorspace cs, int type);

// /*
// Destroy clears the canvas internal data, releases all paints stored by the canvas and destroys the canvas object itself.

// 	@param[in] canvas The Tvg_Canvas object to be destroyed.

// 	@retval TVG_RESULT_INVALID_ARGUMENT An invalid pointer to the Tvg_Canvas object is passed.
// */
// func (canvas *Canvas) Destroy() error {
// 	canvas.bufferPinner.Unpin()
// 	canvas.buffer = nil
// 	return tvg_canvas_destroy(canvas.canvas).error()
// }

// /*
// Add inserts a drawing element into the canvas using a Tvg_Paint object.

// 	@param[in] canvas The Tvg_Canvas object managing the @p paint.
// 	@param[in] paint The Tvg_Paint object to be drawn.

// Only the paints added into the canvas will be drawing targets.
// They are retained by the canvas until you call tvg_canvas_remove()

// 	@return Tvg_Result return values:
// 	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument.
// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION An internal error.

// 	@note The rendering order of the paints is the same as the order as they were added. Consider sorting the paints before adding them if you intend to use layering.
// 	@see tvg_canvas_insert()
// 	@see tvg_canvas_remove()
// */
// func (canvas Canvas) Add(paint Paint) error {
// 	return tvg_canvas_add(canvas.canvas, paint.paint()).error()
// }

// /*
// Insert adds a paint object to the root scene.

// This function appends a paint object to root scene of the canvas. If the optional @p at
// is provided, the new paint object will be inserted immediately before the specified
// paint object in the root scene. If @p at is @c nullptr, the paint object will be added
// to the end of the root scene.

// 	@param[in] canvas The Tvg_Canvas object managing the @p paint.
// 	@param[in] target A pointer to the Paint object to be added into the root scene.
// 										This parameter must not be @c nullptr.
// 	@param[in] at A pointer to an existing Paint object in the root scene before which
// 								the new paint object will be added. If @c nullptr, the new
// 								paint object is added to the end of the root scene. The default is @c nullptr.

// 	@note The ownership of the @p paint object is transferred to the canvas upon addition.
// 	@note The rendering order of the paints is the same as the order as they were added. Consider sorting the paints before adding them if you intend to use layering.

// 	@see tvg_canvas_add()
// 	@see tvg_canvas_remove()
// 	@see tvg_canvas_remove()
// 	@since 1.0
// */
// func (canvas Canvas) Insert(target Paint, paint Paint) error {
// 	var cTarget uintptr
// 	if target != nil {
// 		cTarget = target.paint()
// 	}
// 	return tvg_canvas_insert(canvas.canvas, cTarget, paint.paint()).error()
// }

// /*
// Remove removes a paint object from the root scene.

// This function removes a specified paint object from the root scene. If no paint
// object is specified (i.e., the default @c nullptr is used), the function
// performs to clear all paints from the scene.

// 	@param[in] canvas A Tvg_Canvas object to remove the @p paint.
// 	@param[in] paint A pointer to the Paint object to be removed from the root scene.
// 									If @c nullptr, remove all the paints from the root scene.

// 	@see tvg_canvas_add()
// 	@see tvg_canvas_insert()
// 	@since 1.0
// */
// func (canvas Canvas) Remove(paint Paint) error {
// 	return tvg_canvas_remove(canvas.canvas, paint.paint()).error()
// }

// /*
// Update requests the canvas to update modified paint objects in preparation for rendering.

// This function triggers an internal update for all paint instances that have been modified
// since the last update. It ensures that the canvas state is ready for accurate rendering.

// 	  @param[in] canvas The Tvg_Canvas object to be updated.

// 	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Canvas pointer.
// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION The canvas is not properly prepared.
// 						This may occur if the canvas target has not been set or if the update is called during drawing.
// 						Call tvg_canvas_sync() before trying.

// 	@note Only paint objects that have been changed will be processed.
// 	@note If the canvas is configured with multiple threads, the update may be performed asynchronously.

// 	  @see tvg_canvas_sync()
// */
// func (canvas Canvas) Update(Paint) error {
// 	return tvg_canvas_update(canvas.canvas).error()
// }

// /*
// Draw requests the canvas to render the Paint objects.

// 	@param[in] canvas The Tvg_Canvas object containing elements to be drawn.
// 	@param[in] clear If @c true, clears the target buffer to zero before drawing.

// 	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Canvas pointer.
// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION The canvas is not properly prepared.
// 						This may occur if the canvas target has not been set or if the update is called during drawing.
// 						without calling tvg_canvas_sync() in between.

// 	  @note Clearing the buffer is unnecessary if the canvas will be fully covered

// 		with opaque content. Skipping the clear can improve performance.

// 	  @note Drawing may be performed asynchronously if the thread count is greater than zero.

// 		To ensure the drawing process is complete, call sync() afterwards.

// 	  @note If the canvas has not been updated prior to tvg_canvas_draw(), it may implicitly perform tvg_canvas_update()

// 	  @see tvg_canvas_sync()
// 	  @see tvg_canvas_update()
// */
// func (canvas Canvas) Draw(clear_ bool) error {
// 	return tvg_canvas_draw(canvas.canvas, clear_).error()
// }

// /*
// Sync guarantees that drawing task is finished.

// 	@param[in] canvas The Tvg_Canvas object containing elements which were drawn.

// The Canvas rendering can be performed asynchronously. To make sure that rendering is finished,
// the tvg_canvas_sync() must be called after the tvg_canvas_draw() regardless of threading.

// 	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Canvas pointer.

// 	@see tvg_canvas_draw()
// */
// func (canvas Canvas) Sync() error {
// 	return tvg_canvas_sync(canvas.canvas).error()
// }

// /*
// *
// SetViewport Sets the drawing region of the canvas.

// This function defines a rectangular area of the canvas to be used for drawing operations.
// The specified viewport clips rendering output to the boundaries of that rectangle.

// Please note that changing the viewport is only allowed at the beginning of the rendering sequence—that is, after calling tvg_canvas_sync().

// 	@param[in] canvas The Tvg_Canvas object containing elements which were drawn.
// 	@param[in] x The x-coordinate of the upper-left corner of the rectangle.
// 	@param[in] y The y-coordinate of the upper-left corner of the rectangle.
// 	@param[in] w The width of the rectangle.
// 	@param[in] h The height of the rectangle.

// 	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Canvas pointer.
// 	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the canvas is not in a synced state.

// 	@see tvg_canvas_sync()
// 	@see tvg_swcanvas_set_target()
// 	@see tvg_glcanvas_set_target()
// 	@see tvg_wgcanvas_set_target()

// 	  @warning Changing the viewport is not allowed after calling tvg_canvas_add(),

// 		tvg_canvas_remove(), tvg_canvas_update(), or tvg_canvas_draw().

// 	  @note When the target is reset, the viewport will also be reset to match the target size.
// 	  @since 0.15
// */
// func (canvas Canvas) SetViewport(x int, y int, width int, height int) error {
// 	return tvg_canvas_set_viewport(canvas.canvas, int32(x), int32(y), int32(width), int32(height)).error()
// }
