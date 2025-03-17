package parser

import (
	"testing"

	tok "github.com/aenlemmea/mika/front/token"
)

// Testing the string functionality provided by the context
func TestString(t *testing.T) {
	program := &Context{
		Statements: []Statement{
			&TrStatement{
				PrimToken: tok.Token{Kind: tok.TR, Value: "tr"},
				Name: &Identifier{
					IdToken: tok.Token{Kind: tok.IDENT, Value: "foobaz"},
					Value:   "foobaz",
				},
				Value: &Identifier{
					IdToken: tok.Token{Kind: tok.IDENT, Value: "bar"},
					Value:   "bar",
				},
			},
		},
	}

	if program.String() != "tr foobaz = bar;" {
		t.Errorf("program.String() is incorrect. Got: %q", program.String())
	}
}
