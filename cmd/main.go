package main

import (
	"fmt"
	"os"
	"lang/internal/lexer"
)
func main(){
	fmt.Println(+2+2)
	file,err:=os.Open("../myfile.txt")
	if err!=nil{
		fmt.Println(err)
	}
	lexer.ReadFile(file)
	fmt.Println("end of program")


}


