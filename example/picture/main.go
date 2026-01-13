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

	if err := tvg.EngineInit(2); err != nil {
		panic(err)
	}

	canvas := tvg.SwCanvasCreate(tvg.ENGINE_OPTION_DEFAULT)
	if err := canvas.SwSetTarget(uint(width), uint(width), uint(height), tvg.COLORSPACE_ARGB8888); err != nil {
		panic(err)
	}

	// background
	bg := tvg.ShapeNew()
	if err := bg.AppendRect(0, 0, float32(width), float32(height), 0, 0, true); err != nil {
		panic(err)
	}
	if err := bg.SetFillColor(255, 255, 255, 255); err != nil {
		panic(err)
	}
	if err := canvas.Push(bg); err != nil {
		panic(err)
	}

	// load and resize image
	picture := tvg.PictureNew()
	// if err := picture.LoadDataText(string(data.Gopher), "svg", ""); err != nil {
	if err := picture.LoadData(data.Gopher, "svg", ""); err != nil {
		panic(err)
	}
	w, h, err := picture.GetSize()
	if err != nil {
		panic(err)
	}
	var size float32 = 400
	w = w / h * size
	h = size
	if err := picture.SetSize(w, h); err != nil {
		panic(err)
	}

	// draw image, centred
	if err := picture.Translate((float32(width)-w)/2, (float32(height)-h)/2); err != nil {
		panic(err)
	}
	if err := canvas.Push(picture); err != nil {
		panic(err)
	}

	// finish
	if err := canvas.Draw(true); err != nil {
		panic(err)
	}
	if err := canvas.Sync(); err != nil {
		panic(err)
	}

	buffer := canvas.Buffer()
	swizzle.BGRA(buffer)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = buffer

	file, err := os.Create("example/picture/picture.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("error encoding image: %v", err)
	}
}
