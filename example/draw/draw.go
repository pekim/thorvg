package draw

// import (
// 	tvg "github.com/pekim/thorvg"
// )

// // SimpleShapes draws some shapes on a canvas.
// // It is used by multiple examples, all of which should render identically.
// //
// // What's drawn is similar to what'a shown in the "Quick Start" section of
// // the thorvg readme.
// // https://github.com/thorvg/thorvg?tab=readme-ov-file#quick-start
// func SimpleShapes(canvas tvg.Canvas, width float32, height float32) {
// 	// background
// 	bg := tvg.ShapeNew()
// 	_ = bg.AppendRect(0, 0, width, height, 0, 0, true)
// 	_ = bg.SetFillColor(255, 255, 255, 255)
// 	_ = canvas.Add(bg)

// 	// rectangle
// 	rect := tvg.ShapeNew()
// 	_ = rect.AppendRect(50, 50, 200, 200, 20, 20, true)
// 	_ = rect.SetFillColor(255, 0, 0, 100)
// 	_ = canvas.Add(rect)

// 	// circle, with gradient fill
// 	circle := tvg.ShapeNew()
// 	_ = circle.AppendCircle(400, 400, 100, 100, true)
// 	gradient := tvg.RadialGradientNew()
// 	_ = gradient.RadialSet(400, 400, 200, 400, 400, 0)
// 	_ = gradient.SetColorStops([]tvg.ColorStop{
// 		{Offset: 0.0, R: 255, G: 255, B: 255, A: 255},
// 		{Offset: 1.0, R: 0, G: 0, B: 0, A: 255},
// 	})
// 	_ = circle.SetGradient(gradient)
// 	_ = canvas.Add(circle)

// 	// finish
// 	_ = canvas.Draw(true)
// 	_ = canvas.Sync()
// }
