package lexer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	//"strings"
	//"encoding"
	//"strings"
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
)

//lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).
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

	"'":  SINGLE_QUOTES,
	"\"": DOUBLE_QUOTES,
}


var doubleCharTokens = map[string]TokenType{
	"==": EQUAL_EQUAL,
	"!=": NOT_EQUAL,
	"<=": LESS_EQUAL,
	">=": GREATER_EQUAL,
}
	var tokenlist []Token
	doubletoken:=false
	stringtoken:=false
	for index,value:=range source{
		character:=string(value)
		if doubletoken{
			newToken:=tokenlist[len(tokenlist)-1].Value+character
			if tokenType,ok:=doubleCharTokens[newToken];ok{
			tokenlist[len(tokenlist)-1].Type=tokenType	
			tokenlist[len(tokenlist)-1].Value=newToken
			continue
			}
		}
		if stringtoken{

		}
		tokenType, ok := singleCharTokens[character] 
		if ok {
	if slices.Contains([]TokenType{SINGLE_QUOTES,DOUBLE_QUOTES},tokenType){
				if !stringtoken{
					stringtoken=true
					tokenType = STRING
			tokenlist = append(tokenlist, Token{
				Type:  tokenType,
				Value: character,
				ID:    index,
			})
	
				}else{stringtoken=false
					continue
			}
			}
			if !stringtoken{
			if slices.Contains([]TokenType{GREATER,LESS,NOT,EQUAL},tokenType){
				doubletoken=true
			}
		

			tokenlist = append(tokenlist, Token{
				Type:  tokenType,
				Value: character,
				ID:    index,
			})
	
			
		}
	}	

		// if equal sign then if 	
	}


return tokenlist
}


