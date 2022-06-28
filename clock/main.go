package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	// clock label displays the time
	clock := widget.NewLabel("")
	time.LoadLocation("America/New_York")
	updateTime(clock)

	// Add the label to window
	w.SetContent(clock)
	w.SetContent(clock)

	// Update the time
	go func() {
		time.LoadLocation("Toronto")
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}

// updateTime takes a fyne label and sets its text to the current time
func updateTime(clock *widget.Label) {
	// Locad location of time
	local, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	formatted := time.Now().In(local).Format("Time: 03:04:05")
	clock.SetText(formatted)
}
