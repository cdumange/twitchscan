package pages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Index(t *testing.T) {
	o := generateIndex()

	assert.True(t, o.Visible())
}
