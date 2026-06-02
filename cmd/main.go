package main

import (
	"fmt"
	"os"
	"lang/internal/lexer"
	"lang/internal/parser"
	"lang/internal/repl"
	"strings"
	"lang/internal/compiler"
	"lang/internal/vm"

)

func main(){
	args:=os.Args[1:]
	if len(args)>0 {
		readscript(args[0])
	}else{
	repl.Run()
	}
}

func readscript(path string){
	file,err:=os.Open(path)
	if err!=nil{
		fmt.Println(err)
	}

	tokenlist:=lexer.ReadFile(file)
	rootnode:=parser.Parser(tokenlist)	
	prettyprinter(rootnode,0)
		bytecodelist:=compiler.Compile(rootnode)
		bytearray,constTable,vartable:=vm.ToBytecode(bytecodelist)	
		ans:=vm.Machine(bytearray,constTable,vartable)	
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
	default:
		fmt.Printf("%s???\n", pad)
	}
}
