package ui

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Front the basic first page
type Front struct {
	app.Compo
}

// Render the app renderering
func (f *Front) Render() app.UI {
	return app.Div().Style("color", "white").Body(app.Text("hello world"))
}
