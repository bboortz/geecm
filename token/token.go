package token

// TokenType an int that represents a char like '#'
type TokenType int

// Token a token type and literal value
type Token struct {
	Type    TokenType
	Literal string
}

// Tokens is a array of multiple detected Token
type Tokens []Token

const (
	// COMMENT represents a detected comment
	COMMENT = '#'
	// SPACE represents a detected space
	SPACE = ' '
	// SEMI represents a detected semicolon
	SEMI = ';'
	// EOL represents the end of line
	EOL = 0
	// EOF represents the end of file
	EOF = 0
)
