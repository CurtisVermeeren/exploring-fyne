package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("Hi!")

	// Create a vertial layout and add widgets
	w.SetContent(container.NewVBox(
		// The label and entry widgets are bound together on str
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))

	w.ShowAndRun()
}
