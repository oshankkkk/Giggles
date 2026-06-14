package compiler

import (
	//"fmt"
	"lang/internal/frontend/lexer"
	"lang/internal/frontend/parser"
	"strconv"
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

	if value,ok:=ast.(parser.Condition);ok{
		condPos:=len(list)
		list = append(list, Compile(value.Condition)...)

		jifPos := len(list)
		list = append(list, "JIF", "0")  

		resultCode := Compile(value.Result)

		list = append(list, resultCode...)  // then 

		if value.Looped{
			list = append(list, "JMP",strconv.Itoa(condPos))  // then 

			elseCode := Compile(value.ElseResult)
			list[jifPos+1] = strconv.Itoa(len(list))
			list = append(list, elseCode...)   // else

		}else if value.HasElse {
			elseCode := Compile(value.ElseResult)

			
			jmpPos := len(list)
			list = append(list, "JMP", "0")

			// JIF start of else
			list[jifPos+1] = strconv.Itoa(len(list))

			list = append(list, elseCode...)   // else


		
			// JMP past else if then works
			list[jmpPos+1] = strconv.Itoa(len(list))
		} else {
			// JIF jumps here past then
			list[jifPos+1] = strconv.Itoa(len(list))
		}

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
		switch value.Operator {
		case lexer.PLUS:
			opcode = "ADD"
		case lexer.MINUS:
			opcode = "SUB"
		case lexer.SLASH:
			opcode = "DIV"
		case lexer.STAR:
			opcode = "MUL"
		case lexer.AND:
			opcode = "AND"
		case lexer.OR:
			opcode = "OR"
		case lexer.GREATER:
			opcode = "GT"
		case lexer.LESS:
			opcode = "LT"
		case lexer.GREATER_EQUAL:
			opcode = "GTE"
		case lexer.LESS_EQUAL:
			opcode = "LTE"
		case lexer.EQUAL_EQUAL:
			opcode = "EQ"
		case lexer.NOT_EQUAL:
			opcode = "NEQ"
//		case lexer.EQUAL:
//			opcode="ASS"
		}		

		list = append(list, opcode)
	}
	return list
}


