package parser

import (
	"testing"

	"github.com/aenlemmea/mika/front/lexer"
)

func TestPrattParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-x * y;",
			"((-x) * y)",
		},
		{
			"~-5;",
			"(~(-5))",
		},
		{
			"3 + 4 * 5 eq 3 * 1 + 4 * 5;",
			"((3 + (4 * 5)) eq ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parse := New(lex)
		program := parse.ParseContext()
		checkParserErrors(t, parse)

		got := program.String()
		if got != tt.expected {
			t.Errorf("Expected: %q Got: %q", tt.expected, got)
		}
	}
}
