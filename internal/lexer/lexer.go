package lexer
//lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).
import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(file *os.File)[][]Token{
	scanner := bufio.NewScanner(file)
	var tokenlist [][]Token
	for scanner.Scan() {
		tokens:=lexer(scanner.Text())
		tokenlist = append(tokenlist, tokens)
		fmt.Println(scanner.Text())
	}
	return tokenlist
}

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

func lexer(source string)[]Token {

	pointer:=0
	var tokenlist []Token
	move := func() {
		pointer++
	}
	current := func() byte {
		if pointer < len(source) {
			return source[pointer]
		}
		return 0
	}
	next := func() byte {
		if pointer+1 < len(source) {
			return source[pointer+1]
		}
		return 0
	}

	for pointer<len(source){
		char:=source[pointer]		

		//whitespace	
		if char == ' ' || char == '\t' || char == '\r' || char == '\n' {
			move()
			continue
		}

		// Strings " or '
		if char == '"' || char == '\'' {
			quote := char
			move() // skip opening quote
			start := pointer
			for pointer < len(source) && current() != quote {
				move()
			}
			str := source[start:pointer]
			move() // skip closing quote
			addToken(STRING, str,&tokenlist)
			continue
		}
		// numbers
		if isDigit(char) {
			start := pointer
			for pointer < len(source) && isDigit(current()) {
				move()
			}
			// Handle decimal
			if pointer < len(source) && current() == '.' && isDigit(next()) {
				move() // consume '.'
				for pointer < len(source) && isDigit(current()) {
					move()
				}
			}
			addToken(NUMBER, source[start:pointer],&tokenlist)
			continue
		}


		if pointer+1 < len(source) {
			two := string(source[pointer : pointer+2])
			if tt, ok := doubleCharTokens[two]; ok {
				addToken(tt, two,&tokenlist)
				move()
				move()
				continue
			}
		}

		if tt, ok := singleCharTokens[string(char)]; ok {
			addToken(tt, string(char),&tokenlist)
			move()
			continue
		}
		move()
	} 
	return tokenlist
}

func addToken(t TokenType, val string,tokenlist *[]Token) {
	fmt.Println("adding")
	*tokenlist= append(*tokenlist, Token{ID: idCounter, Type: t, Value: val})
	idCounter++
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
