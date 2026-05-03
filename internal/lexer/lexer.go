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
//	fmt.Println(scanner.Text())
	}
	fmt.Println(tokenlist)
}
type TokenType string

const (
	LEFT_PAREN    TokenType = "("
	RIGHT_PAREN   TokenType = ")"
	LEFT_BRACE    TokenType = "{"
	RIGHT_BRACE   TokenType = "}"
	LEFT_BRACKET  TokenType = "["
	RIGHT_BRACKET TokenType = "]"
	COMMA         TokenType = ","
	DOT           TokenType = "."
	SEMICOLON     TokenType = ";"
	COLON         TokenType = ":"

	PLUS    TokenType = "+"
	MINUS   TokenType = "-"
	STAR    TokenType = "*"
	SLASH   TokenType = "/"
	PERCENT TokenType = "%"

	EQUAL	TokenType = "="
	NOT		TokenType = "!"
	LESS    TokenType = "<"
	GREATER TokenType = ">"
	SINGLE_QUOTES TokenType = "'"
	DOUBLE_QUOTES TokenType = "\""

	EQUAL_EQUAL   TokenType = "=="
	NOT_EQUAL    TokenType = "!="
	LESS_EQUAL    TokenType = "<="
	GREATER_EQUAL TokenType = ">="

)



//lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).
type Token struct{
	ID int
	Type TokenType 
	Value string

}


func lexer(source string)[]Token{
	singleCharTokens := map[string]TokenType{
		string(LEFT_PAREN):    LEFT_PAREN,
		string(RIGHT_PAREN):   RIGHT_PAREN,
		string(LEFT_BRACE):    LEFT_BRACE,
		string(RIGHT_BRACE):   RIGHT_BRACE,
		string(LEFT_BRACKET):  LEFT_BRACKET,
		string(RIGHT_BRACKET): RIGHT_BRACKET,
		string(COMMA):         COMMA,
		string(DOT):           DOT,
		string(SEMICOLON):     SEMICOLON,
		string(COLON):         COLON,
		string(PLUS):          PLUS,
		string(MINUS):         MINUS,
		string(STAR):          STAR,
		string(SLASH):         SLASH,
		string(PERCENT):       PERCENT,

		string(GREATER): GREATER,
		string(LESS): LESS, 
		string(NOT): NOT,
		string(EQUAL): EQUAL,
	}
	doubleCharTokens:=map[string]TokenType{
		//string(SINGLE_QUOTES): SINGLE_QUOTES,
		//string(DOUBLE_QUOTES): DOUBLE_QUOTES,
		string(EQUAL_EQUAL):EQUAL_EQUAL,
		string(NOT_EQUAL):NOT_EQUAL,
		string(LESS_EQUAL):LESS_EQUAL,
		string(GREATER_EQUAL):GREATER_EQUAL,
	}
	var tokenlist []Token
	cont:=false
	for index,value:=range source{
		character:=string(value)
		if cont{
			newToken:=tokenlist[len(tokenlist)-1].Value+character
			if tokenType,ok:=doubleCharTokens[newToken];ok{
			tokenlist[len(tokenlist)-1].Type=tokenType	
			tokenlist[len(tokenlist)-1].Value=newToken
			continue
			}
		}
		tokenType, ok := singleCharTokens[character] 
		if ok {
			tokenlist = append(tokenlist, Token{
				Type:  tokenType,
				Value: character,
				ID:    index,
			})
			if slices.Contains([]string{string(GREATER),string(LESS),string(NOT),string(EQUAL)},character){
				cont=true
			}
		}

		// if equal sign then if 	
	}


return tokenlist
}


