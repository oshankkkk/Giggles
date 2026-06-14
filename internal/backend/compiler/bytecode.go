package compiler
import (
//	"fmt"
	"strconv"

)
var Opcode = map[string]int{
	"PUSH":  1,
	"ADD":   2,
	"SUB":   3,
	"MUL":   4,
	"DIV":   5,
	"VAR_DEC": 6,
	"VAR":   7,
	"AND":   8,
	"OR":    9,
	"TRUE":  10,
	"FALSE": 11,
	"GT":    12, // >
	"LT":    13, // <
	"GTE":   14, // >=
	"LTE":   15, // <=
	"EQ":    16, // ==
	"NEQ":   17, // !=
	"JMP":18,
	"JIF":19,
	"ASS":20,
}

var OpName = map[int]string{
	1:  "PUSH",
	2:  "ADD",
	3:  "SUB",
	4:  "MUL",
	5:  "DIV",
	6:  "VAR_DEC",
	7:  "VAR",
	8:  "AND",
	9:  "OR",
	10: "TRUE",
	11: "FALSE",
	12: "GT",
	13: "LT",
	14: "GTE",
	15: "LTE",
	16: "EQ",
	17: "NEQ",
	18:"JMP",
	19:"JIF",
	20:"ASS",
}

func ToBytecode(program []string)([]byte,[]int,[]string){
	var constantTable []int
	var varConstTable []string
	var bytearray []byte
	for _,val:=range program{
		if opcode,ok:=Opcode[val];ok{

			bytearray = append(bytearray, byte(opcode))
		}else {
			digit,err:=strconv.Atoi(val)
			//not numbers
			if err!=nil{
			varConstTable = append(varConstTable, val)	
			bytearray = append(bytearray, byte(len(varConstTable)-1))
			continue
			}
			constantTable= append(constantTable, digit)
			bytearray = append(bytearray, byte(len(constantTable)-1))
		}
	}
	return bytearray,constantTable,varConstTable
}

func ToBytecode2(program []string) ([]byte, []int, []string) {
    var constantTable []int
    var varConstTable []string
    var bytearray []byte

    // First pass: build IR-index → byte-index map
    irToBytePos := make(map[int]int)
    bytePos := 0
    for i := 0; i < len(program); i++ {
        irToBytePos[i] = bytePos
        val := program[i]
        if _, ok := Opcode[val]; ok {
            bytePos++
            // if this opcode takes an operand, skip it in IR
            if val == "PUSH" || val == "JIF" || val == "JMP" ||
                val == "VAR_DEC" || val == "VAR" {
                i++
                bytePos++ // operand byte
            }
        } else {
            bytePos++ // standalone operand (shouldn't happen at top level)
        }
    }

    // Second pass: emit bytes
    for i := 0; i < len(program); i++ {
        val := program[i]
        if opcode, ok := Opcode[val]; ok {
            bytearray = append(bytearray, byte(opcode))

            // opcodes that take an operand
            switch val {
            case "JMP", "JIF":
                i++
                irTarget, _ := strconv.Atoi(program[i])
                byteTarget := irToBytePos[irTarget] // IR index → byte index
                constantTable = append(constantTable, byteTarget)
                bytearray = append(bytearray, byte(len(constantTable)-1))

            case "PUSH":
                i++
                digit, _ := strconv.Atoi(program[i])
                constantTable = append(constantTable, digit)
                bytearray = append(bytearray, byte(len(constantTable)-1))

            case "VAR_DEC", "VAR":
                i++
                varConstTable = append(varConstTable, program[i])
                bytearray = append(bytearray, byte(len(varConstTable)-1))
            }
        }
    }

    return bytearray, constantTable, varConstTable
}
