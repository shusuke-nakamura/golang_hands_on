package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	v := 0.0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	p := widget.NewProgressBar()
	b := widget.NewButton("Up!", func() {
		v += 0.1
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})
	w.SetContent(
		container.NewVBox(
			l, p, b,
		),
	)

	w.ShowAndRun()
}
