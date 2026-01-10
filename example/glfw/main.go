package main

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"

	tvg "github.com/pekim/thorvg"
	"github.com/pekim/thorvg/example/draw"
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

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "thorvg - GLFW example", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	err = tvg.EngineInit(2)
	if err != nil {
		panic(err)
	}
	canvas := tvg.GlCanvasCreate()

	onSize := func(width int, height int) {
		context := window.GetGLXContext()
		err = canvas.GlSetTarget(nil, nil, unsafe.Pointer(context), 0,
			uint(width), uint(height), tvg.COLORSPACE_ABGR8888S)
		if err != nil {
			panic(err)
		}

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

	for !window.ShouldClose() {
		err := draw.SimpleShapes(canvas, float32(windowWidth), float32(windowHeight))
		if err != nil {
			panic(err)
		}

		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
