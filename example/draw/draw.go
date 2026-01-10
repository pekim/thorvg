package draw

import (
	tvg "github.com/pekim/thorvg"
)

// SimpleShapes draws some shapes on a canvas.
// It is used by multiple examples, all of which should render identically.
//
// What's drawn is similar to what'a shown in the "Quick Start" section of
// the thorvg readme.
// https://github.com/thorvg/thorvg?tab=readme-ov-file#quick-start
func SimpleShapes(canvas tvg.Canvas, width float32, height float32) error {
	// background
	bg := tvg.ShapeNew()
	if err := bg.AppendRect(0, 0, width, height, 0, 0, true); err != nil {
		return err
	}
	if err := bg.SetFillColor(255, 255, 255, 255); err != nil {
		return err
	}
	if err := canvas.Push(bg); err != nil {
		return err
	}

	// rectangle
	rect := tvg.ShapeNew()
	if err := rect.AppendRect(50, 50, 200, 200, 20, 20, true); err != nil {
		return err
	}
	if err := rect.SetFillColor(255, 0, 0, 100); err != nil {
		return err
	}
	if err := canvas.Push(rect); err != nil {
		return err
	}

	// circle, with gradient fill
	circle := tvg.ShapeNew()
	if err := circle.AppendCircle(400, 400, 100, 100, true); err != nil {
		return err
	}
	gradient := tvg.RadialGradientNew()
	if err := gradient.RadialSet(400, 400, 200, 400, 400, 0); err != nil {
		return err
	}
	if err := gradient.SetColorStops([]tvg.ColorStop{
		{Offset: 0.0, R: 255, G: 255, B: 255, A: 255},
		{Offset: 1.0, R: 0, G: 0, B: 0, A: 255},
	}); err != nil {
		return err
	}
	if err := circle.SetGradient(gradient); err != nil {
		return err
	}
	if err := canvas.Push(circle); err != nil {
		return err
	}

	// finish
	if err := canvas.Draw(true); err != nil {
		return err
	}
	if err := canvas.Sync(); err != nil {
		return err
	}

	return nil
}
