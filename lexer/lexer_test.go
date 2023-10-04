package lexer_test

import (
	"testing"

	"monkey/lexer"
	"monkey/token"
)

type testExpect struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []testExpect{
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

	l := lexer.New(input)
	assertInputTokens(t, l, tests)
}

func TestNextToken2(t *testing.T) {
	input := `let four = 4;
	let eleven = 11;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(four, eleven);
	`

	tests := []testExpect{
		{token.LET, "let"},
		{token.IDENT, "four"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "eleven"},
		{token.ASSIGN, "="},
		{token.INT, "11"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "four"},
		{token.COMMA, ","},
		{token.IDENT, "eleven"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	assertInputTokens(t, l, tests)
}

func TestNextTokenKeywords(t *testing.T) {
	input := `let t = true;
	let f = false;

	if (t) {
		return t;
	} else {
		return f;
	}
	`

	tests := []testExpect{
		{token.LET, "let"},
		{token.IDENT, "t"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "f"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "t"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "t"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "f"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	l := lexer.New(input)

	assertInputTokens(t, l, tests)
}

func TestNextTokenOperators(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
	`
	tests := []testExpect{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := lexer.New(input)

	assertInputTokens(t, l, tests)
}

func TestNextTokenEqNotEq(t *testing.T) {
	input := `10 == 10;
	10 != 9;
	`

	tests := []testExpect{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	assertInputTokens(t, l, tests)
}

func assertInputTokens(t *testing.T, lexer *lexer.Lexer, expectations []testExpect) {
	for i, tt := range expectations {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - token type wrong. expexted=%q, got= %q",
				i, tt.expectedType, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expexted=%q, got= %q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
