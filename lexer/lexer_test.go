package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanComment(t *testing.T) {
	input := "# test comment"
	expected := "COMMENT"

	lex := New(input)
	result := Scan(input)

	assert.Equal(t, expected, result)
}
