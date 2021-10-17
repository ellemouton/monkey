package lexer

import (
	"testing"

	"github.com/ellemouton/monkey/token"
)

func TestNextToken1(t *testing.T) {
	type res struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tests := []struct {
		name        string
		input       string
		expectedRes []res
	}{
		{
			name:  "1",
			input: `=+(){},;`,
			expectedRes: []res{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "2",
			input: `!-/*5;
				5 < 10 > 5;`,
			expectedRes: []res{
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
			},
		},
		{
			name: "3",
			input: `if (5 < 10) {
       					return true;
				} else {
       					return false;
				}`,
			expectedRes: []res{
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			name: "4",
			input: `10 == 10;
				10 != 9;`,
			expectedRes: []res{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "5",
			input: `let five = 5;
                                let ten = 10;

                                let add = fn(x, y) {
                                        x + y;
                                };

                                let result = add(five, ten);
                                `,
			expectedRes: []res{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
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
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "6",
			input: `"foobar"
				"foo bar"`,
			expectedRes: []res{
				{token.STRING, "foobar"},
				{token.STRING, "foo bar"},
				{token.EOF, ""},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(test.input)

			for i, tt := range test.expectedRes {
				tok := l.NextToken()

				if tok.Type != tt.expectedType {
					t.Fatalf("tests[%d] - tokentype "+
						"wrong. expected=%q, got %q",
						i, tt.expectedType, tok.Type)
				}

				if tok.Literal != tt.expectedLiteral {
					t.Fatalf("tests[%d] - literal "+
						"wrong. expected=%q, got %q",
						i, tt.expectedLiteral,
						tok.Literal)
				}
			}
		})
	}
}
