package thorvg

// #include "thorvg_capi.h"
import "C"

type Result C.Tvg_Result

const (
	RESULT_SUCCESS                = Result(C.TVG_RESULT_SUCCESS)
	RESULT_INVALID_ARGUMENT       = Result(C.TVG_RESULT_INVALID_ARGUMENT)
	RESULT_INSUFFICIENT_CONDITION = Result(C.TVG_RESULT_INSUFFICIENT_CONDITION)
	RESULT_FAILED_ALLOCATION      = Result(C.TVG_RESULT_FAILED_ALLOCATION)
	RESULT_MEMORY_CORRUPTION      = Result(C.TVG_RESULT_MEMORY_CORRUPTION)
	RESULT_NOT_SUPPORTED          = Result(C.TVG_RESULT_NOT_SUPPORTED)
	RESULT_UNKNOWN                = Result(C.TVG_RESULT_UNKNOWN)
)

type ResultError struct {
	result Result
}

func (err ResultError) Error() string {
	return map[Result]string{
		RESULT_SUCCESS:                "RESULT_SUCCESS : The value returned in case of a correct request execution.",
		RESULT_INVALID_ARGUMENT:       "RESULT_INVALID_ARGUMENT : The value returned in the event of a problem with the arguments given to the API - e.g. empty paths or null pointers.",
		RESULT_INSUFFICIENT_CONDITION: "RESULT_INSUFFICIENT_CONDITION : The value returned in case the request cannot be processed - e.g. asking for properties of an object, which does not exist.",
		RESULT_FAILED_ALLOCATION:      "RESULT_FAILED_ALLOCATION : The value returned in case of unsuccessful memory allocation.",
		RESULT_MEMORY_CORRUPTION:      "RESULT_MEMORY_CORRUPTION : The value returned in the event of bad memory handling - e.g. failing in pointer releasing or casting",
		RESULT_NOT_SUPPORTED:          "RESULT_NOT_SUPPORTED : The value returned in case of choosing unsupported engine features(options).",
		RESULT_UNKNOWN:                "RESULT_UNKNOWN : The value returned in all other cases.",
	}[err.result]
}

func (err ResultError) Result() Result {
	return err.result
}

var errorHandler func(err ResultError)

/*
SetErrorHandler sets an error handler function that will be called when any function
returns an error that represents a Result that is not RESULT_SUCCESS.

Providing nil for the handler will result in errors being returned from functions in
the normal fashion.

The previous error handler, which may be nil, will be returned.

A lot of the functions in this package return an error if the underlying thorvg C function
call returns an unsuccessful result.
Checking every returned potential error in an application can be tedious.
So an application may want to treat all, or most, thorvg errors in a common way.
Perhaps by logging an error, or by exiting.
In which case setting an error handler may be appropriate.

An application may wish to treat most thorvg errors in a common way, but here and
there needs to handle errors from a function call.
In which case it may appropriate to temporarily remove the handler.

	oldHandler := SetErrorHandler(nil)
	if err := tvg.SomeFunction(...); err != nil {
	  // handle the error
	}
	SetErrorHandler(oldHandler)
*/
func SetErrorHandler(handler func(err ResultError)) func(err ResultError) {
	oldHandler := errorHandler
	errorHandler = handler
	return oldHandler
}

func resultError(result C.Tvg_Result) error {
	if result == C.TVG_RESULT_SUCCESS {
		return nil
	}

	err := ResultError{
		result: Result(result),
	}

	if errorHandler != nil {
		errorHandler(err)
	}
	return err
}

/*
ColorSpace is an enumeration specifying the methods of combining the 8-bit color channels into 32-bit color.
*/
type ColorSpace C.Tvg_Colorspace

