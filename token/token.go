package token

type TokenType string

// Token represents a single token. In production this would probably
// include filename and line number to track down lexer and parsing errors
// and Type would probably something that allows for better performance, like
// bytes.
type Token struct {
	// Type tells us if this is an integer, a right bracket, an identifier,
	// a keyword etc.
	Type TokenType

	// Literal is the literal value of the token.
	Literal string
}

const (
	// Special types
	ILLEGAL = "ILLEGAL" // a token/char we dont know about
	EOF     = "EOF"     // end of file. This tells our parser later on that it can stop

	// Identifiers & literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   //  12345

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERICK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Deliminators
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
