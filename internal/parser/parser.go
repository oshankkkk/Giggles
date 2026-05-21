package parser

import (
	"lang/internal/lexer"
//	"mime/multipart"
	"os"
)
func Parser(file *os.File){
	tokenlist:=lexer.ReadFile(file)
	for _,value:=range tokenlist{
	//each statemens
	addsubparser(value[0])
	}
}

type Expression interface{
	expression()
}
type Literal struct{
	value lexer.Token
}
type Groups struct{
	group Expression
}
type Binary struct{
	left Expression
	right Expression
	operator lexer.TokenType
}
type Unary struct{
	single Expression
}
func (s*Literal) expression()
func (s*Groups) expression()
func (s*Unary) expression()
func (s*Binary) expression()

func muldivparser(tokenlist[]lexer.Token,index int)(Expression,int){
	left,index:=numgroupparser(tokenlist,index)
	for index<len(tokenlist){
		char:=tokenlist[index]
		if !(char.Type==lexer.STAR|| char.Type==lexer.SLASH){
			break
		}
		index++
		right,index=numgroupparser(tokenlist,index)
		exp:=Binray{
			left:left,
			operator:char,
			right:right,
		}
	} 
	return left,index

}

func numgroupparser(tokenlist []lexer.Token,index int){
 
}

func walker(statement []lexer.Token){
	for index,value :=range statement{

	}
}
