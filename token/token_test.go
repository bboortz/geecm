package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	res1 := Token{Type: COMMENT, Literal: "#"}
	res2 := Token{Type: COMMENT, Literal: "#"}
	assert.Equal(t, res1.Type, res1.Type)
	assert.Equal(t, res1.Literal, res2.Literal)
}
