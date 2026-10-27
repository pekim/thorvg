package thorvg

import (
	"runtime"
	"unsafe"
)

/*
Canvas is a structure responsible for managing and drawing graphical elements.

It sets up the target buffer, which can be drawn on the screen. It stores the Paint objects (Shape, Scene, Picture).
*/
type Canvas struct {
	canvas       uintptr
	buffer       []byte // only used for a software canvas
	bufferPinner runtime.Pinner
}

/*
Region is a data structure representing an axis-aligned rectangular region.

The region is defined by its minimum and maximum x and y coordinates.

	@note Experimental API
*/
type Region struct {
	XMin, YMin int32
	XMax, YMax int32
}

/*
Buffer returns the buffer target of a software canvas.

The buffer remains the canvas's target, so its contents should not be altered.

If the canvas is not a software canvas, nil will be returned.
If the canvas is a software canvas but has never has a target set, nil will be returned.
*/
func (canvas Canvas) Buffer() []byte {
	return canvas.buffer
}

// Creates a new Software Canvas object with optional rendering engine settings.
//
// This method generates a software canvas instance that can be used for drawing vector graphics. It accepts an optional parameter op to choose between different rendering engine behaviors.
//
//   - op : The rendering engine option.
//
// Return: A new canvas object.
//
// see enum Tvg_Engine_Option
func SwCanvasCreate(option EngineOption) Canvas {
	return Canvas{
		canvas: tvg_swcanvas_create(option),
	}
}

// Sets the buffer used in the rasterization process and defines the used colorspace.
//
// For optimisation reasons TVG does not allocate memory for the output buffer on its own. The buffer of a desirable size should be allocated and owned by the caller.
//
//   - canvas : The canvas object managing the buffer.
//   - buffer : The allocated memory block of the size stride x h.
//   - stride : The stride of the raster image - in most cases same value as w.
//   - w : The width of the raster image.
//   - h : The height of the raster image.
//   - cs : The colorspace value defining the way the 32-bits colors should be read/written.
//
// Returns [RESULT_INVALID_ARGUMENT] : An invalid canvas or buffer pointer passed or one of the stride, w or h being zero.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : if the canvas is performing rendering. Please ensure the canvas is synced.
//
// Returns [RESULT_NOT_SUPPORTED] : The software engine is not supported.
//
// Warning: Do not access buffer during tvg_canvas_draw() - tvg_canvas_sync(). It should not be accessed while the engine is writing on it.
//
// Note: Currently, only TVG_COLORSPACE_ABGR8888, TVG_COLORSPACE_ARGB8888, TVG_COLORSPACE_ABGR8888S, and TVG_COLORSPACE_ARGB8888S are supported for cs.
//
// see Tvg_Colorspace
func (canvas *Canvas) SwSetTarget(stride uint, width uint, height uint, cs ColorSpace) error {
	canvas.bufferPinner.Unpin()
	canvas.buffer = make([]byte, 4*width*height)
	canvas.bufferPinner.Pin(&canvas.buffer)

	result := tvg_swcanvas_set_target(canvas.canvas, &canvas.buffer[0], uint32(stride), uint32(width), uint32(height), cs)
	return result.error()
}

// Creates a new OpenGL/ES Canvas object with optional rendering engine settings.
//
// This method generates a OpenGL/ES canvas instance that can be used for drawing vector graphics. It accepts an optional parameter op to choose between different rendering engine behaviors.
//
//   - op : The rendering engine option.
//
// Return: A new canvas object.
//
// Note: Currently, it does not support TVG_ENGINE_OPTION_SMART_RENDER. The request will be ignored.
//
// see enum Tvg_Engine_Option
//
// Since: 1.0
func GlCanvasCreate(option EngineOption) Canvas {
	return Canvas{
		canvas: tvg_glcanvas_create(option),
	}
}

