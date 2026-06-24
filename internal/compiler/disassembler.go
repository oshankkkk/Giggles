package compiler

import (
	"fmt"
)

// OpcodeNames maps the integer opcode to its human-readable string.
var OpcodeNames = map[Opcode]string{
	PUSH:      "PUSH",
	ADD:       "ADD",
	SUB:       "SUB",
	MUL:       "MUL",
	DIV:       "DIV",
	VAR_DEC:   "VAR_DEC",
	VAR:       "VAR",
	AND:       "AND",
	OR:        "OR",
	TRUE:      "TRUE",
	FALSE:     "FALSE",
	GT:        "GT",
	LT:        "LT",
	GTE:       "GTE",
	LTE:       "LTE",
	EQ:        "EQ",
	NEQ:       "NEQ",
	JMP:       "JMP",
	JIF:       "JIF",
	ASS:       "ASS",
	GETGLOBAL: "GETGLOBAL",
	SETGLOBAL: "SETGLOBAL",
	SETLOCAL:  "SETLOCAL",
	GETLOCAL:  "GETLOCAL",
	NWFRM:     "NWFRM",
	RMFRM:     "RMFRM",
	STOP:      "STOP",
}

func (o Opcode) String() string {
	if name, ok := OpcodeNames[o]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN(%d)", int(o))
}

// Disassemble prints out the compiled bytecode in a human-readable format.
func Disassemble(bytearray []byte, counterTable []int) {
	fmt.Println(len(bytearray),"bytearray")
	fmt.Println("=== BYTECODE DISASSEMBLY ===")
	
	for pc := 0; pc < len(bytearray); {
		opcode := Opcode(bytearray[pc])
		fmt.Printf("%04d: %-10s", pc, opcode.String())

		// Determine if this instruction takes an argument
		switch opcode {
		case PUSH, SETGLOBAL, GETGLOBAL, ASS, JMP, JIF:
			// These opcodes have a 1-byte operand representing an index into the CounterTable
			if pc+1 < len(bytearray) {
				idx := int(bytearray[pc+1])
				val := 0
				if idx < len(counterTable) {
					val = counterTable[idx]
				}
				fmt.Printf(" [byteIdx: %d -> countertablevalue: %d]\n", idx, val)
			} else {
				fmt.Printf(" [MISSING OPERAND]\n")
			}
			pc += 2 // Move past opcode + operand
		default:
			// 1-byte opcode
			fmt.Println()
			pc += 1
		}
	}
	fmt.Println("============================")
}
