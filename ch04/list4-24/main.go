package main

import (
	"myentry"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := myentry.NewMyEntry(func(e *myentry.MyEntry) {
		s := e.Text
		e.SetText("")
		l.SetText("you type '" + s + "'.")
	})

	w.SetContent(
		container.NewVBox(
			l, e,
		),
	)

	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()
}
