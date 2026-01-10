package main

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"

	tvg "github.com/pekim/thorvg"
)

func main() {
	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("thorvg - SDL example", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	var err error
	err = tvg.EngineInit(2)
	if err != nil {
		panic(err)
	}

	canvas := tvg.GlCanvasCreate()

	context := unsafe.Pointer(sdl.GLGetCurrentContext())

	var windowWidth int32
	var windowHeight int32

	resized := func() {
		sdl.GetWindowSize(window, &windowWidth, &windowHeight)
		err = canvas.GlSetTarget(nil, nil, context, 0,
			uint(windowWidth), uint(windowHeight), tvg.COLORSPACE_ABGR8888S)
		if err != nil {
			panic(err)
		}
	}
	resized()

Outer:
	for {
		var event sdl.Event
		for sdl.WaitEvent(&event) {
			switch event.Type() {

			case sdl.EventQuit:
				break Outer

			case sdl.EventKeyDown:
				mod := event.Key().Mod &^ sdl.KeymodNum
				if (mod == sdl.KeymodNone && event.Key().Scancode == sdl.ScancodeEscape) ||
					((mod == sdl.KeymodLCtrl || mod == sdl.KeymodRCtrl) && event.Key().Scancode == sdl.ScancodeQ) {
					break Outer
				}

			case sdl.EventWindowResized:
				resized()

			case sdl.EventWindowExposed:
				// background
				{
					rect := tvg.ShapeNew()
					err = rect.AppendRect(0, 0, float32(windowWidth), float32(windowHeight), 0, 0, true)
					if err != nil {
						panic(err)
					}
					err = rect.SetFillColor(255, 255, 255, 255)
					if err != nil {
						panic(err)
					}
					err = canvas.Push(rect)
					if err != nil {
						panic(err)
					}
				}

				// foreground
				{
					rect := tvg.ShapeNew()
					err = rect.AppendRect(50, 50, 200, 200, 20, 20, true)
					if err != nil {
						panic(err)
					}
					err = rect.SetFillColor(255, 0, 0, 100)
					if err != nil {
						panic(err)
					}
					err = canvas.Push(rect)
					if err != nil {
						panic(err)
					}
				}

				err = canvas.Draw(true)
				if err != nil {
					panic(err)
				}
				err = canvas.Sync()
				if err != nil {
					panic(err)
				}

				sdl.RenderPresent(renderer)
			}
		}
	}
}
