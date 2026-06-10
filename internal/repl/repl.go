package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"lang/internal/frontend/parser"
	"lang/internal/frontend/lexer"
	"lang/internal/backend/compiler"
	"lang/internal/backend/vm"
	
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
		var lexer lexer.Lexer
		lexer.ReadLine(input)
		var parser parser.Parser
		ast:=parser.Run(&lexer)
		bytecodelist:=compiler.Compile(ast)
		bytearray,constTable,vartable:=compiler.ToBytecode(bytecodelist)	
		ans := vm.Machine(bytearray, constTable, vartable, stack, stackpointer, heap)
		fmt.Println(ans)

	}
}
