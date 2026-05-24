package lexer
type TokenType string

const (
	LEFT_PAREN    TokenType = "LEFT_PAREN"
	RIGHT_PAREN   TokenType = "RIGHT_PAREN"
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	COMMA         TokenType = "COMMA"
	DOT           TokenType = "DOT"
	SEMICOLON     TokenType = "SEMICOLON"
	COLON         TokenType = "COLON"

	PLUS    TokenType = "PLUS"
	MINUS   TokenType = "MINUS"
	STAR    TokenType = "STAR"
	SLASH   TokenType = "SLASH"
	PERCENT TokenType = "PERCENT"

	EQUAL         TokenType = "EQUAL"
	NOT           TokenType = "NOT"
	LESS          TokenType = "LESS"
	GREATER       TokenType = "GREATER"
	SINGLE_QUOTES TokenType = "SINGLE_QUOTES"
	DOUBLE_QUOTES TokenType = "DOUBLE_QUOTES"
	STRING TokenType = "STRING"

	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	NOT_EQUAL     TokenType = "NOT_EQUAL"
	LESS_EQUAL    TokenType = "LESS_EQUAL"
	GREATER_EQUAL TokenType = "GREATER_EQUAL"
	IDENTIFIER TokenType = "IDENTIFIER"
	WHITESPACE TokenType ="WHITESPACE"

	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"
	THEN   TokenType = "THEN"
	END    TokenType = "END"
	FUNC   TokenType = "FUNC"
	LOCAL  TokenType = "LOCAL"
	RETURN TokenType = "RETURN"

	WHILE TokenType = "WHILE"
	FOR   TokenType = "FOR"
	BREAK TokenType = "BREAK"
	CONTINUE TokenType = "CONTINUE"

	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
	NIL   TokenType = "NIL"
	NUMBER TokenType = "NUMBER" 
)

type Token struct{
	ID int
	Type TokenType 
	Value string
	Line int
	Column int
}
var idCounter int
var singleCharTokens = map[string]TokenType{
	"(": LEFT_PAREN,
	")": RIGHT_PAREN,
	"{": LEFT_BRACE,
	"}": RIGHT_BRACE,
	"[": LEFT_BRACKET,
	"]": RIGHT_BRACKET,
	",": COMMA,
	".": DOT,
	";": SEMICOLON,
	":": COLON,

	"+": PLUS,
	"-": MINUS,
	"*": STAR,
	"/": SLASH,
	"%": PERCENT,

	">": GREATER,
	"<": LESS,
	"!": NOT,
	"=": EQUAL,

	" ": WHITESPACE,
	"'":  SINGLE_QUOTES,
	"\"": DOUBLE_QUOTES,
}


var doubleCharTokens = map[string]TokenType{
	"==": EQUAL_EQUAL,
	"!=": NOT_EQUAL,
	"<=": LESS_EQUAL,
	">=": GREATER_EQUAL,

}
var keywordTokens=map[string]TokenType{
	"if":       IF,
	"else":     ELSE,
	"then":     THEN,
	"end":      END,
	"func":     FUNC,
	"local":    LOCAL,
	"return":   RETURN,

	"while":    WHILE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,

	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,

}


