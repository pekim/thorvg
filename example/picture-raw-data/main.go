package main

import (
	"bytes"
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

	// load  image
	primateImage, err := png.Decode(bytes.NewReader(data.Primate))
	if err != nil {
		panic(err)
	}
	pix := primateImage.(*image.RGBA).Pix
	swizzle.BGRA(pix)
	picture := tvg.PictureNew()
	_ = picture.LoadRaw(
		pix,
		uint(primateImage.Bounds().Dx()),
		uint(primateImage.Bounds().Dy()),
		tvg.COLORSPACE_ARGB8888,
	)
	w, h, _ := picture.GetSize()

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

	file, err := os.Create("example/picture-raw-data/picture.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("error encoding image: %v", err)
	}
}