const (
	COLORSPACE_ABGR8888  = ColorSpace(C.TVG_COLORSPACE_ABGR8888)  // The channels are joined in the order: alpha, blue, green, red. Colors are alpha-premultiplied.
	COLORSPACE_ARGB8888  = ColorSpace(C.TVG_COLORSPACE_ARGB8888)  // The channels are joined in the order: alpha, red, green, blue. Colors are alpha-premultiplied.
	COLORSPACE_ABGR8888S = ColorSpace(C.TVG_COLORSPACE_ABGR8888S) // The channels are joined in the order: alpha, blue, green, red. Colors are un-alpha-premultiplied. (since 0.13)
	COLORSPACE_ARGB8888S = ColorSpace(C.TVG_COLORSPACE_ARGB8888S) // The channels are joined in the order: alpha, red, green, blue. Colors are un-alpha-premultiplied. (since 0.13)
	COLORSPACE_UNKNOWN   = ColorSpace(C.TVG_COLORSPACE_UNKNOWN)   // Unknown channel data. This is reserved for an initial ColorSpace value. (since 1.0)
)

/*
EngineOption is an enumeration to specify rendering engine behavior.

The availability or behavior of @c ENGINE_OPTION_SMART_RENDER may vary depending on platform or backend support.
It attempts to optimize rendering performance by updating only the regions  of the canvas that have
changed between frames (partial redraw). This can be highly effective in scenarios  where most of the
canvas remains static and only small portions are updated—such as simple animations or GUI interactions.
However, in complex scenes where a large portion of the canvas changes frequently (e.g., full-screen animations
or heavy object movements), the overhead of tracking changes and managing update regions may outweigh the benefits,
resulting in decreased performance compared to the default rendering mode. Thus, it is recommended to benchmark
both modes in your specific use case to determine the optimal setting.

	@since 1.0
*/
type EngineOption C.Tvg_Engine_Option

const (
	ENGINE_OPTION_NONE         = EngineOption(C.TVG_ENGINE_OPTION_NONE)         // No engine options are enabled. This may be used to explicitly disable all optional behaviors.
	ENGINE_OPTION_DEFAULT      = EngineOption(C.TVG_ENGINE_OPTION_DEFAULT)      // Uses the default rendering mode.
	ENGINE_OPTION_SMART_RENDER = EngineOption(C.TVG_ENGINE_OPTION_SMART_RENDER) // Enables automatic partial (smart) rendering optimizations.
)

/*
MaskMethod is an enumeration indicating the method used in the masking of two objects - the target and the source.

	@ingroup ThorVGCapi_Paint
*/
type MaskMethod C.Tvg_Mask_Method

const (
	MASK_METHOD_NONE          = MaskMethod(C.TVG_MASK_METHOD_NONE)          // No Masking is applied.
	MASK_METHOD_ALPHA         = MaskMethod(C.TVG_MASK_METHOD_ALPHA)         // Alpha Masking using the masking target's pixels as an alpha value.
	MASK_METHOD_INVERSE_ALPHA = MaskMethod(C.TVG_MASK_METHOD_INVERSE_ALPHA) // Alpha Masking using the complement to the masking target's pixels as an alpha value.
	MASK_METHOD_LUMA          = MaskMethod(C.TVG_MASK_METHOD_LUMA)          // Alpha Masking using the grayscale (0.2126R + 0.7152G + 0.0722*B) of the masking target's pixels. @since 0.9
	MASK_METHOD_INVERSE_LUMA  = MaskMethod(C.TVG_MASK_METHOD_INVERSE_LUMA)  // Alpha Masking using the grayscale (0.2126R + 0.7152G + 0.0722*B) of the complement to the masking target's pixels. @since 0.11
	MASK_METHOD_ADD           = MaskMethod(C.TVG_MASK_METHOD_ADD)           // Combines the target and source objects pixels using target alpha. (T * TA) + (S * (255 - TA)) @since 1.0
	MASK_METHOD_SUBTRACT      = MaskMethod(C.TVG_MASK_METHOD_SUBTRACT)      // Subtracts the source color from the target color while considering their respective target alpha. (T * TA) - (S * (255 - TA)) @since 1.0
	MASK_METHOD_INTERSECT     = MaskMethod(C.TVG_MASK_METHOD_INTERSECT)     // Computes the result by taking the minimum value between the target alpha and the source alpha and multiplies it with the target color. (T * min(TA, SA)) @since 1.0
	MASK_METHOD_DIFFERENCE    = MaskMethod(C.TVG_MASK_METHOD_DIFFERENCE)    // Calculates the absolute difference between the target color and the source color multiplied by the complement of the target alpha. abs(T - S * (255 - TA)) @since 1.0
	MASK_METHOD_LIGHTEN       = MaskMethod(C.TVG_MASK_METHOD_LIGHTEN)       // Where multiple masks intersect, the highest transparency value is used. @since 1.0
	MASK_METHOD_DARKEN        = MaskMethod(C.TVG_MASK_METHOD_DARKEN)        // Where multiple masks intersect, the lowest transparency value is used. @since 1.0
)

