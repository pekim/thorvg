package thorvg

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

type Picture struct {
	paintCommon
}

type PictureAssetResolver func(paint Paint, src string) bool

// Creates a new Picture object.
//
// This function allocates and returns a new Picture instance. To properly destroy the Picture object, use [Paint.Rel]
//
// Return: The newly created Picture object.
//
// see [Paint.Rel]
func PictureNew() Picture {
	return Picture{
		paintCommon: paintCommon{
			paint_: tvg_picture_new(),
		},
	}
}

// Loads a picture data directly from a file.
//
// ThorVG efficiently caches the loaded data using the specified path as a key. This means that loading the same file again will not result in duplicate operations; instead, ThorVG will reuse the previously loaded picture data.
//
//   - picture : The picture object.
//   - path : The absolute path to the image file.
//
// Returns [RESULT_INVALID_ARGUMENT] : An empty path.
//
// Returns [RESULT_NOT_SUPPORTED] : A file with an unknown extension.
func (picture Picture) Load(path string) error {
	return tvg_picture_load(picture.paint_, path).error()
}

// Loads raw image data in a specific format from a memory block of the given size.
//
// ThorVG efficiently caches the loaded data, using the provided data address as a key when copy is set to false. This allows ThorVG to avoid redundant operations by reusing the previously loaded picture data for the same sharable data, rather than duplicating the load process.
//
//   - picture : The picture object.
//   - data : The memory block where the raw image data is stored.
//   - w : The width of the image in pixels.
//   - h : The height of the image in pixels.
//   - cs : Specifies how the 32-bit color values should be interpreted (read/write).
//   - copy : If true, the data is copied into the engine's local buffer. If false, the data is not copied.
//
// Returns [RESULT_INVALID_ARGUMENT] : No data are provided or the w or h value is zero or less.
//
// Note: TVG_COLORSPACE_XBGR8888 and TVG_COLORSPACE_XRGB8888 overwrite fully covered pixels without alpha blending.
//
// Since: 0.9
func (picture Picture) LoadRaw(data []byte, width uint, height uint, colorSpace ColorSpace) error {
	return tvg_picture_load_raw(picture.paint_, (*uint32)(unsafe.Pointer(&data[0])),
		uint32(width), uint32(height), colorSpace, true).error()
}

// Loads a picture data from a memory block of a given size.
//
// ThorVG efficiently caches the loaded data using the specified data address as a key when the copy has false. This means that loading the same data again will not result in duplicate operations for the sharable data. Instead, ThorVG will reuse the previously loaded picture data.
//
//   - picture : The picture object.
//   - data : A pointer to a memory location where the content of the picture file is stored. A null-terminated string is expected for non-binary data if copy is false
//   - size : The size in bytes of the memory occupied by the data.
//   - mimetype : Mimetype or extension of data such as "jpg", "jpeg", "svg", "svg+xml", "lot", "lottie+json", "png", etc. In case an empty string or an unknown type is provided, the loaders will be tried one by one.
//   - rpath : A resource directory path, if the data needs to access any external resources.
//   - copy : If true the data are copied into the engine local buffer, otherwise they are not.
//
// Returns [RESULT_INVALID_ARGUMENT] : In case a nullptr is passed as the argument or the size is zero or less.
//
// Returns [RESULT_NOT_SUPPORTED] : A file with an unknown extension.
//
// Warning:: It's the user responsibility to release the data memory if the copy is true.
func (picture Picture) LoadDataText(text string, mimetype string, rpath string) error {
	return tvg_picture_load_data(picture.paint_, unsafe.StringData(text), uint32(len(text)),
		mimetype, rpath, true).error()
}

// Loads a picture data from a memory block of a given size.
//
// ThorVG efficiently caches the loaded data using the specified data address as a key when the copy has false. This means that loading the same data again will not result in duplicate operations for the sharable data. Instead, ThorVG will reuse the previously loaded picture data.
//
//   - picture : The picture object.
//   - data : A pointer to a memory location where the content of the picture file is stored. A null-terminated string is expected for non-binary data if copy is false
//   - size : The size in bytes of the memory occupied by the data.
//   - mimetype : Mimetype or extension of data such as "jpg", "jpeg", "svg", "svg+xml", "lot", "lottie+json", "png", etc. In case an empty string or an unknown type is provided, the loaders will be tried one by one.
//   - rpath : A resource directory path, if the data needs to access any external resources.
//   - copy : If true the data are copied into the engine local buffer, otherwise they are not.
//
// Returns [RESULT_INVALID_ARGUMENT] : In case a nullptr is passed as the argument or the size is zero or less.
//
// Returns [RESULT_NOT_SUPPORTED] : A file with an unknown extension.
//
// Warning:: It's the user responsibility to release the data memory if the copy is true.
func (picture Picture) LoadData(data []byte, mimetype string, rpath string) error {
	return tvg_picture_load_data(picture.paint_, &data[0], uint32(len(data)), mimetype, rpath, true).error()
}

