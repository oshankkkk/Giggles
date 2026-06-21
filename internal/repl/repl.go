package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"lang/internal/parser"
	"lang/internal/lexer"
	"lang/internal/compiler"
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
		var lexer lexer.Lexer
		lexer.ReadLine(input)
		var parser parser.Parser
		ast:=parser.Run(&lexer)
		var comp compiler.State
		var vm vm.GVM
		globalscope:=compiler.InitScope()
		comp.ToBytes(ast,globalscope)
		//bytearray,constTable,vartable:=compiler.(bytecodelist)	
		ans := vm.Machine(comp.Buff, comp.CounterTable)
		fmt.Println(ans)

	}
}
