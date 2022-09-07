package ui

import (
	"github.com/cdumange/twitchscan/ui/atoms"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Front the basic first page
type Front struct {
	app.Compo
}

// Render the app renderering
func (f *Front) Render() app.UI {

	return app.Div().Body(atoms.NewText("test", atoms.TextTypeHeader))
}
