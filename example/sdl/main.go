package main

// import (
// 	"unsafe"

// 	"github.com/jupiterrider/purego-sdl3/sdl"

// 	tvg "github.com/pekim/thorvg"
// 	"github.com/pekim/thorvg/example/draw"
// )

// func main() {
// 	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
// 		panic(sdl.GetError())
// 	}

// 	defer sdl.Quit()
// 	if !sdl.Init(sdl.InitVideo) {
// 		panic(sdl.GetError())
// 	}

// 	var window *sdl.Window
// 	var renderer *sdl.Renderer
// 	if !sdl.CreateWindowAndRenderer("thorvg - SDL example", 1280, 720, sdl.WindowResizable, &window, &renderer) {
// 		panic(sdl.GetError())
// 	}
// 	defer sdl.DestroyWindow(window)
// 	defer sdl.DestroyRenderer(renderer)

// 	tvg.SetErrorHandler(func(err tvg.ResultError) { panic(err) })
// 	_ = tvg.EngineInit(2)

// 	canvas := tvg.GlCanvasCreate()

// 	context := unsafe.Pointer(sdl.GLGetCurrentContext())

// 	var windowWidth int32
// 	var windowHeight int32

// 	resized := func() {
// 		sdl.GetWindowSize(window, &windowWidth, &windowHeight)
// 		_ = canvas.GlSetTarget(nil, nil, context, 0,
// 			uint(windowWidth), uint(windowHeight), tvg.COLORSPACE_ABGR8888S)
// 	}
// 	resized()

// Outer:
// 	for {
// 		var event sdl.Event
// 		for sdl.WaitEvent(&event) {
// 			switch event.Type() {

// 			case sdl.EventQuit:
// 				break Outer

// 			case sdl.EventKeyDown:
// 				mod := event.Key().Mod &^ sdl.KeymodNum
// 				if (mod == sdl.KeymodNone && event.Key().Scancode == sdl.ScancodeEscape) ||
// 					((mod == sdl.KeymodLCtrl || mod == sdl.KeymodRCtrl) && event.Key().Scancode == sdl.ScancodeQ) {
// 					break Outer
// 				}

// 			case sdl.EventWindowResized:
// 				resized()

// 			case sdl.EventWindowExposed:
// 				draw.SimpleShapes(canvas, float32(windowWidth), float32(windowHeight))
// 				sdl.RenderPresent(renderer)
// 			}
// 		}
// 	}
// }
