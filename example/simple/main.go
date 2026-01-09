package main

import (
	"image"
	"image/png"
	"log"
	"os"

	tvg "github.com/pekim/thorvg"
)

func main() {
	width := 400
	height := 400

	tvg.EngineInit(2) //nolint:errcheck

	canvas := tvg.SwCanvasCreate(tvg.ENGINE_OPTION_DEFAULT)
	canvas.SwSetTarget(uint(width), uint(width), uint(height), tvg.COLORSPACE_ARGB8888) //nolint:errcheck

	rect := tvg.ShapeNew()
	rect.AppendRect(50, 50, 200, 200, 20, 20, true) //nolint:errcheck
	rect.SetFillColor(100, 100, 100, 255)           //nolint:errcheck
	canvas.Push(rect)                               //nolint:errcheck

	canvas.Draw(true) //nolint:errcheck
	canvas.Sync()     //nolint:errcheck

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = canvas.Buffer()

	file, err := os.Create("example/simple/simple.png")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
}
