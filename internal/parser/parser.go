package parser

import (
	"lang/internal/lexer"
	"fmt"
)
func Parser(){
	//file,err:=os.open("myfile.txt")
	//if err!=nil{
	//	fmt.println(err)
	//	fmt.println("uuy")
	//}

//	tokenlist:=lexer.ReadFile(file)
tokenlist := []lexer.Token{
    {ID: 0, Type: lexer.NUMBER, Value: "2"},
    {ID: 1, Type: lexer.PLUS,   Value: "+"},
    {ID: 2, Type: lexer.NUMBER, Value: "2"},
    {ID: 3, Type: lexer.STAR,   Value: "*"},
    {ID: 4, Type: lexer.NUMBER, Value: "3"},
}
	//fmt.Println(tokenlist)
	exp,_:=addsubparser(tokenlist,0)
	Dw(exp)
}

func Dw(ex Expression){
	ex.expression()
	if _, ok := ex.(Literal); ok {
		return
	}
	if bin, ok := ex.(Binary); ok {
		Dw(bin.left)
		Dw(bin.right)
		return
	}
	if grp, ok := ex.(Groups); ok {
		Dw(grp.value)
	}
}

type Expression interface{
	expression()
}
type Literal struct{

	nodeName string
	value lexer.Token
}
type Groups struct{

	nodeName string
	value Expression
}
type Binary struct{
	nodeName string
	left Expression
	right Expression
	operator lexer.TokenType
}
type Unary struct{
	value Expression
}
func (s Literal) expression(){
fmt.Println(s.nodeName)
}
func (s Groups) expression(){

fmt.Println(s.nodeName)
}
func (s Unary) expression(){
//fmt.Println(s.nodeName)
}
func (s Binary) expression(){
fmt.Println(s.nodeName)}

func addsubparser(tokenlist[]lexer.Token,index int)(Expression,int){
	left,index:=muldivparser(tokenlist,index)						
	var exp,right Expression
	for index<len(tokenlist){
		char:=tokenlist[index]
		if !(char.Type==lexer.PLUS||char.Type==lexer.MINUS){
			//if equals break
			break	
		}
		index++
		right,index=muldivparser(tokenlist,index)
		exp=Binary{
			nodeName: "binary-addsub",
			left:left,
			operator:char.Type,
			right:right,
		}
	}
	if exp == nil {
		return left, index
	}
	return exp,index
}

func muldivparser(tokenlist[]lexer.Token,index int)(Expression,int){
	var right,exp Expression
	left,index:=numgroupparser(tokenlist,index)
	for index<len(tokenlist){
		char:=tokenlist[index]
		if !(char.Type==lexer.STAR|| char.Type==lexer.SLASH){
			break
		}
		index++
		right,index=numgroupparser(tokenlist,index)
		exp=Binary{
			nodeName: "binary-muldiv",
			left:left,
			operator:char.Type,
			right:right,
		}
	}
	if exp == nil {
		return left, index
	}
	return exp,index

}
func numgroupparser(tokenList []lexer.Token, index int) (Expression, int) {
    char := tokenList[index]
    if char.Type == lexer.NUMBER{
		return Literal{nodeName:"lit",value: char}, index+1
    }
	exp,index:=addsubparser(tokenList, index)
	return Groups{nodeName:"bracket",value: exp}, index
}

