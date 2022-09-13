package pages

import (
	"fyne.io/fyne/v2/app"
	"github.com/cdumange/twitchscan/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Index(t *testing.T) {
	a := app.New()

	w := NewIndex(a)

	assert.Equal(t, consts.AppName, w.Title())
}
