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

// Parser provides the access to the lexer and acts as a global subscriber of errors slice. The prfxFnMap, infxFnMap are two maps holding callables for the pratt parser to work on.
// The parser has no access to the AST itself and purely mutates it and creates it, this leads to cleaner sync-"ability" and less overall indirection.
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

// Constructor for building the parser. The failsafe token is needed for detection of incorrect peekToken values.
// Callable registration for the maps are also initialized here.
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
	prs.registerPrefix(token.TRUE, prs.parseBooleanExpr)
	prs.registerPrefix(token.FALSE, prs.parseBooleanExpr)
	prs.registerPrefix(token.LPAREN, prs.parseGroupedExpr)

	prs.infxFnMap = make(map[token.TokenKind]infxFn)
	prs.registerInfix(token.PLUS, prs.parseInfxExpr)
	prs.registerInfix(token.MULT, prs.parseInfxExpr)
	prs.registerInfix(token.DIV, prs.parseInfxExpr)
	prs.registerInfix(token.MINUS, prs.parseInfxExpr)
	prs.registerInfix(token.EQ, prs.parseInfxExpr)
	prs.registerInfix(token.NEQ, prs.parseInfxExpr)

	return prs
}

// Entrypoint of the parser to start parsing. It creates the ancestor nodes and hands over control to `parseStatement()` for further identication.
func (prs *Parser) ParseContext() *Context /* MUTATES */ {
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

// Detect the kind of grammar we are at via the current token.
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

func (prs *Parser) Errors() []string {
	return prs.errors
}

// The actual iterator over tokens, one token at a time design.
func (prs *Parser) setNextToken() /* MUTATES */ {
	prs.currToken = prs.peekToken
	prs.peekToken = prs.lex.NextToken()
}

func (prs *Parser) currTokenIs(t token.TokenKind) bool {
	return prs.currToken.Kind == t
}

func (prs *Parser) peekTokenIs(t token.TokenKind) bool {
	return prs.peekToken.Kind == t
}

// Utility method, majorly provided for non core tool development.
func (prs *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, prs.peekToken.Kind)
	prs.errors = append(prs.errors, msg)
}

// Utility method, majorly provided for non core tool development.
func (prs *Parser) expectPeek(t token.TokenKind) bool /* Mutates */ {
	if prs.peekTokenIs(t) {
		prs.setNextToken()
		return true
	} else {
		prs.peekError(t)
		return false
	}
}

// PARSING CORE METHODS

func (prs *Parser) parseIdentifier() Expression {
	return &Identifier{IdToken: prs.currToken, Value: prs.currToken.Value}
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

func (prs *Parser) parseBooleanExpr() Expression {
	return &Boolean{BoolToken: prs.currToken, Val: prs.currTokenIs(token.TRUE)}
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
