package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Simple")

	// Create a binding for a string
	str := binding.NewString()
	str.Set("Initial value")

	text := widget.NewLabelWithData(str)
	w.SetContent(text)

	go func() {
		// Update the bound string after 2 seconds
		time.Sleep(time.Second * 2)
		str.Set("A new string")
	}()

	w.ShowAndRun()
}
