package thorvg

import "unsafe"

type Text struct {
	paintCommon
}

func (text Text) paint() uintptr {
	return text.paint_
}

/*
TextNew creates a new Text object.

This function allocates and returns a new Text instance.
To properly destroy the Text object, use @ref tvg_paint_rel().

	@return A pointer to the newly created Text object.

	@see tvg_paint_rel()

	@since 0.15
*/
func TextNew() Text {
	return Text{
		paintCommon: paintCommon{
			paint_: tvg_text_new(),
		},
	}
}

/*
SetFont sets the font family for the text.

This function specifies the name of the font to be used when rendering text.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] name The name of the font. This should match a font available through the canvas backend.
									If set to @c nullptr, ThorVG will attempt to select a fallback font available on the engine.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p paint argument.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION  The specified @p name cannot be found.

	@note This function only sets the font family name. Use @ref size() to define the font size.
	@note If the @p name is not specified, ThorVG will select an available fallback font.

	@see tvg_text_set_size()
	@see tvg_font_load()

	@since 1.0
*/
func (text Text) SetFont(name string) error {
	return tvg_text_set_font(text.paint(), name).error()
}

/*
SetSize sets the font size for the text.

This function sets the font size used during text rendering.
The size is specified in point units, and supports floating-point precision
for smooth scaling and animation effects.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] size The font size in points. Must be greater than 0.0.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p paint argument.
	@retval TVG_RESULT_INVALID_ARGUMENT if the @p size is less than or equal to 0.

	@note Use this function in combination with @ref font() to fully define text appearance.
	@note Fractional sizes (e.g., 12.5) are supported for sub-pixel rendering and animations.

	@see tvg_text_set_font()

	@since 1.0
*/
func (text Text) SetSize(size float32) error {
	return tvg_text_set_size(text.paint_, size).error()
}

/*
SetText assigns the given unicode text to be rendered.

This function sets the unicode text that will be displayed by the rendering system.
The text is set according to the specified UTF encoding method, which defaults to UTF-8.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] utf8 The multi-byte text encoded with utf8 string to be rendered.

	@since 1.0
*/
func (text Text) SetText(utf8 string) error {
	return tvg_text_set_text(text.paint_, utf8).error()
}

/*
Align sets text alignment or anchor per axis.

If layout width/height is set on an axis, align within the layout box.
Otherwise, treat it as an anchor within the text bounds which point of
the text box is pinned to the paint position.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] x Horizontal alignment/anchor in [0..1]: 0=left/start, 0.5=center, 1=right/end. (Default is 0)
	@param[in] y Vertical alignment/anchor in [0..1]: 0=top, 0.5=middle, 1=bottom. (Default is 0)

	@since 1.0

	@see tvg_text_layout()
*/
func (text Text) Align(x float32, y float32) error {
	return tvg_text_align(text.paint_, x, y).error()
}

/*
Layout sets the virtual layout box (constraints) for the text.

If width/height is set on an axis, that axis is constrained by a virtual layout box and
the text may wrap/align inside it. If width/height == 0, the axis is
unconstrained and @ref tvg_text_align() acts as an anchor on that axis.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] w Layout width in user space. Use 0 for no horizontal constraint. (Default is 0)
	@param[in] h Layout height in user space. Use 0 for no vertical constraint. (Default is 0)

	@note This defines constraints only; alignment/anchoring is controlled by @ref align().
	@since 1.0

	@see tvg_text_align()
	@see tvg_text_spacing()
*/
func (text Text) Layout(width float32, h float32) error {
	return tvg_text_layout(text.paint_, width, h).error()
}

/*
WrapMode sets the text wrapping mode for this text object.

This method controls how the text is laid out when it exceeds the available space.
The wrapping mode determines whether text is truncated, wrapped by character or word,
or adjusted automatically. An ellipsis mode is also available for truncation with "...".

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] mode The wrapping strategy to apply. Default is @c TVG_TEXT_WRAP_NONE.

	@see Tvg_Text_Wrap
	@since 1.0
*/
func (text Text) WrapMode(mode TextWrap) error {
	return tvg_text_wrap_mode(text.paint_, mode).error()
}

