package token

// TokenType an int that represents a char like '#'
type TokenType int

// Token a token type and literal value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// COMMENT represents a detected comment
	COMMENT = '#'
	// SEMI represents a detected semicolon
	SEMI = ';'
	// EOF represents the end of file
	EOF = 0
)
