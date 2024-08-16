package lexer

import "monkey/token"

/* Methods of Lexer

func (l *Lexer) NextToken() token.Token
func (l *Lexer) readChar()
func (l *Lexer) readIdentifier() string

*/

type Lexer struct {
	input        string
	position     int  // current position (index) in input (points to current char)
	readPosition int  // current reading position (index) in input (after current char)
	ch           byte // current char being examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// TODO: add unicode support (currently ASCII)
// Method, returns next character and advances position in the input
func (l *Lexer) readChar() {
	// end of input reached ?
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Sets current character equal to ASCII's NUL char, (which is then interpreted as EOF?)
	} else {
		// Checks next position with read position, and sets it to this char
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	// After updating the current position, pos and pos+1 are the same.
	l.readPosition += 1 // Update l readPosition to prepare for reading the next character.

}

// The syntax (l *Lexer) peekChar() means we are adding the peekChar() method to Lexer
// l is a pointer to a Lexer instance, and the reference of that instance in our method
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	// This allows for snake case names like foo_bar
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'z' || ch == '_'
}