// Sets the drawing target for rasterization.
//
// This function specifies the drawing target where the rasterization will occur. It can target a specific framebuffer object (FBO) or the main surface.
//
//   - display : The platform-specific display handle (EGLDisplay for EGL). Set nullptr for other systems.
//   - surface : The platform-specific surface handle (EGLSurface for EGL, HDC for WGL). Set nullptr for other systems.
//   - context : The OpenGL context to be used for rendering on this canvas.
//   - id : The GL target ID, usually indicating the FBO ID. A value of 0 specifies the main surface.
//   - w : The width (in pixels) of the raster image.
//   - h : The height (in pixels) of the raster image.
//   - cs : Specifies how the pixel values should be interpreted. Currently, it only allows TVG_COLORSPACE_ABGR8888S as GL_RGBA8.
//
// Note: If display and surface are not provided, the ThorVG GL engine assumes that       the appropriate OpenGL context is already current and will not attempt to bind a new one.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the canvas is currently rendering.         Ensure that [Canvas.Sync] has been called before setting a new target.
//
// Returns [RESULT_NOT_SUPPORTED] : In case the gl engine is not supported.
//
// see [Canvas.Sync]
//
// Since: 1.0
func (canvas *Canvas) GlSetTarget(
	display unsafe.Pointer, surface unsafe.Pointer, context unsafe.Pointer, id int,
	width uint, height uint, colorSpace ColorSpace,
) error {
	result := tvg_glcanvas_set_target(canvas.canvas, display, surface, context, int32(id),
		uint32(width), uint32(height), colorSpace)
	return result.error()
}

// Clears the canvas internal data, releases all paints stored by the canvas and destroys the canvas object itself.
//
//   - canvas : The canvas object to be destroyed.
func (canvas *Canvas) Destroy() error {
	canvas.bufferPinner.Unpin()
	canvas.buffer = nil
	return tvg_canvas_destroy(canvas.canvas).error()
}

// Adds a paint object to the canvas for rendering.
//
// Adds the specified paint into the canvas root scene. Only paints added to the canvas are considered rendering targets. The canvas retains the paint object until it is explicitly removed via tvg_canvas_remove().
//
//   - canvas : A handle to the canvas that will manage the paint object.
//   - paint :  A handle to the paint object to be rendered.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the canvas is not in a valid state to accept new paints.
//
// Note: Ownership of the paint object is transferred to the canvas upon       successful addition. To retain ownership, call [Paint.Ref]       before adding it to the canvas.
//
// Note: The rendering order of paint objects follows the order in which they are       added to the canvas. If layering is required, ensure paints are added in       the desired order.
//
// see [Canvas.Insert]
//
// see [Canvas.Remove]
//
// Since: 1.0
func (canvas Canvas) Add(paint Paint) error {
	return tvg_canvas_add(canvas.canvas, paint.paint()).error()
}

// Inserts a paint object into the canvas root scene.
//
// Inserts a paint object into the root scene of the specified canvas. If the at parameter is provided, the paint object is inserted immediately before the specified paint in the root scene. If at is nullptr, the paint object is appended to the end of the root scene.
//
//   - canvas : A handle to the canvas that will manage the paint object.                   This parameter must be a valid canvas handle.
//   - target : A handle to the paint object to be inserted into the root scene.                   This parameter must not be nullptr.
//   - at :     A handle to an existing paint object in the root scene before                   which target will be inserted. If nullptr, target is                   appended to the end of the root scene.
//
// Note: Ownership of the target object is transferred to the canvas upon       successful addition. To retain ownership, call [Paint.Ref]       before adding it to the canvas.
//
// Note: The rendering order of paint objects follows their order in the root       scene. If layering is required, ensure paints are inserted in the       desired order.
//
// see [Canvas.Add]
//
// see [Canvas.Remove]
//
// Since: 1.0
func (canvas Canvas) Insert(target Paint, paint Paint) error {
	var cTarget uintptr
	if target != nil {
		cTarget = target.paint()
	}
	return tvg_canvas_insert(canvas.canvas, cTarget, paint.paint()).error()
}

// Removes a paint object from the root scene.
//
// This function removes a specified paint object from the root scene. If no paint object is specified (i.e., the default nullptr is used), the function performs to clear all paints from the scene.
//
//   - canvas : A canvas object to remove the paint.
//   - paint : The paint object to be removed from the root scene.                  If nullptr, remove all the paints from the root scene.
//
// see [Canvas.Add]
//
// see [Canvas.Insert]
// Since: 1.0
func (canvas Canvas) Remove(paint Paint) error {
	return tvg_canvas_remove(canvas.canvas, paint.paint()).error()
}

// Requests the canvas to update modified paint objects in preparation for rendering.
//
// This function triggers an internal update for all paint instances that have been modified since the last update. It ensures that the canvas state is ready for accurate rendering.
//
//   - canvas : The canvas object to be updated.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : The canvas is not properly prepared.         This may occur if the canvas target has not been set or if the update is called during drawing.         Call tvg_canvas_sync() before trying.
//
// Note: Only paint objects that have been changed will be processed.
//
// Note: If the canvas is configured with multiple threads, the update may be performed asynchronously.
//
// see [Canvas.Sync]
func (canvas Canvas) Update(Paint) error {
	return tvg_canvas_update(canvas.canvas).error()
}

