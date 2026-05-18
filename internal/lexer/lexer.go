package lexer
//lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).
import (
	"bufio"
	"fmt"
	"os"
	"slices"
)
func ReadFile(file *os.File){
	scanner := bufio.NewScanner(file)
	var tokenlist [][]Token
	for scanner.Scan() {
		tokens:=lexer(scanner.Text())
		tokenlist = append(tokenlist, tokens)
		fmt.Println(scanner.Text())
	}
	fmt.Println(tokenlist)
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

)


type Token struct{
	ID int
	Type TokenType 
	Value string
}

func lexer(source string)[]Token{

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
	//Scope highlighting plugin
	var tokenlist []Token
	stringtoken:=false
	newlist:=""
	for index,value:=range source{
		character:=string(value)
		tokenType, singleCharToken:= singleCharTokens[character]

		if !(singleCharToken){
			newlist+=character
			continue
		}

		newstringtoken:=newlist
		if len(newlist)>1 && !stringtoken{

			newstringtoken=newlist[:len(newlist)-1]}

			//tokenising the seperated
			if slices.Contains([]TokenType{SINGLE_QUOTES,DOUBLE_QUOTES},tokenType){
				if !stringtoken{
					stringtoken=true
				}else{
					stringtoken=false
					tokenlist = append(tokenlist, Token{
						Type: STRING,
						Value: newstringtoken,
						ID:    index,
					})
				}
			}else
			if tokenType,ok:=doubleCharTokens[ newstringtoken];ok{
				tokenlist = append(tokenlist, Token{
					Type:  tokenType,
					Value: newlist,
					ID:    index,
				})


			}else if tokenType,ok:=keywordTokens[newstringtoken]; ok{
				tokenlist = append(tokenlist, Token{
					Type:  tokenType,
					Value: newlist,
					ID:    index,
				})

			}else if stringtoken{
				continue
			}else {
				tokenlist = append(tokenlist, Token{
					Type: IDENTIFIER,
					Value: newlist,
					ID:    index,
				})

			}

			//tokenising the seperator
			tokenlist = append(tokenlist, Token{
				Type:  tokenType,
				Value: character,
				ID:    index,
			})

			newlist=""
		}

		return tokenlist
	}