// Sets the asset resolver callback for handling external resources (e.g., images and fonts).
//
// This callback is invoked when an external asset reference (such as an image source or file path) is encountered in a Picture object. It allows the user to provide a custom mechanism for loading or substituting assets, such as loading from an external source or a virtual filesystem.
//
//   - resolver : A user-defined function that handles the resolution of asset paths.                     The function should return true if the asset was successfully resolved by the user, or false if it was not.
//   - data : A pointer to user-defined data that will be passed to the callback each time it is invoked.                 This can be used to maintain context or access external resources.
//
// Returns [RESULT_INVALID_ARGUMENT] : A nullptr passed as the picture argument.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the picture is already loaded.
//
// Note: This function must be called before [Picture.Load]       Setting the resolver after loading will have no effect on asset resolution for that asset.
//
// Note: If false is returned by resolver, ThorVG will attempt to resolve the resource using its internal resolution mechanism as a fallback.
//
// Note: To unset the resolver, pass nullptr as the resolver parameter.
//
// Note: Experimental API
//
// see Tvg_Picture_Asset_Resolver
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

// Resizes the picture content to the given width and height.
//
// The picture content is resized while keeping the default size aspect ratio. The scaling factor is established for each of dimensions and the smaller value is applied to both of them.
//
//   - picture : The picture object.
//   - w : A new width of the image in pixels.
//   - h : A new height of the image in pixels.
func (picture Picture) SetSize(width float32, height float32) error {
	return tvg_picture_set_size(picture.paint_, width, height).error()
}

// Gets the size of the loaded picture.
//
//   - picture : The picture object.
//   - w : A width of the image in pixels.
//   - h : A height of the image in pixels.
func (picture Picture) GetSize() (float32, float32, error) {
	var width float32
	var height float32
	result := tvg_picture_get_size(picture.paint_, &width, &height)
	return width, height, result.error()
}

// Sets the normalized origin point of the Picture object.
//
// This method defines the origin point of the Picture using normalized coordinates. Unlike a typical pivot point used only for transformations, this origin affects both the transformation behavior and the actual rendering position of the Picture.
// The specified origin becomes the reference point for positioning the Picture on the canvas. For example, setting the origin to (0.5f, 0.5f) moves the visual center of the picture to the position specified by Paint::translate().
// The coordinates are given in a normalized range relative to the picture's bounds: - (0.0f, 0.0f): top-left corner - (0.5f, 0.5f): center - (1.0f, 1.0f): bottom-right corner
//
//   - picture : The picture object.
//   - x : The normalized x-coordinate of the origin point (range: 0.0f to 1.0f).
//   - y : The normalized y-coordinate of the origin point (range: 0.0f to 1.0f).
//
// Note: This origin directly affects how the Picture is placed on the canvas when using       transformations such as translate(), rotate(), or scale().
//
// see [Paint.Translate]
//
// see [Paint.Rotate]
//
// see [Paint.Scale]
//
// see [Paint.SetTransform]
//
// see [Picture.GetOrigin]
//
// Since: 1.0
func (picture Picture) SetOrigin(x float32, y float32) error {
	return tvg_picture_set_origin(picture.paint_, x, y).error()
}

// Gets the normalized origin point of the Picture object.
//
// This method retrieves the current origin point of the Picture, expressed in normalized coordinates relative to the picture’s bounds.
//
//   - picture : The picture object.
//   - x : The normalized x-coordinate of the origin (range: 0.0f to 1.0f).
//   - y : The normalized y-coordinate of the origin (range: 0.0f to 1.0f).
//
// see [Picture.SetOrigin]
// Since: 1.0
func (picture Picture) GetOrigin() (float32, float32, error) {
	var x float32
	var y float32
	result := tvg_picture_get_origin(picture.paint_, &x, &y)
	return x, y, result.error()
}

// Retrieve a paint object from the Picture scene by its Unique ID.
//
// This function searches for a paint object within the Picture scene that matches the provided id.
//
//   - picture : The picture object.
//   - id : The Unique ID of the paint object.
//
// Return: The paint object that matches the given identifier, or nullptr if no matching paint object is found.
//
// Note: Setting [Picture.SetAccessible] to true enables more efficient access.
//
// see [Accessor.GenerateId]
// Since: 1.0
func (picture Picture) GetPaint(id uint) (Paint, bool) {
	paint := tvg_picture_get_paint(picture.paint_, uint32(id))
	return newPaint(paint)
}

// Sets the image filtering method for rendering this picture.
//
// Specifies how the image data should be filtered when it is scaled or transformed during rendering. This affects the visual quality and performance of the output.
//
//   - picture : A Tvg_Paint pointer to the picture object.
//   - method : The filtering method to apply. The initial value is TVG_FILTER_METHOD_BILINEAR.
//
// see Tvg_Filter_Method
// Since: 1.1
func (picture Picture) SetFilter(method FilterMethod) error {
	return tvg_picture_set_filter(picture.paint_, method).error()
}

// Enable or disable accessible mode for a Picture.
//
// When accessible mode is enabled, the Picture maintains an internal mapping of ID-accessible vector assets nodes (such as SVG), allowing efficient access to Paint objects and their associated identifier information via Accessor APIs.
// When disabled, no additional mapping is maintained and all nodes are treated as general traversal targets.
//
//   - picture : The target Picture object.
//   - accessible : Set to true to enable accessible mode, or false to disable it.
//
// see [Accessor.GenerateId]
//
// see [Accessor.GetName]
//
// see [Picture.GetPaint]
//
// Since: 1.1
func (picture Picture) SetAccessible(accessible bool) error {
	return tvg_picture_set_accessible(picture.paint_, accessible).error()
}

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
