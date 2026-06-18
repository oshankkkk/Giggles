package main

import (
	"fmt"
	"os"
	"lang/internal/repl"
	"strings"
	"lang/internal/compiler"
	"lang/internal/lexer"
	"lang/internal/parser"
	"lang/internal/vm"
)

  	var stack = make([]int, 1024)
    var stackpointer int
	var heap=make(map[string]int)



func main(){	
	args:=os.Args[1:]
	if len(args)>0 {
		readscript(args[0])
	}else{
		 repl.Run(&stack, &stackpointer, &heap)
	}
}

func readscript(path string){
	var lexer lexer.Lexer
	lexer.ReadFile(path)
	var parser parser.Parser
	rootnode:=parser.Run(&lexer)	
	prettyprinter(rootnode,0)
		globalscope:=compiler.InitScope()
		var vm vm.GVM
		var code compiler.State
		bytecodelist:=code.ToBytes(rootnode,globalscope)
		//bytearray,constTable,vartable:=compiler.(bytecodelist)	
		ans := vm.Machine(bytecodelist,code.CounterTable)
		fmt.Println(ans)


	fmt.Println("end of program")

}

func prettyprinter(ex parser.ASTNode, indent int) {
	pad := strings.Repeat("  ", indent)
	switch n := ex.(type) {
	case parser.Program:
		fmt.Printf("%sProgram\n", pad)
		for _, s := range n.Statements {
			prettyprinter(s, indent+1)
		}

	case parser.VarDecl:
		fmt.Printf("%sVarDecl(%s)\n", pad, n.Name.Value)
		prettyprinter(n.Value, indent+1)
	case parser.ExprStatement:
		fmt.Printf("%sExprStatement\n", pad)
		prettyprinter(n.Expr, indent+1)
	case parser.Literal:
		fmt.Printf("%sLiteral(%s)\n", pad, n.Value.Value)
	case parser.Identifier:
		fmt.Printf("%sIdentifier(%s)\n", pad, n.Name.Value)
	case parser.Binary:
		fmt.Printf("%sBinary(%s)\n", pad, n.Operator)
		prettyprinter(n.Left, indent+1)
		prettyprinter(n.Right, indent+1)
	case parser.Groups:
		fmt.Printf("%sGroup\n", pad)
		prettyprinter(n.Value, indent+1)
	case parser.Unary:
		fmt.Printf("%sUnary\n", pad)
		prettyprinter(n.Value, indent+1)
	case parser.Condition:
		fmt.Printf("%sCondition\n", pad)
		prettyprinter(n.Condition, indent+1)
		for _, s := range n.Result {
			prettyprinter(s, indent+1)
		}
		for _, s := range n.ElseResult {
			prettyprinter(s, indent+1)
		}

	default:
		fmt.Printf("%s???\n", pad)
	}
}
