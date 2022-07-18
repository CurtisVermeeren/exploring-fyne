package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Conversion")
	// Float value bound to slider
	f := binding.NewFloat()
	// String value for label bound to f float value
	str := binding.FloatToString(f)
	// String value with formatting from f float as percent
	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)

	// Create a layout container with 3 widgets
	w.SetContent(container.NewVBox(
		widget.NewSliderWithData(0, 100.0, f),
		widget.NewLabelWithData(str),
		widget.NewLabelWithData(short),
	))

	w.ShowAndRun()
}
