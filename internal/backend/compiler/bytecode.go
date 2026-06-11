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
