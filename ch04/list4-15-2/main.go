package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		container.NewGridWithRows(3,
			widget.NewButton("One", nil),
			widget.NewButton("Two", nil),
			widget.NewButton("Three", nil),
			widget.NewButton("Four", nil),
			layout.NewSpacer(),
			widget.NewButton("Five", nil),
			widget.NewButton("Six", nil),
			layout.NewSpacer(),
			widget.NewButton("Seven", nil),
			widget.NewButton("Eight", nil),
			widget.NewButton("Nine", nil),
			widget.NewButton("Ten", nil),
		),
	)
	w.ShowAndRun()
}
