package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // sera utilizado/chamado quando algum caracter ou algo que nao tenha sido identificado seja utilizado
	EOF     = "EOF"     // end of file

	// identificadores e literals
	IDENT = "IDENT"
	INT   = "INT"
	STRING = "STRING"

 // operafores
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// delimitadores
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

type Token struct {
	Type    TokenType
	Literal string
}

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
