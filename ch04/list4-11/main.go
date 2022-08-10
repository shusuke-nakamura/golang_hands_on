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
	ne := widget.NewEntry()
	pe := widget.NewPasswordEntry()

	w.SetContent(
		container.NewVBox(
			l,
			widget.NewForm(
				widget.NewFormItem("Name", ne),
				widget.NewFormItem("Pass", pe),
			),
			widget.NewButton("OK", func() {
				l.SetText(ne.Text + " & " + pe.Text)
			}),
		),
	)

	w.ShowAndRun()
}
