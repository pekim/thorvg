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

	if err := tvg.FontLoadData("DejaVuSans", data.DejaVuSans, ""); err != nil {
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
	if err := canvas.Add(bg); err != nil {
		panic(err)
	}

	// draw text
	text := tvg.TextNew()
	if err := text.SetFont("DejaVuSans"); err != nil {
		panic(err)
	}
	if err := text.SetColor(0, 0, 0); err != nil {
		panic(err)
	}
	if err := text.SetSize(18); err != nil {
		panic(err)
	}
	// if err := text.WrapMode(tvg.TEXT_WRAP_WORD); err != nil {
	// 	panic(err)
	// }
	if err := text.Translate(20, 20); err != nil {
		panic(err)
	}
	// if err := text.SetText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."); err != nil {
	if err := text.SetText("Lorem ipsum dolor sit amet, consectetur..."); err != nil {
		panic(err)
	}
	if err := canvas.Add(text); err != nil {
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

	file, err := os.Create("example/text/text.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("error encoding image: %v", err)
	}
}
