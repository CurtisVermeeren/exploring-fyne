package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

/*

Widgets Include:

Accordion menu list
Buttons
Cards
Checkbox
Entry
FileIcon
Forms
Hyperlinks
Icons
Labels
Progress bars
Radio buttons
Select dropdowns
Line separators
Sliders
Text grids
Toolbars

Lists
Tables
Trees

App tabs
Scroll container
Split containers

*/

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")

	// Widgets are special canvas objects with interactive elements to them
	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}
