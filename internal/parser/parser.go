package parser

import (
	"lang/internal/lexer"
	"fmt"
	"strings"
)
func Parser(tokenlist [][]lexer.Token){
//	tokenlist:=lexer.ReadFile(file)
//tokenlist := []lexer.Token{
//    {ID: 0, Type: lexer.NUMBER, Value: "2"},
//    {ID: 1, Type: lexer.PLUS,   Value: "+"},
//    {ID: 2, Type: lexer.NUMBER, Value: "2"},
//    {ID: 3, Type: lexer.STAR,   Value: "*"},
//    {ID: 4, Type: lexer.NUMBER, Value: "3"},
//}
pointer:=0

//[[{0 LEFT_PAREN (} {1 NUMBER 1} {2 PLUS +} {3 NUMBER 2} {4 PLUS +} {5 NUMBER 3} {6 RIGHT_PAREN )} {7 STAR *} {8 NUMBER 12}] []]
	exp:=addsubparser(tokenlist[0],&pointer)
	Pp(exp,0)
}

func Pp(ex Expression, indent int) {
    pad := strings.Repeat("  ", indent)
    switch n := ex.(type) {
    case Literal:
        fmt.Printf("%sLiteral(%s)\n", pad, n.value.Value)
    case Binary:
        fmt.Printf("%sBinary(%s)\n", pad, n.operator)
        Pp(n.left, indent+1)
        Pp(n.right, indent+1)
    case Groups:
        fmt.Printf("%sGroup\n", pad)
        Pp(n.value, indent+1)
    case Unary:
        fmt.Printf("%sUnary\n", pad)
        Pp(n.value, indent+1)
    default:
        fmt.Printf("%s???\n", pad)
    }
}

func addsubparser(tokenlist[]lexer.Token,pointer *int)(Expression){
	left:=muldivparser(tokenlist,pointer)						
	var exp,right Expression
	for *pointer<len(tokenlist){
		char:=tokenlist[*pointer]
		if !(char.Type==lexer.PLUS||char.Type==lexer.MINUS){
			//if equals break
			break	
		}
		*pointer++
		right=muldivparser(tokenlist,pointer)
		exp=Binary{
			nodeName: "binary-addsub",
			left:left,
			operator:char.Type,
			right:right,
		}
	}
	if exp == nil {
		return left
	}
	return exp
}

func muldivparser(tokenlist[]lexer.Token, pointer *int)(Expression){
	var right,exp Expression
	left:=numgroupparser(tokenlist,pointer)
	for *pointer<len(tokenlist){
		char:=tokenlist[*pointer]
		if !(char.Type==lexer.STAR|| char.Type==lexer.SLASH){
			break
		}
		*pointer++
		right =numgroupparser(tokenlist,pointer)
		exp=Binary{
			nodeName: "binary-muldiv",
			left:left,
			operator:char.Type,
			right:right,
		}
	}
	if exp == nil {
		return left
	}
	return exp

}
func numgroupparser(tokenList []lexer.Token, pointer *int) (Expression) {
    char := tokenList[*pointer]
    if char.Type == lexer.NUMBER{
		*pointer++
		return Literal{nodeName:"lit",value: char} 
    }
	*pointer++
	exp:=addsubparser(tokenList, pointer)
	return Groups{nodeName:"bracket",value: exp}
}

