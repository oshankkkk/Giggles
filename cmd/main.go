package main

import (
	"fmt"
	"os"
	"lang/internal/lexer"
	"lang/internal/parser"
	"lang/internal/repl"
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
	parser.Parser(tokenlist)	

	fmt.Println("end of program")
}
