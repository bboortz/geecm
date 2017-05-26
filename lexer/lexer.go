package lexer

// Lexer stores info for lexing input
type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition + 1
	l.readPosition++
}

// NextToken gets the next token from input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case token.COMMENT:
		tok = newToken(token.COMMENT, l.ch)
	}

	l.readChar()

	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}
}

// Scan is the scanner method for the lexer
func Scan(str string) string {
	return "COMMENT"
}

// New creates a new lexer from an input
func New(in string) *Lexer {
	l := &Lexer{input: in}
	l.readChar()
	return l
}
