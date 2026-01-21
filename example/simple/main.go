package main

import (
	"image"
	"image/png"
	"log"
	"os"

	tvg "github.com/pekim/thorvg"
	"github.com/pekim/thorvg/example/draw"
	"github.com/pekim/thorvg/swizzle"
)

func main() {
	width := 600
	height := 600

	tvg.SetErrorHandler(func(err tvg.ResultError) { panic(err) })
	_ = tvg.EngineInit(2)

	canvas := tvg.SwCanvasCreate(tvg.ENGINE_OPTION_DEFAULT)
	_ = canvas.SwSetTarget(uint(width), uint(width), uint(height), tvg.COLORSPACE_ARGB8888)
	draw.SimpleShapes(canvas, float32(width), float32(height))

	buffer := canvas.Buffer()
	swizzle.BGRA(buffer)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = buffer

	file, err := os.Create("example/simple/simple.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("error encoding image: %v", err)
	}
}
