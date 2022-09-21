package lexer

import (
	"fmt"
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt:= range tests {
		tok := l.NextToken()
		// fmt.Printf("[%d]...%q, %q", i, tok.Type, tok.Literal)
		// fmt.Println()
		// fmt.Printf("[%d]...%q, %q", i, tt.expectedType, tt.expectedLiteral)
		// fmt.Println()
		// fmt.Println()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}