package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

/*
Dialog Options Include:

Colour picker
Conform action
File opener
Form with validation
Information popup
Custom dialog
*/

/*
Keyboard shortcuts can be defined.

Fyne has an API for storing user configurations in a cross platform manner.

Fyne release V2.2.0 Added support for a system tray menu. Displays an icon on macOS, windows, linux that when tapped will pop out a menu as specified by the app.

Data binding added in Fyne release V2.0.0 makes it easier to connect many widgets to a data source that will update over time.
*/

/*
Drawing Actions Include:

Rectangle
Text
Line
Circle
Image
Raster
Gradient
Animation
*/

func main() {
	myApp := app.New()
	// A Canvas is the area which an application is drawn within
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	// CanvasObjects can be drawn on a Canvas
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rectangle := canvas.NewRectangle(blue)
	myCanvas.SetContent(rectangle)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		// A CanvasObject can be updated with Refresh
		rectangle.FillColor = green
		rectangle.Refresh()
	}()

	// setContentToText(myCanvas)
	// setContentToCircle(myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func SetContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func SetContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
