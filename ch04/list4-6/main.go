package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	c := widget.NewCheck("Check!", func(f bool) {
		if f {
			l.SetText("CHECKED!!")
		} else {
			l.SetText("not checked!.")
		}
	})
	c.SetChecked(true)
	w.SetContent(
		container.NewVBox(
			l, c,
		),
	)

	w.ShowAndRun()
}
