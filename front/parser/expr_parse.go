package parser

import "github.com/aenlemmea/mika/front/token"

// Function types that returns Expression.
type (
	prfxFn func() Expression
	infxFn func(left Expression) Expression
)

const (
	LOWEST      = iota
	EQUALS      = iota
	LESSGREATER = iota
	SUM         = iota
	PROD        = iota
	PREFIX      = iota
	CALL        = iota
)

var precedences = map[token.TokenKind]int{
	token.EQ:      EQUALS,
	token.NEQ:     EQUALS,
	token.GREATER: LESSGREATER,
	token.LESS:    LESSGREATER,
	token.PLUS:    SUM,
	token.MINUS:   SUM,
	token.DIV:     PROD,
	token.MULT:    PROD,
}

func (prs *Parser) peekPrec() int {
	if p, ok := precedences[prs.peekToken.Kind]; ok {
		return p
	}
	return LOWEST
}

func (prs *Parser) currPrec() int {
	if p, ok := precedences[prs.currToken.Kind]; ok {
		return p
	}
	return LOWEST
}
