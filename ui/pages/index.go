package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/cdumange/twitchscan/consts"
	"github.com/cdumange/twitchscan/ui/images"
)

// NewIndex will return the main app window
func NewIndex(app fyne.App) fyne.Window {
	w := app.NewWindow(consts.AppName)
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(canvas.NewImageFromResource(images.ResourceHarleyPng))
	w.CenterOnScreen()
	return w
}
