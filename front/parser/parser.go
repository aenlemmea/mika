package parser

import (
	"github.com/aenlemmea/mika/front/lexer"
	"github.com/aenlemmea/mika/front/token"
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
	
	prs.peekToken.Kind = token.ILLEGAL /* Failsafe */
	prs.setNextToken()
	prs.setNextToken()

	return prs
}

func (prs *Parser) setNextToken() {
	prs.currToken = prs.peekToken
	prs.peekToken = prs.lex.NextToken()
}

func (prs *Parser) ParseContext() (*Context) {
	program := &Context{}
	program.Statements = []Statement{}

	for prs.currToken.Kind != token.EOF {
		statem := prs.parseStatement() 
		if statem != nil {
			program.Statements = append(program.Statements, statem)
		}
		prs.setNextToken()
	}

	return program
}

func (prs *Parser) parseStatement() Statement {
	switch prs.currToken.Kind {
		case token.TR:
			return prs.parseTrStatement()
		default:
			return nil
	}
}

func (prs *Parser) parseTrStatement() *TrStatement {
	statem := &TrStatement{Token: prs.currToken}

	if !prs.expectPeek(token.IDENT) {
		return nil
	}

	statem.Name = &Identifier{Token: prs.currToken, Value: prs.currToken.Literal}
	if !prs.expectPeek(token.ASSIGN) {
		return nil
	}

	for !prs.currTokenIs(token.SEMICOLON) {
		prs.setNextToken()
	}

	return statem
}

func (prs *Parser) currTokenIs(t token.TokenKind) bool {
	return prs.currToken.Kind == t
}

func (prs *Parser) peekTokenIs(t token.TokenKind) bool {
	return prs.peekToken.Kind == t
}

func (prs *Parser) expectPeek(t token.TokenKind) bool /* Mutates */
	if prs.peekTokenIs(t) {
		prs.setNextToken()
		return true
	} else {
		return false
	}
}
