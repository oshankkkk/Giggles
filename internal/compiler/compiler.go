package compiler

import (
	"lang/internal/lexer"
	"lang/internal/parser"
	"strconv"
)

type Compiler struct {
CounterTable  []int
}

func (c *Compiler) Compile(ast parser.ASTNode) []byte{
	var bytecode []byte
	if value, ok := ast.(parser.Program); ok {
		for _, val := range value.Statements {
			bytecode = append(bytecode, c.Compile(val)...)
		}
	}
	if value, ok := ast.(parser.ExprStatement); ok {
		bytecode = append(bytecode, c.Compile(value.Expr)...)
	}
	if value, ok := ast.(parser.Groups); ok {
		bytecode = append(bytecode, c.Compile(value.Value)...)
	}
//	if value, ok := ast.(parser.VarDecl); ok {
//		bytecode = append(bytecode, c.Compile(value.Value)...)
//		bytecode = append(bytecode, byte(VAR_DEC), value.Name.Value)
//	}
//	if value, ok := ast.(parser.Identifier); ok {
//		bytecode = append(bytecode, byte(VAR), value.Name.Value)
//	}
	if value, ok := ast.(parser.Condition); ok {
		condPos := len(bytecode)
		bytecode = append(bytecode, c.Compile(value.Condition)...)
		jifPos := len(bytecode)
		bytecode = append(bytecode, byte(JIF), 0)
		resultCode := []byte{}
		for _, r := range value.Result {
			resultCode = append(resultCode, c.Compile(r)...)
		}
		bytecode = append(bytecode, resultCode...)
		if value.Looped {
			bytecode = append(bytecode, byte(JMP), byte(condPos))
			elseCode := []byte{}
			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.Compile(e)...)
			}
			bytecode[jifPos+1] = byte(len(bytecode))
			bytecode = append(bytecode, elseCode...)
		} else if value.HasElse {
			elseCode := []byte{}
			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.Compile(e)...)
			}
			jmpPos := len(bytecode)
			bytecode = append(bytecode, byte(JMP), 0)
			bytecode[jifPos+1] = byte(len(bytecode))
			bytecode = append(bytecode, elseCode...)
			bytecode[jmpPos+1] = byte(len(bytecode))
		} else {
			bytecode[jifPos+1] = byte(len(bytecode))
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
		bytecode = append(bytecode, c.Compile(value.Left)...)
		bytecode = append(bytecode, c.Compile(value.Right)...)
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
)


