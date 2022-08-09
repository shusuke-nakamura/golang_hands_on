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
	sl := widget.NewSelect([]string{"Eins", "Twei", "Drei"}, func(s string) {
		l.SetText("selected: " + s)
	})

	w.SetContent(
		container.NewVBox(
			l, sl,
		),
	)

	w.ShowAndRun()
}
