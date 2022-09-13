package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/cdumange/twitchscan/ui/pages"
)

func main() {
	a := app.New()
	index := pages.NewIndex(a)
	index.ShowAndRun()
}
