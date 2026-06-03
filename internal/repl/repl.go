package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"lang/internal/lexer"
	"lang/internal/compiler"
	"lang/internal/parser"
	"lang/internal/vm"
	
)
func check(err error){
	if err!=nil{
		fmt.Printf("err %s",err)
	}
}
func Run(stack *[]int,stackpointer *int,heap *map[string]int){
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
		ast:=parser.Parser(tokens)
		bytecodelist:=compiler.Compile(ast)
		bytearray,constTable,vartable:=vm.ToBytecode(bytecodelist)	
		ans := vm.Machine(bytearray, constTable, vartable, stack, stackpointer, heap)
		fmt.Println(ans)

	}
}
