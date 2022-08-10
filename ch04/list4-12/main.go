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
	// 参考書のGroupがfyne v2ではないので、checkgroupとformを使用して同じことを実現した。
	// ckg := widget.NewCheckGroup([]string{"check 1", "check 2"}, func(s []string) {
	// 	re := "result: "
	// 	for _, v := range s {
	// 		re += v
	// 	}
	// 	l.SetText(re)
	// })
	ckg := widget.NewCheckGroup([]string{"check 1", "check 2"}, nil)
	// ck1 := widget.NewCheck("check 1", nil)
	// ck2 := widget.NewCheck("check 2", nil)
	w.SetContent(
		container.NewVBox(
			l,
			widget.NewForm(
				widget.NewFormItem("Group", ckg),
			),
			widget.NewButton("OK", func() {
				re := "result: "
				for _, v := range ckg.Selected {
					re += v
				}
				l.SetText(re)
			}),
		),
	)
	w.ShowAndRun()
}
