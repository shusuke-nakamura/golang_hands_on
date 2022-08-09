package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	s := widget.NewSlider(0.0, 100.)
	b := widget.NewButton("Check", func() {
		l.SetText("value: " + strconv.Itoa(int(s.Value)))
	})

	w.SetContent(
		container.NewVBox(
			l, s, b,
		),
	)

	w.ShowAndRun()
}
