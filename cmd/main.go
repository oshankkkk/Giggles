package main

import (
	"fmt"
	//"os"
	//"lang/internal/lexer"
	"lang/internal/parser"
)
func main(){
	//fmt.Println(+2+2)
	//file,err:=os.Open("../myfile.txt")
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//tokenlist:=lexer.ReadFile(file)
	parser.Parser()	
	fmt.Println("end of program")


}


