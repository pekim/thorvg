package draw

import (
	tvg "github.com/pekim/thorvg"
)

// SimpleShapes draws some shapes on a canvas.
//
// It is used by multiple examples, all of which should render identically.
func SimpleShapes(canvas tvg.Canvas, width float32, height float32) error {
	var err error

	// background
	{
		rect := tvg.ShapeNew()
		err = rect.AppendRect(0, 0, width, height, 0, 0, true)
		if err != nil {
			return err
		}
		err = rect.SetFillColor(255, 255, 255, 255)
		if err != nil {
			return err
		}
		err = canvas.Push(rect)
		if err != nil {
			return err
		}
	}

	// foreground
	{
		rect := tvg.ShapeNew()
		err = rect.AppendRect(50, 50, 200, 200, 20, 20, true)
		if err != nil {
			return err
		}
		err = rect.SetFillColor(255, 0, 0, 100)
		if err != nil {
			return err
		}
		err = canvas.Push(rect)
		if err != nil {
			return err
		}
	}

	// finish
	err = canvas.Draw(true)
	if err != nil {
		return err
	}
	err = canvas.Sync()
	if err != nil {
		return err
	}

	return nil
}
