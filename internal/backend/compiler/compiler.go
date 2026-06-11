package compiler

import (
	"lang/internal/frontend/lexer"
	"lang/internal/frontend/parser"
)
var bytecode []byte

// Binary, literall 
func Compile(ast parser.ASTNode)[]string{
	var opcode string
	list:=[]string{}
	if value,ok:=ast.(parser.Program);ok{
		for _,val:=range value.Statements{
			list = append(list, Compile(val)...)
		}		
	}

	if value,ok:=ast.(parser.ExprStatement);ok{
		list = append(list, Compile(value.Expr)...)
	}

	if value,ok:=ast.(parser.Groups);ok{
		list = append(list, Compile(value.Value)...)

	}
	if value,ok:=ast.(parser.VarDecl);ok{
		list = append(list, Compile(value.Value)...)
		list = append(list, "VAR_DEC",value.Name.Value)

	}
	if value,ok:=ast.(parser.Identifier);ok{
		list = append(list, "VAR",value.Name.Value)

	}
	if value,ok:=ast.(parser.Literal);ok{
		//intvalue,err:=strconv.Atoi(value.Value.Value)
		//check(err)
		if value.Value.Type== lexer.TRUE{

			return append(list,"TRUE")
		}else if value.Value.Type== lexer.FALSE{
			return append(list,"FALSE")
		}

		return append(list,"PUSH",value.Value.Value)
	}

	if value,ok:=ast.(parser.Binary);ok{
		list = append(list, Compile(value.Left)...)
		list = append(list, Compile(value.Right)...)
		switch value.Operator{
		case lexer.PLUS:
			opcode="ADD"
		case lexer.MINUS:
			opcode="SUB"
		case lexer.SLASH:
			opcode="DIV"
		case lexer.STAR:
			opcode="MUL"
		case lexer.AND:
			opcode="AND"
		case lexer.OR:
			opcode="OR"
	case lexer.TRUE:
			opcode="TRUE"
	case lexer.FALSE:
			opcode="FALSE"
		}
		list = append(list, opcode)
	}
	return list
}


