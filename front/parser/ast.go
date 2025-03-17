package parser

import (
	"bytes"

	tok "github.com/aenlemmea/mika/front/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Represents a program context. Always the "root" to hold other nodes.
type Context struct {
	Statements []Statement
}

func (ctx *Context) String() string {
	var out bytes.Buffer

	for _, s := range ctx.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ctx *Context) TokenLiteral() string {
	if len(ctx.Statements) > 0 {
		return ctx.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// TR Statements :: Statement
type TrStatement struct {
	PrimToken tok.Token
	Name      *Identifier
	Value     Expression
}

func (trs *TrStatement) statementNode()       {}
func (trs *TrStatement) TokenLiteral() string { return trs.PrimToken.Value }
func (trs *TrStatement) String() string {
	var out bytes.Buffer

	out.WriteString(trs.TokenLiteral() + " ")
	out.WriteString(trs.Name.String() + " ")
	out.WriteString("= ")

	if trs.Value != nil {
		out.WriteString(trs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// RET Statements :: Statement
type RetStatement struct {
	RetToken tok.Token
	RetValue Expression
}

func (rs *RetStatement) statementNode()       {}
func (rs *RetStatement) TokenLiteral() string { return rs.RetToken.Value }
func (rs *RetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.RetValue != nil {
		out.WriteString(rs.RetValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// EXPR Statement :: Statement
type ExprStatement struct {
	ExprToken tok.Token
	Expr      Expression
}

func (expr *ExprStatement) statementNode()       {}
func (expr *ExprStatement) TokenLiteral() string { return expr.ExprToken.Value }
func (expr *ExprStatement) String() string {
	if expr.Expr != nil {
		return expr.Expr.String()
	}
	return ""
}

// IDENTIFIER :: Expression
type Identifier struct {
	IdToken tok.Token
	Value   string
}

func (ids *Identifier) expressionNode()      {}
func (ids *Identifier) TokenLiteral() string { return ids.IdToken.Value }
func (ids *Identifier) String() string {
	return ids.Value
}

// INTEGER VALUES :: Expression
type IntVal struct {
	IntToken tok.Token
	Value    int64
}

func (intv *IntVal) expressionNode()      {}
func (intv *IntVal) TokenLiteral() string { return intv.IntToken.Value }
func (intv *IntVal) String() string {
	return intv.IntToken.Value
}

// PREFIX EXPR :: Expression
type PrfxExpr struct {
	PrfxToken tok.Token
	Operator  string
	Right     Expression
}

func (prxe *PrfxExpr) expressionNode()      {}
func (prxe *PrfxExpr) TokenLiteral() string { return prxe.PrfxToken.Value }
func (prxe *PrfxExpr) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(prxe.Operator)
	out.WriteString(prxe.Right.String())
	out.WriteString(")")

	return out.String()
}

// INFIX EXPR :: Expression
type InfxExpr struct {
	InfxToken tok.Token
	Left      Expression
	Operator  string
	Right     Expression
}

func (infxe *InfxExpr) expressionNode()      {}
func (infxe *InfxExpr) TokenLiteral() string { return infxe.InfxToken.Value }
func (infxe *InfxExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infxe.Left.String())
	out.WriteString(" " + infxe.Operator + " ")
	out.WriteString(infxe.Right.String())
	out.WriteString(")")

	return out.String()
}
