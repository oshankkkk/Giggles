package main

import (
	"fmt"
	"os"
	"lang/internal/lexer"
	"lang/internal/parser"
)

func main(){
file,err:=os.Open("lol.lol")
if err!=nil{
	fmt.Println(err)
}

tokenlist:=lexer.ReadFile(file)
fmt.Println(tokenlist)
	parser.Parser(tokenlist)	

	fmt.Println("end of program")
	

}


