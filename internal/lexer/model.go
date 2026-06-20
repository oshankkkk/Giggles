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
	ILLEGAL TokenType = "ILLEGAL"
	WHITESPACE TokenType = "WHITESPACE"
	EOF        TokenType = "EOF"

//	LET    TokenType = "LET"
	TYPEDEFF TokenType = "TYPEDEFF"
	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"
	THEN   TokenType = "THEN"
	END    TokenType = "END"
	FN   TokenType = "FN"
	LOCAL  TokenType = "LOCAL"
	RETURN TokenType = "RETURN"
	MUT TokenType = "MUT"
	MAIN TokenType = "MAIN"

//	WHILE TokenType = "WHILE"
	FOR   TokenType = "FOR"
	BREAK TokenType = "BREAK"
	CONTINUE TokenType = "CONTINUE"

	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
	NIL   TokenType = "NIL"
	NUMBER TokenType = "NUMBER" 
	AND TokenType = "AND"
	OR TokenType = "OR"
)
type Token struct{
	ID int
	Type TokenType 
	Value string
	Line int
	Column int
	StartIndex int
	EndIndex int
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
	"&&": AND,
	"||":OR,

}
var keywordTokens=map[string]TokenType{
	"if":       IF,
	"else":     ELSE,
	"then":     THEN,
	"end":      END,
	"fn":     FN,
	"local":    LOCAL,
	"return":   RETURN,

	"int": TYPEDEFF,
	"string": TYPEDEFF,
	"double": TYPEDEFF,
	"bool": TYPEDEFF,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"main": MAIN,

	"nil":      NIL,
	"true": TRUE,
	"false":FALSE,
	"mut": MUT,

}


