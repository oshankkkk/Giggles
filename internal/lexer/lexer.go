package lexer

import (
	"bufio"
	"fmt"
	"os"
	//"fmt"
	//"strings"
)
func ReadFile(file *os.File){
	scanner := bufio.NewScanner(file)
	var tokenlist [][]Token
	for scanner.Scan() {
		token:=lexer(scanner.Text())
		tokenlist = append(tokenlist, token)

	fmt.Println(scanner.Text())
	}
	fmt.Println(tokenlist)
}

//lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).
type Token struct{
	ID int
	Type string
	Value rune

}
func lexer(source string)[]Token{
	var tokenlist []Token
	for index,value:=range source{
		switch string(value) {
		case "(":
		tokenlist = append(tokenlist, 	Token{Type: "LParam", Value: value,ID: index})
		case ")":
			tokenlist = append(tokenlist, Token{Type: "RParan", Value: value, ID: index})
		case ",":
			tokenlist=append(tokenlist, Token{Type: "Comma", Value: value,ID: index})
		case "\"":
			tokenlist = append(tokenlist, Token{Type: "DQuotation",Value:value, ID: index })
		case " ":
			tokenlist = append(tokenlist,Token{Type: "WhiteSpace",Value:value, ID: index })
//		default:
//			tokenlist = append(tokenlist, Token{Type: "None",Value:value, ID: index })
		}	
	}
	return tokenlist
}


