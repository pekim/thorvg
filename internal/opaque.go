package internal

// Pointers to opaque structs.

/*
Canvas is a structure responsible for managing and drawing graphical elements.

It sets up the target buffer, which can be drawn on the screen. It stores the Paint objects (Shape, Scene, Picture).
*/
type Canvas uintptr

/*
Paint is a structure representing a graphical element.

@warning The TvgPaint objects cannot be shared between Canvases.
*/
type Paint uintptr

/*
Gradient is a structure representing a gradient fill of a Paint object.
*/
type Gradient uintptr

/*
Saver is a structure representing an object that enables to save a Paint object into a file.
*/
type Saver uintptr

/*
Animation is a structure representing an animation controller object.
*/
type Animation uintptr

/*
Accessor is a structure representing an object that enables iterating through a scene's descendents.
*/
type Accessor uintptr
