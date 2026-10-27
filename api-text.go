package thorvg

import "unsafe"

type Text struct {
	paintCommon
}

func (text Text) paint() uintptr {
	return text.paint_
}

// Creates a new Text object.
//
// This function allocates and returns a new Text instance. To properly destroy the Text object, use [Paint.Rel]
//
// Return: The newly created Text object.
//
// see [Paint.Rel]
//
// Since: 0.15
func TextNew() Text {
	return Text{
		paintCommon: paintCommon{
			paint_: tvg_text_new(),
		},
	}
}

// Sets the font family for the text.
//
// This function specifies the name of the font to be used when rendering text.
//
//   - text : The text object.
//   - name : The name of the font. This should match a font available through the canvas backend.                 If set to nullptr, ThorVG will attempt to select a fallback font available on the engine.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] :  The specified name cannot be found.
//
// Note: This function only sets the font family name. Use size() to define the font size.
//
// Note: If the name is not specified, ThorVG will select an available fallback font.
//
// see [Text.SetSize]
//
// see [FontLoad]
//
// Since: 1.0
func (text Text) SetFont(name string) error {
	return tvg_text_set_font(text.paint(), name).error()
}

// Sets the font size for the text.
//
// This function sets the font size used during text rendering. The size is specified in point units, and supports floating-point precision for smooth scaling and animation effects.
//
//   - text : The text object.
//   - size : The font size in points. Must be greater than 0.0.
//
// Returns [RESULT_INVALID_ARGUMENT] : if the size is less than or equal to 0.
//
// Note: Use this function in combination with font() to fully define text appearance.
//
// Note: Fractional sizes (e.g., 12.5) are supported for sub-pixel rendering and animations.
//
// see [Text.SetFont]
//
// Since: 1.0
func (text Text) SetSize(size float32) error {
	return tvg_text_set_size(text.paint_, size).error()
}

// Assigns the given unicode text to be rendered.
//
// This function sets the unicode text that will be displayed by the rendering system. The text is set according to the specified UTF encoding method, which defaults to UTF-8.
//
//   - text : The text object.
//   - utf8 : The multi-byte text encoded with utf8 string to be rendered.
//
// see [Text.GetText]
//
// Since: 1.0
func (text Text) SetText(utf8 string) error {
	return tvg_text_set_text(text.paint_, utf8).error()
}

// Sets text alignment or anchor per axis.
//
// If layout width/height is set on an axis, align within the layout box. Otherwise, treat it as an anchor within the text bounds which point of the text box is pinned to the paint position.
//
//   - text : The text object.
//   - x : Horizontal alignment/anchor in [0..1]: 0=left/start, 0.5=center, 1=right/end. (The initial value is 0)
//   - y : Vertical alignment/anchor in [0..1]: 0=top, 0.5=middle, 1=bottom. (The initial value is 0)
//
// Since: 1.0
//
// see [Text.Layout]
func (text Text) Align(x float32, y float32) error {
	return tvg_text_align(text.paint_, x, y).error()
}

// Sets the virtual layout box (constraints) for the text.
//
// If width/height is set on an axis, that axis is constrained by a virtual layout box and the text may wrap/align inside it. If width/height == 0, the axis is unconstrained and [Text.Align] acts as an anchor on that axis.
//
//   - text : The text object.
//   - w : Layout width in user space. Use 0 for no horizontal constraint. (The initial value is 0)
//   - h : Layout height in user space. Use 0 for no vertical constraint. (The initial value is 0)
//
// Note: This defines constraints only; alignment/anchoring is controlled by align().
// Since: 1.0
//
// see [Text.Align]
//
// see [Text.Spacing]
func (text Text) Layout(width float32, h float32) error {
	return tvg_text_layout(text.paint_, width, h).error()
}

// Sets the text wrapping mode for this text object.
//
// This method controls how the text is laid out when it exceeds the available space. The wrapping mode determines whether text is truncated, wrapped by character or word, or adjusted automatically. An ellipsis mode is also available for truncation with "...".
//
//   - text : The text object.
//   - mode : The wrapping strategy to apply. The initial value is TVG_TEXT_WRAP_NONE.
//
// see Tvg_Text_Wrap
//
// see [Text.LineCount]
// Since: 1.0
func (text Text) WrapMode(mode TextWrap) error {
	return tvg_text_wrap_mode(text.paint_, mode).error()
}

