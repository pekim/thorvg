package main

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"

	tvg "github.com/pekim/thorvg"
	"github.com/pekim/thorvg/example/data"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	windowWidth := 800
	windowHeight := 600

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "thorvg - GLFW lottie example", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	tvg.SetErrorHandler(func(err tvg.ResultError) { panic(err) })
	_ = tvg.EngineInit(2)
	canvas := tvg.GlCanvasCreate()

	onSize := func(width int, height int) {
		context := window.GetGLXContext()
		_ = canvas.GlSetTarget(nil, nil, unsafe.Pointer(context), 0,
			uint(width), uint(height), tvg.COLORSPACE_ABGR8888S)

		windowWidth = width
		windowHeight = height
	}
	onSize(1, 1)

	window.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
		onSize(width, height)
	})
	window.SetKeyCallback(func(_w *glfw.Window, key glfw.Key, _scancode int, _action glfw.Action, mods glfw.ModifierKey) {
		if (mods == 0 && key == glfw.KeyEscape) || (mods == glfw.ModControl && key == glfw.KeyQ) {
			window.SetShouldClose(true)
		}
	})

	draw := false
	window.SetRefreshCallback(func(_ *glfw.Window) {
		draw = true
	})

	for !window.ShouldClose() {
		if draw {
			// background
			bg := tvg.ShapeNew()
			_ = bg.AppendRect(0, 0, float32(windowWidth), float32(windowHeight), 0, 0, true)
			_ = bg.SetFillColor(255, 255, 255, 255)
			_ = canvas.Push(bg)

			// picture
			picture := tvg.PictureNew()
			_ = picture.LoadData(data.LottieGearsAnimation, "lottie+json", "")
			width, height, _ := picture.GetSize()
			scale := min(
				float32(windowWidth)/width,
				float32(windowHeight)/height,
			)
			_ = picture.Scale(scale)
			width *= scale
			height *= scale
			_ = picture.Translate(float32(windowWidth-int(width))/2, float32(windowHeight-int(height))/2)
			_ = canvas.Push(picture)

			// finish
			_ = canvas.Draw(true)
			_ = canvas.Sync()

			window.SwapBuffers()
			draw = false
		}

		glfw.WaitEvents()
	}
}