/*
BlendMethod is an enumeration indicates the method used for blending paint. Please refer to the respective formulas for each method.

	@ingroup ThorVGCapi_Paint

	@since 0.15
*/
type BlendMethod C.Tvg_Blend_Method

const (
	BLEND_METHOD_NORMAL      = BlendMethod(C.TVG_BLEND_METHOD_NORMAL)      // Perform the alpha blending(default). S if (Sa == 255), otherwise (Sa * S) + (255 - Sa) * D
	BLEND_METHOD_MULTIPLY    = BlendMethod(C.TVG_BLEND_METHOD_MULTIPLY)    // Takes the RGB channel values from 0 to 255 of each pixel in the top layer and multiples them with the values for the corresponding pixel from the bottom layer. (S * D)
	BLEND_METHOD_SCREEN      = BlendMethod(C.TVG_BLEND_METHOD_SCREEN)      // The values of the pixels in the two layers are inverted, multiplied, and then inverted again. (S + D) - (S * D)
	BLEND_METHOD_OVERLAY     = BlendMethod(C.TVG_BLEND_METHOD_OVERLAY)     // Combines Multiply and Screen blend modes. (2 * S * D) if (2 * D < Da), otherwise (Sa * Da) - 2 * (Da - S) * (Sa - D)
	BLEND_METHOD_DARKEN      = BlendMethod(C.TVG_BLEND_METHOD_DARKEN)      // Creates a pixel that retains the smallest components of the top and bottom layer pixels. min(S, D)
	BLEND_METHOD_LIGHTEN     = BlendMethod(C.TVG_BLEND_METHOD_LIGHTEN)     // Only has the opposite action of Darken Only. max(S, D)
	BLEND_METHOD_COLORDODGE  = BlendMethod(C.TVG_BLEND_METHOD_COLORDODGE)  // Divides the bottom layer by the inverted top layer. D / (255 - S)
	BLEND_METHOD_COLORBURN   = BlendMethod(C.TVG_BLEND_METHOD_COLORBURN)   // Divides the inverted bottom layer by the top layer, and then inverts the result. 255 - (255 - D) / S
	BLEND_METHOD_HARDLIGHT   = BlendMethod(C.TVG_BLEND_METHOD_HARDLIGHT)   // The same as Overlay but with the color roles reversed. (2 * S * D) if (S < Sa), otherwise (Sa * Da) - 2 * (Da - S) * (Sa - D)
	BLEND_METHOD_SOFTLIGHT   = BlendMethod(C.TVG_BLEND_METHOD_SOFTLIGHT)   // The same as Overlay but with applying pure black or white does not result in pure black or white. (1 - 2 * S) * (D ^ 2) + (2 * S * D)
	BLEND_METHOD_DIFFERENCE  = BlendMethod(C.TVG_BLEND_METHOD_DIFFERENCE)  // Subtracts the bottom layer from the top layer or the other way around, to always get a non-negative value. (S - D) if (S > D), otherwise (D - S)
	BLEND_METHOD_EXCLUSION   = BlendMethod(C.TVG_BLEND_METHOD_EXCLUSION)   // The result is twice the product of the top and bottom layers, subtracted from their sum. s + d - (2 * s * d)
	BLEND_METHOD_HUE         = BlendMethod(C.TVG_BLEND_METHOD_HUE)         // Combine with HSL(Sh + Ds + Dl) then convert it to RGB.
	BLEND_METHOD_SATURATION  = BlendMethod(C.TVG_BLEND_METHOD_SATURATION)  // Combine with HSL(Dh + Ss + Dl) then convert it to RGB.
	BLEND_METHOD_COLOR       = BlendMethod(C.TVG_BLEND_METHOD_COLOR)       // Combine with HSL(Sh + Ss + Dl) then convert it to RGB.
	BLEND_METHOD_LUMINOSITY  = BlendMethod(C.TVG_BLEND_METHOD_LUMINOSITY)  // Combine with HSL(Dh + Ds + Sl) then convert it to RGB.
	BLEND_METHOD_ADD         = BlendMethod(C.TVG_BLEND_METHOD_ADD)         // Simply adds pixel values of one layer with the other. (S + D)
	BLEND_METHOD_COMPOSITION = BlendMethod(C.TVG_BLEND_METHOD_COMPOSITION) // Used for intermediate composition. @since 1.0
)

