package compiler

import (
	"fmt"
	"lang/internal/lexer"
	"lang/internal/parser"
	"strconv"
)
type State struct {
CounterTable  []int
}

func (c *State) checkTable(){}
func (c *State) ToBytes(ast parser.ASTNode,scope *Scope) []byte{
	var bytecode []byte
	if value, ok := ast.(parser.Program); ok {
		for _, val := range value.Statements {
			bytecode = append(bytecode, c.ToBytes(val,scope)...)
		}
	}
	if value, ok := ast.(parser.ExprStatement); ok {
		bytecode = append(bytecode, c.ToBytes(value.Expr,scope)...)
	}
	if value, ok := ast.(parser.Groups); ok {
		bytecode = append(bytecode, c.ToBytes(value.Value,scope)...)
	}
	
	if value, ok := ast.(parser.VarDecl); ok {
//check for scope and put the flat bytes here with pointer address and var table index
//value	
	bytecode = append(bytecode, c.ToBytes(value.Value,scope)...)
// the variable name is after vardec
		scope.AddVariable(value.Name.Value)
		bytecode = append(bytecode, byte(SETGLOBAL))
		//what this opcode does
	}
	if value, ok := ast.(parser.Identifier); ok {
//get local/global
		if info,ok:=scope.VarLookup(value.Name.Value);ok{
			//this means the var is in the vm now
		bytecode = append(bytecode, byte(GETGLOBAL),byte(info.id))
	}else{
		panic("var undefined")
	}
}
	if value, ok := ast.(parser.Condition); ok {
		condPos := len(bytecode)
		bytecode = append(bytecode, c.ToBytes(value.Condition,scope)...)
		jifPos := len(bytecode)
		fmt.Println(len(c.CounterTable),"jejeje")
		//placeholder 0 
	 	bytecode = append(bytecode, byte(JIF), 0)
		
		resultCode := []byte{}
		//local :=EnterScope(scope)
		for _, r := range value.Result {
			resultCode = append(resultCode, c.ToBytes(r,scope)...)
		}
		bytecode = append(bytecode, resultCode...)
		if value.Looped {
			bytecode = append(bytecode, byte(JMP), byte(condPos))
			elseCode := []byte{}

			//local :=EnterScope(scope)
			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.ToBytes(e,scope)...)
			}
			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)

			bytecode = append(bytecode, elseCode...)
		} else if value.HasElse {
			elseCode := []byte{}

			//local :=EnterScope(scope)
			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.ToBytes(e,scope)...)
			}
			jmpPos := len(bytecode)

			bytecode = append(bytecode, byte(JMP), 0)

//	bytecode[jifPos+1] = byte(len(bytecode))

			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)


			bytecode = append(bytecode, elseCode...)

		//	bytecode[jmpPos+1] = byte(len(bytecode))

			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jmpPos+1] = byte(len(c.CounterTable)-1)


		} else {
		c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)


//			bytecode[jifPos+1] = byte(len(bytecode))
		}
	}


	if value, ok := ast.(parser.Literal); ok {
		if value.Value.Type == lexer.TRUE {
			return append(bytecode, byte(TRUE))
		} else if value.Value.Type == lexer.FALSE {
			return append(bytecode, byte(FALSE))
		}
		index,err:=strconv.Atoi(value.Value.Value)
		if err!=nil{
			panic(err)
		}
		c.CounterTable = append(c.CounterTable, index)
		return append(bytecode, byte(PUSH), byte(len(c.CounterTable)-1))
	}
	if value, ok := ast.(parser.Binary); ok {
		bytecode = append(bytecode, c.ToBytes(value.Left,scope)...)
		bytecode = append(bytecode, c.ToBytes(value.Right,scope)...)
		var opcode Opcode
		switch value.Operator {
		case lexer.MINUS:
			opcode = SUB
		case lexer.PLUS:
			opcode = ADD
		case lexer.SLASH:
			opcode = DIV
		case lexer.STAR:
			opcode = MUL
		case lexer.AND:
			opcode = AND
		case lexer.OR:
			opcode = OR
		case lexer.GREATER:
			opcode = GT
		case lexer.LESS:
			opcode = LT
		case lexer.GREATER_EQUAL:
			opcode = GTE
		case lexer.LESS_EQUAL:
			opcode = LTE
		case lexer.EQUAL_EQUAL:
			opcode = EQ
		case lexer.NOT_EQUAL:
			opcode = NEQ
		case lexer.EQUAL:
			opcode = ASS
		}
		bytecode = append(bytecode, byte(opcode))
	}


	return bytecode
}

type Opcode int

const (
	PUSH Opcode = iota + 1
	ADD
	SUB
	MUL
	DIV
	VAR_DEC
	VAR
	AND
	OR
	TRUE
	FALSE
	GT
	LT
	GTE
	LTE
	EQ
	NEQ
	JMP
	JIF
	ASS
	GETGLOBAL
	SETGLOBAL
	SETLOCAL
	GETLOCAL
)