// Requests the canvas to render the Paint objects.
//
//   - canvas : The canvas object containing elements to be drawn.
//   - clear : If true, clears the target buffer to zero before drawing.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : The canvas is not properly prepared.         This may occur if the canvas target has not been set or if the update is called during drawing.         without calling tvg_canvas_sync() in between.
//
// Note: Clearing the buffer is unnecessary if the canvas will be fully covered       with opaque content. Skipping the clear can improve performance.
//
// Note: Drawing may be performed asynchronously if the thread count is greater than zero.       To ensure the drawing process is complete, call sync() afterwards.
//
// Note: If the canvas has not been updated prior to tvg_canvas_draw(), it may implicitly perform tvg_canvas_update()
//
// see [Canvas.Sync]
//
// see [Canvas.Update]
func (canvas Canvas) Draw(clear_ bool) error {
	return tvg_canvas_draw(canvas.canvas, clear_).error()
}

// Guarantees that drawing task is finished.
//
//   - canvas : The canvas object containing elements which were drawn.
//
// The Canvas rendering can be performed asynchronously. To make sure that rendering is finished, the tvg_canvas_sync() must be called after the tvg_canvas_draw() regardless of threading.
//
// see [Canvas.Draw]
func (canvas Canvas) Sync() error {
	return tvg_canvas_sync(canvas.canvas).error()
}

// Returns the number of dirty regions of the canvas.
//
//   - canvas : The canvas object.
//
// Return: The number of regions available for retrieval, or zero if canvas is NULL.
//
// Note: This function must be called after [Canvas.Draw] and before [Canvas.Sync]
//
// Note: If the active rendering backend does not support TVG_ENGINE_OPTION_SMART_RENDER, this function returns 1.
//
// Note: Experimental API
//
// see [Canvas.GetDirtyRegions]
func (canvas Canvas) GetDirtyRegionsCnt() uint32 {
	return tvg_canvas_get_dirty_regions_cnt(canvas.canvas)
}

// Retrieves the dirty regions of the canvas.
//
//   - canvas : The canvas object.
//   - regions : An array to receive the dirty regions.
//   - cnt : The number of Tvg_Region elements that the array can hold.
//
// Returns [RESULT_INVALID_ARGUMENT] : If canvas or regions is NULL, or cnt is zero.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If called outside the interval after tvg_canvas_draw() and before tvg_canvas_sync().
//
// Note: This function must be called after [Canvas.Draw] and before [Canvas.Sync]
//
// Note: If the active rendering backend does not support TVG_ENGINE_OPTION_SMART_RENDER,       a single region covering the entire render target is returned.
//
// Note: Experimental API
//
// see [Canvas.GetDirtyRegionsCnt]
func (canvas Canvas) GetDirtyRegions(cnt uint32) ([]Region, error) {
	regions := make([]Region, cnt)
	result := tvg_canvas_get_dirty_regions(canvas.canvas, &regions[0], cnt)
	return regions, result.error()
}

// Sets the drawing region of the canvas.
//
// This function defines a rectangular area of the canvas to be used for drawing operations. The specified viewport clips rendering output to the boundaries of that rectangle.
// Please note that changing the viewport is only allowed at the beginning of the rendering sequence—that is, after calling tvg_canvas_sync().
//
//   - canvas : The canvas object containing elements which were drawn.
//   - x : The x-coordinate of the upper-left corner of the rectangle.
//   - y : The y-coordinate of the upper-left corner of the rectangle.
//   - w : The width of the rectangle.
//   - h : The height of the rectangle.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the canvas is not in a synced state.
//
// see [Canvas.Sync]
//
// see [SwcanvasSetTarget]
//
// see [GlcanvasSetTarget]
//
// see [WgcanvasSetTarget]
//
// Warning: Changing the viewport is not allowed after calling tvg_canvas_add(),          tvg_canvas_remove(), tvg_canvas_update(), or tvg_canvas_draw().
//
// Note: When the target is reset, the viewport will also be reset to match the target size.
// Since: 0.15
func (canvas Canvas) SetViewport(x int, y int, width int, height int) error {
	return tvg_canvas_set_viewport(canvas.canvas, int32(x), int32(y), int32(width), int32(height)).error()
}
