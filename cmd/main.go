package main

import (
	"fmt"
	"os"
	"lang/internal/lexer"
	//"lang/internal/parser"
)
func main(){
file,err:=os.Open("myfile.txt")
if err!=nil{
	fmt.Println(err)
}
	tokenlist:=lexer.ReadFile(file)
	fmt.Println(tokenlist)
//	parser.Parser()	

	fmt.Println("end of program")


}


