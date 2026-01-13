package thorvg

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

type Picture struct {
	paintCommon
}

type PictureAssetResolver func(paint Paint, src string) bool

/*
@brief Creates a new Picture object.

This function allocates and returns a new Picture instance.
To properly destroy the Picture object, use @ref tvg_paint_rel().

	@return A pointer to the newly created Picture object.

	@see tvg_paint_rel()
*/
func PictureNew() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: tvg_picture_new(),
		},
	}
}

/*
Load Loads a picture data directly from a file.

ThorVG efficiently caches the loaded data using the specified @p path as a key.
This means that loading the same file again will not result in duplicate operations;
instead, ThorVG will reuse the previously loaded picture data.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] path The absolute path to the image file.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer or an empty @p path.
	@retval TVG_RESULT_NOT_SUPPORTED A file with an unknown extension.
*/
func (picture Picture) Load(path string) error {
	return tvg_picture_load(picture.paint_, path).error()
}

// TVG_API Tvg_Result tvg_picture_load(Tvg_Paint picture, char* path);

/*
LoadRaw loads raw image data in a specific format from a memory block of the given size.

ThorVG efficiently caches the loaded data, using the provided @p data address as a key
when @p copy is set to @c false. This allows ThorVG to avoid redundant operations
by reusing the previously loaded picture data for the same sharable @p data,
rather than duplicating the load process.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] data A pointer to the memory block where the raw image data is stored.
	@param[in] w The width of the image in pixels.
	@param[in] h The height of the image in pixels.
	@param[in] cs Specifies how the 32-bit color values should be interpreted (read/write).
	@param[in] copy If @c true, the data is copied into the engine's local buffer. If @c false, the data is not copied.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer or no data are provided or the @p w or @p h value is zero or less.

	@since 0.9
*/
func (picture Picture) LoadRaw(data []byte, width uint, height uint, colorSpace ColorSpace) error {
	return tvg_picture_load_raw(picture.paint_, (*uint32)(unsafe.Pointer(&data[0])),
		uint32(width), uint32(height), colorSpace, true).error()
}

/*
LoadDataText loads a picture data from a string.
The string should be a text based image format, such as svg.

ThorVG efficiently caches the loaded data using the specified @p data address as a key
when the @p copy has @c false. This means that loading the same data again will not result in duplicate operations
for the sharable @p data. Instead, ThorVG will reuse the previously loaded picture data.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] data A pointer to a memory location where the content of the picture file is stored. A null-terminated string is expected for non-binary data if @p copy is @c false
	@param[in] size The size in bytes of the memory occupied by the @p data.
	@param[in] mimetype Mimetype or extension of data such as "jpg", "jpeg", "svg", "svg+xml", "lot", "lottie+json", "png", etc. In case an empty string or an unknown type is provided, the loaders will be tried one by one.
	@param[in] rpath A resource directory path, if the @p data needs to access any external resources.
	@param[in] copy If @c true the data are copied into the engine local buffer, otherwise they are not.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument or the @p size is zero or less.
	@retval TVG_RESULT_NOT_SUPPORTED A file with an unknown extension.

	@warning: It's the user responsibility to release the @p data memory if the @p copy is @c true.
*/
func (picture Picture) LoadDataText(text string, mimetype string, rpath string) error {
	return tvg_picture_load_data(picture.paint_, unsafe.StringData(text), uint32(len(text)),
		mimetype, rpath, true).error()
}

/*
LoadData loads a picture data from a memory block of a given size.

ThorVG efficiently caches the loaded data using the specified @p data address as a key
when the @p copy has @c false. This means that loading the same data again will not result in duplicate operations
for the sharable @p data. Instead, ThorVG will reuse the previously loaded picture data.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] data A pointer to a memory location where the content of the picture file is stored. A null-terminated string is expected for non-binary data if @p copy is @c false
	@param[in] size The size in bytes of the memory occupied by the @p data.
	@param[in] mimetype Mimetype or extension of data such as "jpg", "jpeg", "svg", "svg+xml", "lot", "lottie+json", "png", etc. In case an empty string or an unknown type is provided, the loaders will be tried one by one.
	@param[in] rpath A resource directory path, if the @p data needs to access any external resources.
	@param[in] copy If @c true the data are copied into the engine local buffer, otherwise they are not.

	@retval TVG_RESULT_INVALID_ARGUMENT In case a @c nullptr is passed as the argument or the @p size is zero or less.
	@retval TVG_RESULT_NOT_SUPPORTED A file with an unknown extension.

	@warning: It's the user responsibility to release the @p data memory if the @p copy is @c true.
*/
func (picture Picture) LoadData(data []byte, mimetype string, rpath string) error {
	return tvg_picture_load_data(picture.paint_, &data[0], uint32(len(data)), mimetype, rpath, true).error()
}

