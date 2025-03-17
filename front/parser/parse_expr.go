package parser

import (
	"fmt"

	"github.com/aenlemmea/mika/front/token"
)

// Function types that returns Expression.
type (
	prfxFn func() Expression
	infxFn func(left Expression) Expression
)

// Priority values for supported operators
const (
	LOWEST      = iota
	EQUALS      = iota
	LESSGREATER = iota
	SUM         = iota
	PROD        = iota
	PREFIX      = iota
	CALL        = iota
)

// Precedence map for priority values
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

// Actual method based on the Pratt Parsing algorithm responsible for seeing forward and calling functions based on precedences
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

// Parse the expression nodes.
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

// Parse the prefix expressions
func (prs *Parser) parsePrfxExpr() Expression {
	expr := &PrfxExpr{
		PrfxToken: prs.currToken,
		Operator:  prs.currToken.Value,
	}

	prs.setNextToken()
	expr.Right = prs.parseExpr(PREFIX)
	return expr
}

// Parse the infix expressions
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
