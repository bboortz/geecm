package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsComment(t *testing.T) {
	assert.Equal(t, true, isComment("#"))
	assert.NotEqual(t, true, isComment("!"))
}
