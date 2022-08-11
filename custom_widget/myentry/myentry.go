package myentry

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MyEntry is custom entry
type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

// NewEntry create MyEntry
func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

// KeyDown is Keydown Event.
func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.entered(e)
	default:
		e.Entry.KeyDown(key)
	}
}