/*
Spacing sets the spacing scale factors for text layout.

This function adjusts the letter spacing (horizontal space between glyphs) and
line spacing (vertical space between lines of text) using scale factors.

Both values are relative to the font's default metrics:
- The letter spacing is applied as a scale factor to the glyph's advance width.
- The line spacing is applied as a scale factor to the glyph's advance height.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] letter The scale factor for letter spacing.
											Values > 1.0 increase spacing, values < 1.0 decrease it.
											Must be greater than or equal to 0.0. (default: 1.0)

	@param[in] line The scale factor for line spacing.
										Values > 1.0 increase line spacing, values < 1.0 decrease it.
										Must be greater than or equal to 0.0. (default: 1.0)

	  @since 1.0
*/
func (text Text) Spacing(letter float32, line float32) error {
	return tvg_text_spacing(text.paint_, letter, line).error()
}

/*
SetItalic applies an italic (slant) transformation to the text.

This function applies a shear transformation to simulate an italic (oblique) style
for the current text object. The shear factor determines the degree of slant
applied along the X-axis.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] shear The shear factor to apply. A value of 0.0 applies no slant, while values around 0.5 result in a strong slant.
									Must be in the range [0.0, 0.5]. Recommended value is 0.18.

	@note The @p shear factor will be clamped to the valid range if it exceeds the limits.
	@note This does not require the font itself to be italic.
				It visually simulates the effect by applying a transformation matrix.

	@warning Excessive slanting may cause visual distortion depending on the font and size.

	@see tvg_text_set_font()

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p paint argument.

	@since 1.0
*/
func (text Text) SetItalic(shear float32) error {
	return tvg_text_set_italic(text.paint_, shear).error()
}

/*
SetOutline sets an outline (stroke) around the text object.

This function adds an outline to the text with the specified width and RGB color.
The outline enhances the visibility of the text by rendering a stroke around its glyphs.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param width The width of the outline. Must be positive value. (The default is 0)
	@param r     Red component of the outline color (0–255).
	@param g     Green component of the outline color (0–255).
	@param b     Blue component of the outline color (0–255).

	@note To disable the outline, set @p width to 0.
	@see tvg_text_set_fill_color() to set the main text fill color.

	@since 1.0
*/
func (text Text) SetOutline(width float32, r uint8, g uint8, b uint8) error {
	return tvg_text_set_outline(text.paint_, width, r, g, b).error()
}

/*
SetColor sets the text solid color.

	@param[in] paint A Tvg_Paint pointer to the text object.
	@param[in] r The red color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] g The green color channel value in the range [0 ~ 255]. The default value is 0.
	@param[in] b The blue color channel value in the range [0 ~ 255]. The default value is 0.

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p paint argument.

	@note Either a solid color or a gradient fill is applied, depending on what was set as last.

	@see tvg_text_set_font()
	@see tvg_text_set_outline()

	@since 0.15
*/
func (text Text) SetColor(r uint8, g uint8, b uint8) error {
	return tvg_text_set_color(text.paint_, r, g, b).error()
}

/*
SetGradient sets the gradient fill for the text.

	@param[in] text A Tvg_Paint pointer to the text object.
	@param[in] grad The linear or radial gradient fill

	@retval TVG_RESULT_INVALID_ARGUMENT A @c nullptr passed as the @p paint argument.
	@retval TVG_RESULT_MEMORY_CORRUPTION An invalid Tvg_Gradient pointer.

	@note Either a solid color or a gradient fill is applied, depending on what was set as last.
	@see tvg_text_set_font()

	@since 0.15
*/
func (text Text) SetGradient(gradient Gradient) error {
	return tvg_text_set_gradient(text.paint_, gradient.gradient()).error()
}

/*
GetTextMetrics retrieves the layout metrics of the text object.

Fills the provided @ref Tvg_Text_Metrics structure with the font layout values of this text object,
such as ascent, descent, linegap, and line advance.

The returned values reflect the font size applied to the text object,
but do not include any transformations (e.g., scale, rotation, or translation).

	@param[in] text The text object.
	@param[out] metrics A pointer to a @ref Tvg_Text_Metrics structure to be filled with the resulting values.

	@return TVG_RESULT_INSUFFICIENT_CONDITION if no font or size has been set yet.

	@see Tvg_Text_Metrics
	@note Experimental API
*/
func (text Text) GetTextMetrics() (TextMetrics, error) {
	var metrics TextMetrics
	result := tvg_text_get_text_metrics(text.paint_, &metrics)
	return metrics, result.error()
}

