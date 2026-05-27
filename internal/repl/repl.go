package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"lang/internal/lexer"
	"lang/internal/parser"
	
)
func check(err error){
	if err!=nil{
		fmt.Printf("err %s",err)
	}
}
func Run(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interactive Shell (type 'exit' to quit)")
	for {
		fmt.Print(">> ")		
		input, err := reader.ReadString('\n')
		check(err)
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}
		tokens:=lexer.Readline(input)
		parser.Parser(tokens)
		
	}
}
