package atoms

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Text struct {
	app.Compo

	value string
	class string
}

func (f Text) Render() app.UI {
	return app.Span().Class(f.class).Text(f.value)
}

func NewText(value, style string) *Text {
	return &Text{value: value, class: style}
}
