package main

import (
	"fmt"
	"runtime"
	"time"
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
	// var canvas tvg.Canvas

	animation := tvg.LottieAnimationNew()
	picture := animation.GetPicture()
	_ = picture.LoadData(data.LottieGearsAnimation, "lottie+json", "")

	onSize := func(width int, height int) {
		windowWidth = width
		windowHeight = height

		_ = canvas.Destroy()
		_ = animation.Del()
		animation = tvg.LottieAnimationNew()
		picture = animation.GetPicture()
		_ = picture.LoadData(data.LottieGearsAnimation, "lottie+json", "")

		context := window.GetGLXContext()
		canvas = tvg.GlCanvasCreate()

		_ = canvas.GlSetTarget(nil, nil, unsafe.Pointer(context), 0, uint(width), uint(height), tvg.COLORSPACE_ABGR8888S)

		// background
		bg := tvg.ShapeNew()
		_ = bg.AppendRect(0, 0, float32(windowWidth), float32(windowHeight), 0, 0, true)
		_ = bg.SetFillColor(255, 255, 255, 255)
		_ = canvas.Add(bg)

		// picture
		picWidth, picHeight, _ := picture.GetSize()
		scale := min(
			float32(windowWidth)/float32(picWidth),
			float32(windowHeight)/float32(picHeight),
		)
		_ = picture.Scale(scale)
		scaledWidth := float32(picWidth) * scale
		scaledHeight := float32(picHeight) * scale
		_ = picture.Translate(float32(windowWidth-int(scaledWidth))/2, float32(windowHeight-int(scaledHeight))/2)
		_ = canvas.Add(picture)

		// finish
		_ = canvas.Draw(true)
		_ = canvas.Sync()

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

	start := time.Now()

	go func() {
		for {
			time.Sleep(time.Second / 60)
			glfw.PostEmptyEvent()
		}
	}()

	for !window.ShouldClose() {
		// fmt.Println("draw")
		var progress float32
		duration, _ := animation.GetDuration()
		duration *= 1000.0
		elapsed := time.Since(start).Milliseconds()
		if elapsed == 0 || duration == 0 {
			progress = 0.0
		} else {
			forward := (elapsed/int64(duration))%2 == 0
			if elapsed%int64(duration) == 0 {
				if forward {
					progress = 0.0
				} else {
					progress = 1.0
				}
			} else {
				progress = float32(elapsed%int64(duration)) / duration
			}
		}

		totalFrame, _ := animation.GetTotalFrame()
		frame := progress * totalFrame
		// fmt.Println("frame", frame)
		errorHandler := tvg.SetErrorHandler(nil)
		err := animation.SetFrame(frame)
		if err == nil {
			_ = canvas.Update(picture)
			_ = canvas.Draw(false)
			_ = canvas.Sync()
		} else if err.(tvg.ResultError).Result() != tvg.RESULT_INSUFFICIENT_CONDITION {
			fmt.Println(err)
		}
		tvg.SetErrorHandler(errorHandler)

		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
