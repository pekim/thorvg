package draw

import (
	tvg "github.com/pekim/thorvg"
)

// SimpleShapes draws some shapes on a canvas.
//
// It is used by multiple examples, all of which should render identically.
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

	// foreground
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

	// finish
	if err := canvas.Draw(true); err != nil {
		return err
	}
	if err := canvas.Sync(); err != nil {
		return err
	}

	return nil
}
