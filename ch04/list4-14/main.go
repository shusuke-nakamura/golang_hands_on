package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Bottom", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)

	w.SetContent(
		// fyne.NewContainerWithLayoutは古い
		// container.NewXXXを使用するようにとコンパイルにワーニングが表示されたので、
		// container.NewXXXで再実装
		// fyne.NewContainerWithLayout(
		// 	layout.NewBorderLayout(
		// 		bt, bb, bl, br,
		// 	),
		// 	bt, bb, bl, br,
		// 	widget.NewLabel("Conter."),
		// ),
		container.NewBorder(
			bt, bb, bl, br,
			widget.NewLabel("Conter."),
		),
	)
	w.ShowAndRun()
}
