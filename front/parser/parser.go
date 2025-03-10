package parser

import (
	"errors"
	"github.com/aenlemmea/mika/front/token"
	"github.com/aenlemmea/mika/front/lexer"
)


// TODO See https://stackoverflow.com/questions/37135193/how-to-set-default-values-in-go-structs
// for the problems with such a design. 
// TODO Instead extract out a device interface and have parser be a implementer.
type Parser struct {
	lex *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	prs := &Parser{lex : lex}
	
	prs.peekToken = token.ILLEGAL /* Failsafe */
	prs.setNextToken()
	prs.setNextToken()

	return p
}

func (prs *Parser) setNextToken() {
	prs.currToken = prs.peekToken
	prs.peekToken = prs.lex.NextToken()
}

func (prs *Parser) ParseEntry() (*ast.Context, error) {
	return nil
}