// Set the spacing scale factors for text layout.
//
// This function adjusts the letter spacing (horizontal space between glyphs) and line spacing (vertical space between lines of text) using scale factors.
// Both values are relative to the font's default metrics: - The letter spacing is applied as a scale factor to the glyph's advance width. - The line spacing is applied as a scale factor to the glyph's advance height.
//
//   - text : The text object.
//
//   - letter : The scale factor for letter spacing.                   Values > 1.0 increase spacing, values < 1.0 decrease it.                   Must be greater than or equal to 0.0. The initial value is 1.0.
//
//   - line : The scale factor for line spacing.                 Values > 1.0 increase line spacing, values < 1.0 decrease it.                 Must be greater than or equal to 0.0. The initial value is 1.0.
//
// Since: 1.0
func (text Text) Spacing(letter float32, line float32) error {
	return tvg_text_spacing(text.paint_, letter, line).error()
}

// Apply an italic (slant) transformation to the text.
//
// This function applies a shear transformation to simulate an italic (oblique) style for the current text object. The shear factor determines the degree of slant applied along the X-axis.
//
//   - text : The text object.
//   - shear : The shear factor to apply. A value of 0.0 applies no slant, while values around 0.5 result in a strong slant.                  Must be in the range [0.0, 0.5]. Recommended value is 0.18.
//
// Note: The shear factor will be clamped to the valid range if it exceeds the limits.
//
// Note: This does not require the font itself to be italic.       It visually simulates the effect by applying a transformation matrix.
//
// Warning: Excessive slanting may cause visual distortion depending on the font and size.
//
// see [Text.SetFont]
//
// Since: 1.0
func (text Text) SetItalic(shear float32) error {
	return tvg_text_set_italic(text.paint_, shear).error()
}

// Sets an outline (stroke) around the text object.
//
// This function adds an outline to the text with the specified width and RGB color. The outline enhances the visibility of the text by rendering a stroke around its glyphs.
//
//   - text : The text object.
//   - width : The width of the outline. Must be positive value. (The initial value is 0)
//   - r :     Red component of the outline color (0–255).
//   - g :     Green component of the outline color (0–255).
//   - b :     Blue component of the outline color (0–255).
//
// Note: To disable the outline, set width to 0.
//
// see [Text.SetFillColor() To Set The Main Text Fill Color]
//
// Since: 1.0
func (text Text) SetOutline(width float32, r uint8, g uint8, b uint8) error {
	return tvg_text_set_outline(text.paint_, width, r, g, b).error()
}

// Sets the text solid color.
//
//   - paint : The text object.
//   - r : The red color channel value in the range [0 ~ 255]. The initial value is 0.
//   - g : The green color channel value in the range [0 ~ 255]. The initial value is 0.
//   - b : The blue color channel value in the range [0 ~ 255]. The initial value is 0.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Text.SetFont]
//
// see [Text.SetOutline]
//
// Since: 0.15
func (text Text) SetColor(r uint8, g uint8, b uint8) error {
	return tvg_text_set_color(text.paint_, r, g, b).error()
}

// Sets the gradient fill for the text.
//
//   - text : The text object.
//   - grad : The linear or radial gradient fill
//
// Returns [RESULT_MEMORY_CORRUPTION] : An invalid gradient object.
//
// Note: Either a solid color or a gradient fill is applied, depending on what was set as last.
//
// see [Text.SetFont]
//
// Since: 0.15
func (text Text) SetGradient(gradient Gradient) error {
	return tvg_text_set_gradient(text.paint_, gradient.gradient()).error()
}

// Retrieves the layout metrics of the text object.
//
// Fills the provided Tvg_Text_Metrics structure with the font layout values of this text object, such as ascent, descent, linegap, and line advance.
// The returned values reflect the font size applied to the text object, but do not include any transformations (e.g., scale, rotation, or translation).
//
//   - text : The text object.
//   - metrics : A pointer to a Tvg_Text_Metrics structure to be filled with the resulting values.
//
// Return: TVG_RESULT_INSUFFICIENT_CONDITION if no font or size has been set yet.
//
// see Tvg_Text_Metrics
//
// Note: Experimental API
func (text Text) GetTextMetrics() (TextMetrics, error) {
	var metrics TextMetrics
	result := tvg_text_get_text_metrics(text.paint_, &metrics)
	return metrics, result.error()
}

