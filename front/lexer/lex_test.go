package lexer

import (
	"testing"
	"github.com/aenlemmea/mika/front/token"

)

func TestNextToken(t *testing.T) {
	input := `tr a = 2;
tr b = 10;

tr add = fn(x, y) {
	x + y;
};

tr res = add:(a, b);

< 5 > 10;
!-/*10;

if 2 > 3 {
	ret true;
} else {
	ret false;
}

5 neq 10;

10 eq 10;

`

	tests := []struct {
		expectedKind token.TokenKind
		expectedValue string
	}{
		{token.TR, "tr"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON,";"},
		{token.TR, "tr"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.TR, "tr"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LCURLY, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RCURLY, "}"},
		{token.SEMICOLON, ";"},
		{token.TR, "tr"},
		{token.IDENT, "res"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LESS, "<"},
		{token.INT, "5"},
		{token.GREATER, ">"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MULT, "*"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.INT, "2"},
		{token.GREATER, ">"},
		{token.INT, "3"},
		{token.LCURLY, "{"},
		{token.RET, "ret"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RCURLY, "}"},
		{token.ELSE, "else"},
		{token.LCURLY, "{"},
		{token.RET, "ret"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RCURLY, "}"},
		{token.INT, "5"},
		{token.NEQ, "neq"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.EQ, "eq"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Kind != tt.expectedKind {
			t.Fatalf("tests[%d] - incorrect tokenkind. Expected: %q Got: %q",
			i, tt.expectedKind, tok.Kind)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - incorrect value. Expected: %q Got: %q",
			i, tt.expectedValue, tok.Value)
		}
	}
}