/*
Type is an enumeration indicating the ThorVG object type value.

ThorVG's drawing objects can return object type values, allowing you to identify the specific type of each object.

	@ingroup ThorVGCapi_Paint

	@see tvg_paint_get_type()
	@see tvg_gradient_get_type()

	@since 1.0
*/
type Type C.Tvg_Type

const (
	TYPE_UNDEF       = Type(C.TVG_TYPE_UNDEF)       // Undefined type.
	TYPE_SHAPE       = Type(C.TVG_TYPE_SHAPE)       // A shape type paint.
	TYPE_SCENE       = Type(C.TVG_TYPE_SCENE)       // A scene type paint.
	TYPE_PICTURE     = Type(C.TVG_TYPE_PICTURE)     // A picture type paint.
	TYPE_TEXT        = Type(C.TVG_TYPE_TEXT)        // A text type paint.
	TYPE_LINEAR_GRAD = Type(C.TVG_TYPE_LINEAR_GRAD) // A linear gradient type.
	TYPE_RADIAL_GRAD = Type(C.TVG_TYPE_RADIAL_GRAD) // A radial gradient type.
)

/*
PathCommand is an enumeration specifying the values of the path commands accepted by ThorVG.
*/
type PathCommand C.Tvg_Path_Command

const (
	PATH_COMMAND_CLOSE    = PathCommand(C.TVG_PATH_COMMAND_CLOSE)    // Ends the current sub-path and connects it with its initial point - corresponds to Z command in the svg path commands.
	PATH_COMMAND_MOVE_TO  = PathCommand(C.TVG_PATH_COMMAND_MOVE_TO)  // Sets a new initial point of the sub-path and a new current point - corresponds to M command in the svg path commands.
	PATH_COMMAND_LINE_TO  = PathCommand(C.TVG_PATH_COMMAND_LINE_TO)  // Draws a line from the current point to the given point and sets a new value of the current point - corresponds to L command in the svg path commands.
	PATH_COMMAND_CUBIC_TO = PathCommand(C.TVG_PATH_COMMAND_CUBIC_TO) // Draws a cubic Bezier curve from the current point to the given point using two given control points and sets a new value of the current point - corresponds to C command in the svg path commands.
)

/*
StrokeCap is an enumeration determining the ending type of a stroke in the open sub-paths.
*/
type StrokeCap C.Tvg_Stroke_Cap

const (
	STROKE_CAP_BUTT   = StrokeCap(C.TVG_STROKE_CAP_BUTT)   // The stroke ends exactly at each of the two endpoints of a sub-path. For zero length sub-paths no stroke is rendered.
	STROKE_CAP_ROUND  = StrokeCap(C.TVG_STROKE_CAP_ROUND)  // The stroke is extended in both endpoints of a sub-path by a half circle, with a radius equal to the half of a stroke width. For zero length sub-paths a full circle is rendered.
	STROKE_CAP_SQUARE = StrokeCap(C.TVG_STROKE_CAP_SQUARE) // The stroke is extended in both endpoints of a sub-path by a rectangle, with the width equal to the stroke width and the length equal to the half of the stroke width. For zero length sub-paths the square is rendered with the size of the stroke width.
)

