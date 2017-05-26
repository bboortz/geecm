package lexer

import (
	"github.com/bboortz/geecm/token"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestScanComment(t *testing.T) {
	input := "# test comment"
	expected := token.Token{Type: token.COMMENT, Literal: "#"}

	lex := New(input)

	tok := lex.NextToken()
	spew.Dump(expected)
	spew.Dump(tok)
	spew.Dump(token.COMMENT)
	assert.Equal(t, expected.Type, tok.Type)
	assert.Equal(t, expected.Literal, tok.Literal)

}
