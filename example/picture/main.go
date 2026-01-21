package main

import (
	"image"
	"image/png"
	"log"
	"os"

	tvg "github.com/pekim/thorvg"
	"github.com/pekim/thorvg/example/data"
	"github.com/pekim/thorvg/swizzle"
)

func main() {
	width := 600
	height := 600

	tvg.SetErrorHandler(func(err tvg.ResultError) { panic(err) })
	_ = tvg.EngineInit(2)

	canvas := tvg.SwCanvasCreate(tvg.ENGINE_OPTION_DEFAULT)
	_ = canvas.SwSetTarget(uint(width), uint(width), uint(height), tvg.COLORSPACE_ARGB8888)

	// background
	bg := tvg.ShapeNew()
	_ = bg.AppendRect(0, 0, float32(width), float32(height), 0, 0, true)
	_ = bg.SetFillColor(255, 255, 255, 255)
	_ = canvas.Add(bg)

	// load and resize image
	picture := tvg.PictureNew()
	_ = picture.LoadDataText(string(data.Gopher), "svg", "")
	w, h, _ := picture.GetSize()
	var size float32 = 400
	w = w / h * size
	h = size
	_ = picture.SetSize(w, h)

	// draw image, centred
	_ = picture.Translate((float32(width)-w)/2, (float32(height)-h)/2)
	_ = canvas.Add(picture)

	// finish
	_ = canvas.Draw(true)
	_ = canvas.Sync()

	buffer := canvas.Buffer()
	swizzle.BGRA(buffer)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = buffer

	file, err := os.Create("example/picture/picture.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	_ = png.Encode(file, img)
}
