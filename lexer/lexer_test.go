package lexer

import (
	"testing"

	"github.com/YazanAssaf/tiny-interpreter/token"
)

func TestNextToken(t *testing.T) {
	src := `=+(){},;`

	expectedResults := []struct {
		expectedType    token.TokenType
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

	lexer := NewLexer(src)

	for i, result := range expectedResults {
		token := lexer.NextToken()

		if token.Type != result.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, result.expectedType, token.Type)
		}

		if token.Literal != result.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, result.expectedLiteral, token.Literal)
		}

	}
}
