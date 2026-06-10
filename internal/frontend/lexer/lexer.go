package lexer

import (
	"os"
	"fmt"
)

type Lexer struct {
	pointer int
	source string
	line int
	column int
	idCounter int
}

func (l *Lexer) ReadFile(filename string){
	content,err:=os.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}
	l.source=string(content)
}

func (l *Lexer) ReadLine(line string){
	l.source=line
}

func (l *Lexer) move (){
		l.column++
		l.pointer++
}

func (l *Lexer) current() byte {
	if l.pointer < len(l.source) {
		return l.source[l.pointer]
	}
	return 0
}

func (l *Lexer)peek() byte {
	if l.pointer+1 < len(l.source) {
		return l.source[l.pointer+1]
	}
	return 0
}


func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l *Lexer) NextToken() Token {
	for l.pointer < len(l.source) {
		char := l.current()

		// whitespace
		if char == ' ' || char == '\t' || char == '\r' || char == '\n' {
			if char == '\n' {
				l.line++
				l.column = 0
			} else {
				l.column++
			}
			l.move()
			continue
		}

		// single char tokens
		if tt, ok := singleCharTokens[string(char)]; ok {
			tok := Token{
				ID:    l.idCounter,
				Type:  tt,
				Value: string(char),
				Line:  l.line,
				Column: l.column,
			}
			l.idCounter++
			l.move()
			return tok
		}

		// strings
		if char == '"' || char == '\'' {
			quote := char
			startCol := l.column

			l.move() // skip opening quote
			start := l.pointer

			for l.pointer < len(l.source) && l.current() != quote {
				if l.current() == '\n' {
					l.line++
					l.column = 0
				} else {
					l.column++
				}
				l.move()
			}

			str := l.source[start:l.pointer]
			l.move() // skip closing quote

			tok := Token{
				ID:     l.idCounter,
				Type:   STRING,
				Value:  str,
				Line:   l.line,
				Column: startCol,
			}
			l.idCounter++
			return tok
		}

		// numbers
		if isDigit(char) {
			start := l.pointer
			startCol := l.column

			for l.pointer < len(l.source) && isDigit(l.current()) {
				l.move()
			}

			// decimal part
			if l.pointer < len(l.source) && l.current() == '.' &&
				l.peek() != 0 && isDigit(l.peek()) {

				l.move() // consume '.'

				for l.pointer < len(l.source) && isDigit(l.current()) {
					l.move()
				}
			}

			tok := Token{
				ID:     l.idCounter,
				Type:   NUMBER,
				Value:  l.source[start:l.pointer],
				Line:   l.line,
				Column: startCol,
			}
			l.idCounter++
			return tok
		}

		// double char tokens
		if l.pointer+1 < len(l.source) {

			fmt.Println("hi")
			two := l.source[l.pointer : l.pointer+2]

			if tt, ok := doubleCharTokens[two]; ok {
				startCol := l.column

				l.move()
				l.move()

				tok := Token{
					ID:     l.idCounter,
					Type:   tt,
					Value:  two,
					Line:   l.line,
					Column: startCol,
				}
				l.idCounter++
				return tok
			}
		}

		// identifiers / keywords
		if isAlpha(char) {
			start := l.pointer
			startCol := l.column

			for l.pointer < len(l.source) &&
				(isAlpha(l.current()) || isDigit(l.current()) || l.current() == '_') {
				l.move()
			}

			word := l.source[start:l.pointer]

			var tokType TokenType
			if tt, ok := keywordTokens[word]; ok {
				tokType = tt
			} else {
				tokType = IDENTIFIER
			}

			tok := Token{
				ID:     l.idCounter,
				Type:   tokType,
				Value:  word,
				Line:   l.line,
				Column: startCol,
			}
			l.idCounter++
			return tok
		}

		// unknown character fallback
		startCol := l.column
		l.move()

		tok := Token{
			ID:     l.idCounter,
			Type:   ILLEGAL,
			Value:  string(char),
			Line:   l.line,
			Column: startCol,
		}
		l.idCounter++
		return tok
	}

	return Token{
		ID:    l.idCounter,
		Type:  EOF,
		Value: "",
		Line:  l.line,
		Column: l.column,
	}
}
