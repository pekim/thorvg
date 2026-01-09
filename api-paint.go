package thorvg

/*
Paint is a structure representing a graphical element.

@warning The TvgPaint objects cannot be shared between Canvases.
*/
type Paint interface {
	paint() uintptr
}