/*
GetGlyphMetrics retrieves the layout metrics of a glyph in the text object.

Fills the provided @ref Tvg_Glyph_Metrics structure with the horizontal layout values
of the specified glyph, such as advance, left-side bearing, and bounding box.

The returned values reflect the font size applied to the text object,
but do not include any transformations (e.g., scale, rotation, or translation).

The input character must be a single UTF-8 encoded character.

	@param[in] text The text object.
	@param[in] ch A pointer to a UTF-8 encoded character.
	@param[out] metrics A pointer to a @ref Tvg_Glyph_Metrics structure to be filled with the resulting values.

	@return TVG_RESULT_INSUFFICIENT_CONDITION if no font or size has been set yet.
	@return TVG_RESULT_INVALID_ARGUMENT if the given character is invalid or not supported.

	@see Tvg_Glyph_Metrics
	@note Currently, ThorVG only supports horizontal text layout.
	@note Experimental API
*/
func (text Text) GetGlyphMetrics(ch rune) (GlyphMetrics, error) {
	var metrics GlyphMetrics
	result := tvg_text_get_glyph_metrics(text.paint_, (*byte)(unsafe.Pointer(&ch)), &metrics)
	return metrics, result.error()
}

/*
FontLoad loads a scalable font data from a file.

ThorVG efficiently caches the loaded data using the specified @p path as a key.
This means that loading the same file again will not result in duplicate operations;
instead, ThorVG will reuse the previously loaded font data.

	@param[in] path The path to the font file.

	@retval TVG_RESULT_INVALID_ARGUMENT An invalid @p path passed as an argument.
	@retval TVG_RESULT_NOT_SUPPORTED When trying to load a file with an unknown extension.

	@see tvg_font_unload()

	@since 0.15
*/
func FontLoad(path string) error {
	return tvg_font_load(path).error()
}

/*
FontLoadData loads a scalable font data from a memory block of a given size.

ThorVG efficiently caches the loaded font data using the specified @p name as a key.
This means that loading the same fonts again will not result in duplicate operations.
Instead, ThorVG will reuse the previously loaded font data.

	@param[in] name The name under which the font will be stored and accessible (e.x. in a @p tvg_text_set_font API).
	@param[in] data A pointer to a memory location where the content of the font data is stored.
	@param[in] size The size in bytes of the memory occupied by the @p data.
	@param[in] mimetype Mimetype or extension of font data. In case a @c nullptr or an empty "" value is provided the loader will be determined automatically.
	@param[in] copy If @c true the data are copied into the engine local buffer, otherwise they are not (default).

	@retval TVG_RESULT_INVALID_ARGUMENT If no name is provided or if @p size is zero while @p data points to a valid memory location.
	@retval TVG_RESULT_NOT_SUPPORTED When trying to load a file with an unknown extension.
	@retval TVG_RESULT_INSUFFICIENT_CONDITION When trying to unload the font data that has not been previously loaded.

	@warning: It's the user responsibility to release the @p data memory.

	@note To unload the font data loaded using this API, pass the proper @p name and @c nullptr as @p data.

	@since 0.15
*/
func FontLoadData(name string, data []byte, mimetype string) error {
	return tvg_font_load_data(name, &data[0], uint32(len(data)), mimetype, true).error()
}

/*
FontUnloadData unloads a font previously loaded with FontLoadData.

The name should match the name of the previously load font.
*/
func FontUnloadData(name string) error {
	return tvg_font_load_data(name, nil, 0, "", false).error()
}

/*
FontUnload unloads the specified scalable font data that was previously loaded.

This function is used to release resources associated with a font file that has been loaded into memory.

	@param[in] path The path to the loaded font file.

	@retval TVG_RESULT_INSUFFICIENT_CONDITION The loader is not initialized.

	@note If the font data is currently in use, it will not be immediately unloaded.
	@see tvg_font_load()

	@since 0.15
*/
func FontUnload(path string) error {
	return tvg_font_unload(path).error()
}

/*
Duplicate duplicates a Text.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (text Text) Duplicate() Text {
	return Text{
		paintCommon: text.duplicate(),
	}
}
