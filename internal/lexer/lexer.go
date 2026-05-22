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