// Retrieves the layout metrics of a glyph in the text object.
//
// Fills the provided Tvg_Glyph_Metrics structure with the horizontal layout values of the specified glyph, such as advance, left-side bearing, and bounding box.
// The returned values reflect the font size applied to the text object, but do not include any transformations (e.g., scale, rotation, or translation).
// The input character must be a single UTF-8 encoded character.
//
//   - text : The text object.
//   - ch : A pointer to a UTF-8 encoded character.
//   - metrics : A pointer to a Tvg_Glyph_Metrics structure to be filled with the resulting values.
//   - next : An optional pointer that receives the position immediately                  following the processed UTF-8 character.
//
// Return: TVG_RESULT_INSUFFICIENT_CONDITION if no font or size has been set yet.
//
// Return: TVG_RESULT_INVALID_ARGUMENT if the given character is invalid or not supported.
//
// see Tvg_Glyph_Metrics
//
// Note: Currently, ThorVG only supports horizontal text layout.
//
// Note: Experimental API
func (text Text) GetGlyphMetrics(ch rune) (GlyphMetrics, error) {
	var metrics GlyphMetrics
	result := tvg_text_get_glyph_metrics(text.paint_, (*byte)(unsafe.Pointer(&ch)), &metrics, nil)
	return metrics, result.error()
}

// Loads a scalable font data from a file.
//
// ThorVG efficiently caches the loaded data using the specified path as a key. This means that loading the same file again will not result in duplicate operations; instead, ThorVG will reuse the previously loaded font data.
//
//   - path : The path to the font file.
//
// Returns [RESULT_INVALID_ARGUMENT] : An invalid path passed as an argument.
//
// Returns [RESULT_NOT_SUPPORTED] : When trying to load a file with an unknown extension.
//
// see [FontUnload]
//
// Since: 0.15
func FontLoad(path string) error {
	return tvg_font_load(path).error()
}

// Loads a scalable font data from a memory block of a given size.
//
// ThorVG efficiently caches the loaded font data using the specified name as a key. This means that loading the same fonts again will not result in duplicate operations. Instead, ThorVG will reuse the previously loaded font data.
//
//   - name : The name under which the font will be stored and accessible (e.x. in a tvg_text_set_font API).
//   - data : A pointer to a memory location where the content of the font data is stored.
//   - size : The size in bytes of the memory occupied by the data.
//   - mimetype : Mimetype or extension of font data. In case a nullptr or an empty "" value is provided the loader will be determined automatically.
//   - copy : If true the data are copied into the engine local buffer, otherwise they are not (default).
//
// Returns [RESULT_INVALID_ARGUMENT] : If no name is provided or if size is zero while data points to a valid memory location.
//
// Returns [RESULT_NOT_SUPPORTED] : When trying to load a file with an unknown extension.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : When trying to unload the font data that has not been previously loaded.
//
// Warning:: It's the user responsibility to release the data memory.
//
// Note: To unload the font data loaded using this API, pass the proper name and nullptr as data.
//
// Since: 0.15
func FontLoadData(name string, data []byte, mimetype string) error {
	return tvg_font_load_data(name, &data[0], uint32(len(data)), mimetype, true).error()
}

// Loads a scalable font data from a memory block of a given size.
//
// ThorVG efficiently caches the loaded font data using the specified name as a key. This means that loading the same fonts again will not result in duplicate operations. Instead, ThorVG will reuse the previously loaded font data.
//
//   - name : The name under which the font will be stored and accessible (e.x. in a tvg_text_set_font API).
//   - data : A pointer to a memory location where the content of the font data is stored.
//   - size : The size in bytes of the memory occupied by the data.
//   - mimetype : Mimetype or extension of font data. In case a nullptr or an empty "" value is provided the loader will be determined automatically.
//   - copy : If true the data are copied into the engine local buffer, otherwise they are not (default).
//
// Returns [RESULT_INVALID_ARGUMENT] : If no name is provided or if size is zero while data points to a valid memory location.
//
// Returns [RESULT_NOT_SUPPORTED] : When trying to load a file with an unknown extension.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : When trying to unload the font data that has not been previously loaded.
//
// Warning:: It's the user responsibility to release the data memory.
//
// Note: To unload the font data loaded using this API, pass the proper name and nullptr as data.
//
// Since: 0.15
func FontUnloadData(name string) error {
	return tvg_font_load_data(name, nil, 0, "", false).error()
}

// Unloads the specified scalable font data that was previously loaded.
//
// This function is used to release resources associated with a font file that has been loaded into memory.
//
//   - path : The path to the loaded font file.
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : The loader is not initialized.
//
// Note: If the font data is currently in use, it will not be immediately unloaded.
//
// see [FontLoad]
//
// Since: 0.15
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
