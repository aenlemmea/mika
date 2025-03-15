package parser

import (
	"fmt"
	"strconv"

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

	errors []string

	prfxFnMap map[token.TokenKind]prfxFn
	infxFnMap map[token.TokenKind]infxFn
}

func (prs *Parser) registerPrefix(tokKind token.TokenKind, fn prfxFn) {
	prs.prfxFnMap[tokKind] = fn
}

// TODO Generics
func (prs *Parser) registerInfix(tokKind token.TokenKind, fn infxFn) {
	prs.infxFnMap[tokKind] = fn
}

func New(lex *lexer.Lexer) *Parser {
	prs := &Parser{lex: lex,
		errors: []string{},
	}

	prs.peekToken.Kind = token.ILLEGAL /* Failsafe */
	prs.setNextToken()
	prs.setNextToken()

	prs.prfxFnMap = make(map[token.TokenKind]prfxFn)
	prs.registerPrefix(token.IDENT, prs.parseIdentifier)
	prs.registerPrefix(token.INT, prs.parseIntValExpr)
	prs.registerPrefix(token.BANG, prs.parsePrfxExpr)
	prs.registerPrefix(token.MINUS, prs.parsePrfxExpr)
	prs.registerPrefix(token.LOWAND, prs.parsePrfxExpr)

	prs.infxFnMap = make(map[token.TokenKind]infxFn)
	prs.registerInfix(token.PLUS, prs.parseInfxExpr)
	prs.registerInfix(token.MULT, prs.parseInfxExpr)
	prs.registerInfix(token.DIV, prs.parseInfxExpr)
	prs.registerInfix(token.MINUS, prs.parseInfxExpr)
	prs.registerInfix(token.EQ, prs.parseInfxExpr)
	prs.registerInfix(token.NEQ, prs.parseInfxExpr)

	return prs
}

func (prs *Parser) parseIdentifier() Expression {
	return &Identifier{IdToken: prs.currToken, Value: prs.currToken.Value}
}

func (prs *Parser) Errors() []string {
	return prs.errors
}

func (prs *Parser) setNextToken() {
	prs.currToken = prs.peekToken
	prs.peekToken = prs.lex.NextToken()
}

func (prs *Parser) ParseContext() *Context {
	program := &Context{}
	program.Statements = []Statement{}

	for !prs.currTokenIs(token.EOF) {
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
	case token.RET:
		return prs.parseRetStatement()
	default:
		return prs.parseExprStatement()
	}
}
func (prs *Parser) parseExprStatement() *ExprStatement {
	statem := &ExprStatement{ExprToken: prs.currToken}

	statem.Expr = prs.parseExpr(LOWEST)

	if prs.peekTokenIs(token.SEMICOLON) {
		prs.setNextToken()
	} else {
		msg := fmt.Sprintf("Missing semicolon @ Expr. Seeing: %q", prs.peekToken)
		prs.errors = append(prs.errors, msg)
	}
	return statem
}

func (prs *Parser) parseExpr(priority int) Expression {
	prefix := prs.prfxFnMap[prs.currToken.Kind]
	if prefix == nil {
		return nil
	}

	leftExpr := prefix()

	for !prs.peekTokenIs(token.SEMICOLON) && priority < prs.peekPrec() {
		infx := prs.infxFnMap[prs.peekToken.Kind]
		if infx == nil {
			return leftExpr
		}

		prs.setNextToken()
		leftExpr = infx(leftExpr)
	}
	return leftExpr
}

func (prs *Parser) parseTrStatement() *TrStatement {
	statem := &TrStatement{PrimToken: prs.currToken}

	if !prs.expectPeek(token.IDENT) {
		return nil
	}

	statem.Name = &Identifier{IdToken: prs.currToken, Value: prs.currToken.Value}
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

func (prs *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, prs.peekToken.Kind)
	prs.errors = append(prs.errors, msg)
}

func (prs *Parser) expectPeek(t token.TokenKind) bool /* Mutates */ {
	if prs.peekTokenIs(t) {
		prs.setNextToken()
		return true
	} else {
		prs.peekError(t)
		return false
	}
}

func (prs *Parser) parseIntValExpr() Expression {
	val := &IntVal{IntToken: prs.currToken}

	value, err := strconv.ParseInt(prs.currToken.Value, 0, 64)
	if err != nil {
		prs.errors = append(prs.errors, "Could not parse integer")
		return nil
	}
	val.Value = value

	return val
}

func (prs *Parser) parseRetStatement() *RetStatement {
	statem := &RetStatement{RetToken: prs.currToken}

	prs.setNextToken()

	for !prs.currTokenIs(token.SEMICOLON) {
		prs.setNextToken()
	}

	return statem
}

func (prs *Parser) parsePrfxExpr() Expression {
	expr := &PrfxExpr{
		PrfxToken: prs.currToken,
		Operator:  prs.currToken.Value,
	}

	prs.setNextToken()
	expr.Right = prs.parseExpr(PREFIX)
	return expr
}

func (prs *Parser) parseInfxExpr(left Expression) Expression {
	expr := &InfxExpr{
		InfxToken: prs.currToken,
		Operator:  prs.currToken.Value,
		Left:      left,
	}

	priority := prs.currPrec()
	prs.setNextToken()
	expr.Right = prs.parseExpr(priority)
	return expr
}
