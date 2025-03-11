package parser

import (
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
		if !testTrStatement(t, statem, tt.expectedIdent) { return }
	}
}

func testTrStatement(t *testing.T, aststatem Statement, name string) bool {
	if  aststatem.TokenLiteral() != "tr" {
		t.Errorf(" statem.TokenLiteral not 'let'. Got: %q", aststatem.TokenLiteral())
		return false
	}
	trstatem, ok :=  aststatem.(*TrStatement)
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
