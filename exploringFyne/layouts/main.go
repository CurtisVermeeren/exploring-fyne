package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

/*

Layouts Include:

Horizontal box
Vertical box
Center
Form
Grid
Grid wrap
Border
Max
Padded

Layouts can be combined for more complex application structures.

*/

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("hello", green)
	text2 := canvas.NewText("world", green)

	// Objects can be moved around if no layout is set
	text2.Move(fyne.NewPos(20, 20))
	//content := container.NewWithoutLayout(text1, text2)

	// Or create a new grid layout with two comments and add the text objects to each
	content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
