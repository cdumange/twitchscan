package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/cdumange/twitchscan/consts"
)

// NewIndex will return the main app window
func NewIndex(app fyne.App) fyne.Window {
	w := app.NewWindow(consts.AppName)
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(widget.NewLabel("Hello World!"))
	w.CenterOnScreen()
	return w
}