/*
StrokeJoin is an enumeration specifying how to fill the area outside the gradient bounds.
*/
type StrokeJoin C.Tvg_Stroke_Join

const (
	STROKE_JOIN_MITER = StrokeJoin(C.TVG_STROKE_JOIN_MITER) // The outer corner of the joined path segments is spiked. The spike is created by extension beyond the join point of the outer edges of the stroke until they intersect. In case the extension goes beyond the limit, the join style is converted to the Bevel style.
	STROKE_JOIN_ROUND = StrokeJoin(C.TVG_STROKE_JOIN_ROUND) // The outer corner of the joined path segments is rounded. The circular region is centered at the join point.
	STROKE_JOIN_BEVEL = StrokeJoin(C.TVG_STROKE_JOIN_BEVEL) // The outer corner of the joined path segments is bevelled at the join point. The triangular region of the corner is enclosed by a straight line between the outer corners of each stroke.
)

/*
StrokeFill is an enumeration specifying how to fill the area outside the gradient bounds.
*/
type StrokeFill C.Tvg_Stroke_Fill

const (
	STROKE_FILL_PAD     = StrokeFill(C.TVG_STROKE_FILL_PAD)     // The remaining area is filled with the closest stop color.
	STROKE_FILL_REFLECT = StrokeFill(C.TVG_STROKE_FILL_REFLECT) // The gradient pattern is reflected outside the gradient area until the expected region is filled.
	STROKE_FILL_REPEAT  = StrokeFill(C.TVG_STROKE_FILL_REPEAT)  // The gradient pattern is repeated continuously beyond the gradient area until the expected region is filled.
)

/*
FillRule is an enumeration specifying the algorithm used to establish which parts of the shape are treated as the inside of the shape.
*/
type FillRule C.Tvg_Fill_Rule

const (
	FILL_RULE_NON_ZERO = FillRule(C.TVG_FILL_RULE_NON_ZERO) // A line from the point to a location outside the shape is drawn. The intersections of the line with the path segment of the shape are counted. Starting from zero, if the path segment of the shape crosses the line clockwise, one is added, otherwise one is subtracted. If the resulting sum is non zero, the point is inside the shape.
	FILL_RULE_EVEN_ODD = FillRule(C.TVG_FILL_RULE_EVEN_ODD) // A line from the point to a location outside the shape is drawn and its intersections with the path segments of the shape are counted. If the number of intersections is an odd number, the point is inside the shape.
)

/*
TextWrap is an enumeration specifying how text should be wrapped and truncated.
*/
type TextWrap C.Tvg_Text_Wrap

const (
	TEXT_WRAP_NONE        = TextWrap(C.TVG_TEXT_WRAP_NONE)        // Do not wrap text. Text is rendered on a single line and may overflow the bounding area.
	TEXT_WRAP_CHARACTER   = TextWrap(C.TVG_TEXT_WRAP_CHARACTER)   // Wrap at the character level. If a word cannot fit, it is broken into individual characters to fit the line.
	TEXT_WRAP_WORD        = TextWrap(C.TVG_TEXT_WRAP_WORD)        // Wrap at the word level. Words that do not fit are moved to the next line.
	TEXT_WRAP_SMART       = TextWrap(C.TVG_TEXT_WRAP_SMART)       // Smart choose wrapping method: word wrap first, falling back to character wrap if a word does not fit.
	TEXT_WRAP_ELLIPSIS    = TextWrap(C.TVG_TEXT_WRAP_ELLIPSIS)    // Truncate overflowing text and append an ellipsis ("...") at the end. Typically used for single-line labels.
	TEXT_WRAP_HYPHENATION = TextWrap(C.TVG_TEXT_WRAP_HYPHENATION) // Reserved. No Support.
)
