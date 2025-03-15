package parser

import (
	"fmt"
	"testing"

	"github.com/aenlemmea/mika/front/lexer"
)

func TestTrStatements(t *testing.T) {
	input := `
tr x = 5;
tr y = 10;
tr foobar = 12345;
`
	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseContext()
	checkParserErrors(t, parse)

	if program == nil {
		t.Fatalf("Context not found")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statem := program.Statements[i]
		if !testTrStatement(t, statem, tt.expectedIdent) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, prs *Parser) {
	errors := prs.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errros", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}

func testTrStatement(t *testing.T, aststatem Statement, name string) bool {
	if aststatem.TokenLiteral() != "tr" {
		t.Errorf(" statem.TokenLiteral not 'let'. Got: %q", aststatem.TokenLiteral())
		return false
	}
	trstatem, ok := aststatem.(*TrStatement)
	if !ok {
		t.Errorf(" statem not *TrStatement. Got: %T", aststatem)
		return false
	}

	if trstatem.Name.Value != name {
		t.Errorf("trstatem.Name.Value not '%s'. Got: %s", name, trstatem.Name.Value)
		return false
	}

	if trstatem.Name.TokenLiteral() != name {
		t.Errorf(" statem.Name not '%s'. Got: %s", name, trstatem.Name)
		return false
	}

	return true
}

func TestRetStatements(t *testing.T) {
	input := `
ret foo;
ret baz;
ret 232;
`
	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseContext()
	checkParserErrors(t, parse)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got: %d", len(program.Statements))
	}

	for _, statem := range program.Statements {
		retStatem, ok := statem.(*RetStatement)
		if !ok {
			t.Errorf("retStatem is not *RetStatement. Got=%T", statem)
			continue
		}

		if retStatem.TokenLiteral() != "ret" {
			t.Errorf("retStatem.TokenLiteral() is not 'return', Got=%q", retStatem.TokenLiteral())
		}
	}
}

func TestIdExpression(t *testing.T) {
	input := "foobar;"

	lex := lexer.New(input)
	parse := New(lex)
	program := parse.ParseContext()
	checkParserErrors(t, parse)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has not enough statements. Got: %d", len(program.Statements))
	}

	statem, ok := program.Statements[0].(*ExprStatement)

	if !ok {
		t.Fatalf("Program.Statements[0] is not ExprStatement. Got: %T", program.Statements[0])
	}

	ident, ok := statem.Expr.(*Identifier)
	if !ok {
		t.Fatalf("Expr not Identifier. Got: %T", statem.Expr)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. Got: %s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. Got: %s", "foobar", ident.TokenLiteral())
	}
}

func TestIntValExpr(t *testing.T) {
	input := "5;"

	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseContext()
	checkParserErrors(t, parse)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has not enough statements. Got: %d", len(program.Statements))
	}

	statem, ok := program.Statements[0].(*ExprStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ExprStatement. Got: %T", program.Statements[0])
	}

	val, ok := statem.Expr.(*IntVal)
	if !ok {
		t.Fatalf("Expr is not IntVal. Got: %T", statem.Expr)
	}

	if val.Value != 5 {
		t.Errorf("IntVal value is not 5. Got: %d", val.Value)
	}

	if val.TokenLiteral() != "5" {
		t.Errorf("val.TokenLiteral is not %s. Got: %s", "5", val.TokenLiteral())
	}
}

func TestPrfxParseExpr(t *testing.T) {
	prfxInput := []struct {
		input    string
		operator string
		intVal   int64
	}{
		{"!10;", "!", 10},
		{"-15;", "-", 15},
		{"~15;", "~", 15},
	}

	for _, tt := range prfxInput {
		lex := lexer.New(tt.input)
		parse := New(lex)
		program := parse.ParseContext()
		checkParserErrors(t, parse)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. Got: %d\n", len(program.Statements))
		}
		statem, ok := program.Statements[0].(*ExprStatement)

		if !ok {
			t.Fatalf("statem is not ExprStatement. Got: %T", program.Statements[0])
		}

		expr, ok := statem.Expr.(*PrfxExpr)
		if !ok {
			t.Fatalf("statem is not PrfxExpr. Got: %T", statem.Expr)
		}

		if expr.Operator != tt.operator {
			t.Fatalf("Operator is not '%s', Got: %s", tt.operator, expr.Operator)
		}

		if !testIntVal(t, expr.Right, tt.intVal) {
			return
		}
	}
}

func testIntVal(t *testing.T, expr Expression, val int64) bool {
	intgr, ok := expr.(*IntVal)
	if !ok {
		t.Errorf("expr is not IntVal. Got: %T", expr)
		return false
	}

	if intgr.Value != val {
		t.Errorf("intgr.Value is not %d. Got: %d", val, intgr.Value)
	}

	if intgr.TokenLiteral() != fmt.Sprintf("%d", val) {
		t.Errorf("integ.TokenLiteral is not %d. Got: %s", val, intgr.TokenLiteral())
		return false
	}

	return true
}

func TestInfxParseExpr(t *testing.T) {
	infxInput := []struct {
		input    string
		leftVal  int64
		operator string
		rightVal int64
	}{
		{"10 + 10;", 10, "+", 10},
		{"5 - 5;", 5, "-", 5},
		//{"5 * 5;", 5, "*", 5},
		{"5 eq 5;", 5, "eq", 5},
		{"5 neq 5;", 5, "neq", 5},
	}

	for _, tt := range infxInput {
		lex := lexer.New(tt.input)
		parse := New(lex)
		program := parse.ParseContext()
		checkParserErrors(t, parse)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. Got: %d\n", len(program.Statements))
		}
		statem, ok := program.Statements[0].(*ExprStatement)

		if !ok {
			t.Fatalf("statem is not ExprStatement. Got: %T", program.Statements[0])
		}

		expr, ok := statem.Expr.(*InfxExpr)
		if !ok {
			t.Fatalf("statem is not PrfxExpr. Got: %T", statem.Expr)
		}

		if !testIntVal(t, expr.Left, tt.leftVal) {
			return
		}

		if expr.Operator != tt.operator {
			t.Fatalf("Operator is not '%s', Got: %s", tt.operator, expr.Operator)
		}

		if !testIntVal(t, expr.Right, tt.rightVal) {
			return
		}
	}
}
