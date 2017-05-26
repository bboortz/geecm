package token

// TokenType an int that represents a char like '<'
type TokenType int

// Token a token type and literal value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	COMMENT = '#'
	SEMI    = ';'
	EOF     = 0
)

// isComment checks if input string is a comment
func isComment(str byte) bool {

	switch str {
	case COMMENT:
		return true
	}
	return false

}
