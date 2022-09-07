package atoms

import (
	"testing"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TestClassed(t *testing.T) {
	v := "a value"
	t.Run(TextTypeHeader, func(t *testing.T) {

		c := NewText(v, TextTypeHeader)

		d := app.NewClientTester(c)
		defer d.Close()

		app.TestMatch(c, app.TestUIDescriptor{
			Path:     app.TestPath(0),
			Expected: app.Span().Class(TextTypeHeader).Text(v),
		})
	})
}
