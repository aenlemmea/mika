package parser

import tok "github.com/aenlemmea/mika/front/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Context struct {
	Statements []Statement
}

func (ctx *Context) TokenLiteral() string {
	if len(ctx.Statements) > 0 {
		return ctx.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type TrStatement struct {
	PrimToken tok.Token
	Name *Identifier
	Value Expression
}

func (trs *TrStatement) statementNode() {}
func (trs *TrStatement) TokenLiteral() string { return trs.PrimToken.Value }

type Identifier struct {
	IdToken tok.Token
	Value string
}

func (ids *Identifier) expressionNode() {}
func (ids *Identifier) TokenLiteral() string { retrun ids.IdToken.Value }



