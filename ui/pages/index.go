package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// NewIndex will return the main app window
func NewIndex(app fyne.App) fyne.Window {
	w := app.NewWindow("Hello World !")

	w.SetContent(widget.NewLabel("Hello World!"))
	return w
}
