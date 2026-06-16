package vm

import (
	"fmt"
	"lang/internal/backend/compiler"
)

type GVM struct {
	programCounter int
	stack          []int
	stackpointer   int
}

func (g *GVM) debugPrint(opcode int) {
	fmt.Println(
		"ProCount:", g.programCounter,
		"OPname:", compiler.OpName[opcode],
		"SPointer:", g.stackpointer,
		"STACK:", g.stack[:g.stackpointer],
	)
}

func (g *GVM) MathOps(opcode int) {
	var ans int
	switch compiler.OpName[opcode] {
	case "ADD":
		fmt.Println("eheeeeeee")
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = left + right
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "SUB":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right - left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "MUL":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right * left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "DIV":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right / left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	}
}

func (g *GVM) Comparisons(opcode int) {
	var ans int
	switch compiler.OpName[opcode] {
	case "GT":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right > left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "LT":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right < left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "GTE":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right >= left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "LTE":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right <= left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "EQ":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right == left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	}
}

func (g *GVM) BoolOps(opcode int) {
	var ans int
	switch compiler.OpName[opcode] {
	case "AND":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(toBool(right) && toBool(left))
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "OR":
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(toBool(right) || toBool(left))
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case "TRUE":
		g.stack[g.stackpointer] = toInt(true)
		g.stackpointer++
		g.programCounter++
	case "FALSE":
		g.stack[g.stackpointer] = toInt(false)
		g.stackpointer++
		g.programCounter++
	}
}

func (g *GVM) Machine(bytearray []byte, counterTable []int) int {
	var ans int
	for g.programCounter < len(bytearray) {
		opcode := int(bytearray[g.programCounter])
		switch compiler.OpName[opcode] {
		case "PUSH":
			g.programCounter++
			number := counterTable[int(bytearray[g.programCounter])]
			fmt.Println(number)
			g.stack[g.stackpointer] = number
			g.stackpointer++
			g.programCounter++
		case "NEQ":
			g.stackpointer--
			left := g.stack[g.stackpointer]
			g.stackpointer--
			right := g.stack[g.stackpointer]
			ans = toInt(right != left)
			g.stack[g.stackpointer] = ans
			g.stackpointer++
			g.programCounter++
		case "JMP":
			g.programCounter++
			address := counterTable[int(bytearray[g.programCounter])]
			g.programCounter = address
		case "JIF":
			g.programCounter++
			g.stackpointer--
			if !toBool(g.stack[g.stackpointer]) {
				address := counterTable[int(bytearray[g.programCounter])]
				g.programCounter = address
			} else {
				g.programCounter++
			}
		case "ADD", "SUB", "MUL", "DIV":
			g.MathOps(opcode)
		case "GT", "LT", "GTE", "LTE", "EQ":
			g.Comparisons(opcode)
		case "AND", "OR", "TRUE", "FALSE":
			g.BoolOps(opcode)
		}
	}
	fmt.Println(g.stack)
	return ans
}

func toBool(val int) bool {
	return val != 0
}

func toInt(val bool) int {
	if val {
		return 1
	}
	return 0
}
//	case "VAR_DEC":
	//		programCounter++
	//		*stackpointer--
    //        globalvar:= varConstTable[int(bytearray[programCounter])]
    //        (*heap)[globalvar] = (*stack)[*stackpointer]
    //        programCounter++
	//	case "VAR":
	//		programCounter++
	//		//var string index
	//		ident:=varConstTable[int(bytearray[programCounter])]
	//		//give the variable string(heap key)
	//		value:=(*heap)[ident]
	//		(*stack)[*stackpointer]=value
	//		*stackpointer++
	//		(*stack)[*stackpointer]=int(bytearray[programCounter])
	//		*stackpointer++
    //        programCounter++
	//   case "ASS":
	//	   	*stackpointer--
	//		newval:=(*stack)[*stackpointer]
	//	   	*stackpointer--
	//		index:=(*stack)[*stackpointer]
	//		ident:=varConstTable[index]
	//		(*heap)[ident]=newval
	//		*stackpointer++
	//		programCounter++


