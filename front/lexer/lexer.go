package lexer

import "github.com/aenlemmea/mika/front/token"

type Lexer struct {
	input string // Later change this to a FileStringRep object (filename, contents)
	pos int
	readpos int
	ch byte
}


func (l *Lexer) peekChar() byte {
	if l.readpos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readpos]
	}
}

// Move ahead in the input. Do NOT use for peeking.
func (l *Lexer) readChar() /* MUTATES */ {
	if l.readpos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readpos]
	}
	l.pos = l.readpos
	l.readpos += 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() 
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case ':':
			tok = newToken(token.COLON, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '{':
			tok = newToken(token.LCURLY, l.ch)
		case '}':
			tok = newToken(token.RCURLY, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '*':
			tok = newToken(token.MULT, l.ch)
		case '/':
			tok = newToken(token.DIV, l.ch)
		case '!':
			tok = newToken(token.BANG, l.ch)
		case '~':
			tok = newToken(token.LOWAND, l.ch)
		case '<':
			tok = newToken(token.LESS, l.ch)
		case '>':
			tok = newToken(token.GREATER, l.ch)
		case 0:
			tok.Value = ""
			tok.Kind = token.EOF
		default:
			if isLetter(l.ch) {
				tok.Value = l.readIdentifier()
				tok.Kind = token.SearchIdentInKeyword(tok.Value)
				return tok
			} else if isDigit(l.ch) {
				tok.Kind = token.INT
				tok.Value = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() /* MUTATES */ {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string /* MUTATES */ {
	position := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9';
}

func (l *Lexer) readIdentifier() string /* MUTATES */  {
	position := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenKind token.TokenKind, ch byte) token.Token {
	return token.Token{Kind: tokenKind, Value: string(ch)}
}