/*
SetAssetResolver sets the asset resolver callback for handling external resources (e.g., images and fonts).

This callback is invoked when an external asset reference (such as an image source or file path)
is encountered in a Picture object. It allows the user to provide a custom mechanism for loading
or substituting assets, such as loading from an external source or a virtual filesystem.

	@param[in] resolver A user-defined function that handles the resolution of asset paths.
											The function should return @c true if the asset was successfully resolved by the user, or @c false if it was not.
	@param[in] data A pointer to user-defined data that will be passed to the callback each time it is invoked.
									This can be used to maintain context or access external resources.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p picture argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the @p picture is already loaded.

	@note This function must be called before @ref tvg_picture_load()
				Setting the resolver after loading will have no effect on asset resolution for that asset.
	@note If @c false is returned by @p resolver, ThorVG will attempt to resolve the resource using its internal resolution mechanism as a fallback.
	@note To unset the resolver, pass @c nullptr as the @p resolver parameter.
	@note Experimental API

	@see Tvg_Picture_Asset_Resolver
*/
func (picture Picture) SetAssetResolver(resolver PictureAssetResolver) error {
	puregoResolver := purego.NewCallback(func(cPaint uintptr, src *byte, _data uintptr) bool {
		paint, ok := newPaint(cPaint)
		if !ok {
			// This should never occur, so provide a non-specific Paint type.
			paint = paintCommon{paint_: cPaint}
		}

		return resolver(paint, goString(src))
	})

	return tvg_picture_set_asset_resolver(picture.paint_, puregoResolver, 0).error()
}

// TVG_API Tvg_Result tvg_picture_set_asset_resolver(Tvg_Paint picture, Tvg_Picture_Asset_Resolver resolver, void* data);

/*
SetSize resizes the picture content to the given width and height.

The picture content is resized while keeping the default size aspect ratio.
The scaling factor is established for each of dimensions and the smaller value is applied to both of them.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] w A new width of the image in pixels.
	@param[in] h A new height of the image in pixels.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (picture Picture) SetSize(width float32, height float32) error {
	return tvg_picture_set_size(picture.paint_, width, height).error()
}

/*
GetSize gets the size of the loaded picture.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[out] w A width of the image in pixels.
	@param[out] h A height of the image in pixels.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.
*/
func (picture Picture) GetSize() (float32, float32, error) {
	var width float32
	var height float32
	result := tvg_picture_get_size(picture.paint_, &width, &height)
	return width, height, result.error()
}

/*
SetOrigin sets the normalized origin point of the Picture object.

This method defines the origin point of the Picture using normalized coordinates.
Unlike a typical pivot point used only for transformations, this origin affects both
the transformation behavior and the actual rendering position of the Picture.

The specified origin becomes the reference point for positioning the Picture on the canvas.
For example, setting the origin to (0.5f, 0.5f) moves the visual center of the picture
to the position specified by Paint::translate().

The coordinates are given in a normalized range relative to the picture's bounds:
- (0.0f, 0.0f): top-left corner
- (0.5f, 0.5f): center
- (1.0f, 1.0f): bottom-right corner

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] x The normalized x-coordinate of the origin point (range: 0.0f to 1.0f).
	@param[in] y The normalized y-coordinate of the origin point (range: 0.0f to 1.0f).

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@note This origin directly affects how the Picture is placed on the canvas when using
				transformations such as translate(), rotate(), or scale().

	@see tvg_paint_translate()
	@see tvg_paint_rotate()
	@see tvg_paint_scale()
	@see tvg_paint_set_transform()
	@see tvg_picture_get_origin()

	@since 1.0
*/
func (picture Picture) SetOrigin(x float32, y float32) error {
	return tvg_picture_set_origin(picture.paint_, x, y).error()
}

/*
GetOrigin gets the normalized origin point of the Picture object.

This method retrieves the current origin point of the Picture, expressed
in normalized coordinates relative to the picture’s bounds.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[out] x The normalized x-coordinate of the origin (range: 0.0f to 1.0f).
	@param[out] y The normalized y-coordinate of the origin (range: 0.0f to 1.0f).

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid Tvg_Paint pointer.

	@see tvg_picture_set_origin()
	@since 1.0
*/
func (picture Picture) GetOrigin() (float32, float32, error) {
	var x float32
	var y float32
	result := tvg_picture_get_origin(picture.paint_, &x, &y)
	return x, y, result.error()
}

/*
GetPaint retrieves a paint object from the Picture scene by its Unique ID.

This function searches for a paint object within the Picture scene that matches the provided @p id.

	@param[in] picture A Tvg_Paint pointer to the picture object.
	@param[in] id The Unique ID of the paint object.

	@return A pointer to the paint object that matches the given identifier, or @c nullptr if no matching paint object is found.

	@see tvg_accessor_generate_id()
	@since 1.0
*/
func (picture Picture) GetPaint(id uint) (Paint, bool) {
	paint := tvg_picture_get_paint(picture.paint_, uint32(id))
	return newPaint(paint)
}

// TVG_API Tvg_Paint tvg_picture_get_paint(Tvg_Paint picture, uint32_t id);

/*
Duplicate duplicates a Picture.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (picture Picture) Duplicate() Picture {
	return Picture{
		paintCommon: picture.duplicate(),
	}
}
