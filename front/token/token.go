package token

type TokenKind string

type Token struct {
	Kind TokenKind
	Value string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"
	STR = "STR"

	ASSIGN = "="
	PLUS = "+"
	DIV = "/"
	MULT = "*"
	
	MINUS = "-"
	BANG = "!"
	LOWAND = "~"

	GREATER = ">"
	LESS = "<"
	EQ = "EQ"
	NEQ = "NEQ"

	COLON = ":"
	SEMICOLON = ";"
	COMMA = ","

	LPAREN = "("
	RPAREN = ")"
	LCURLY = "{"
	RCURLY = "}"

	FUNCTION = "FUNCTION"
	TR = "TR"
	IMPORT = "IMPORT"
	AS = "AS"
	IF = "IF"
	ELSE = "ELSE"
	RET = "RET"
	MUTATES = "MUTATES"

	TRUE = "true"
	FALSE = "false"
)

var keywords = // TODO: Somehow have this be non var.
	map[string]TokenKind {
		"fn": FUNCTION,
		"tr": TR,
		"import": IMPORT,
		"as": AS,
		"if" : IF,
		"else" : ELSE,
		"ret" : RET,
		"true" : TRUE,
		"false" : FALSE,
		"neq" : NEQ,
		"eq" : EQ,
		"mutates" : MUTATES,
}

func SearchIdentInKeyword(ident string) TokenKind {
	// keykind is same as IDENT!
	if keykind, ok := keywords[ident]; ok {
		return keykind
	}
	return IDENT
}
