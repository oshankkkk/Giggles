package vm

import (
	"fmt"
	"lang/internal/compiler"
)

type GVM struct {
	programCounter int
	stack          []int
	stackpointer   int
}

func (g *GVM) debugPrint(opcode int) {
	fmt.Println(
		"ProCount:", g.programCounter,
		"OPname:", compiler.Opcode(opcode),
		"SPointer:", g.stackpointer,
		"STACK:", g.stack[:g.stackpointer],
	)
}

func (g *GVM) Machine(bytearray []byte, counterTable []int) int {
	g.stack = make([]int, 1024)
	var ans int
	for g.programCounter < len(bytearray) {
		opcode := int(bytearray[g.programCounter])
		g.debugPrint(opcode)
		switch compiler.Opcode(opcode) {
		case compiler.PUSH:
			g.programCounter++
			number := counterTable[int(bytearray[g.programCounter])]
			g.stack[g.stackpointer] = number
			g.stackpointer++
			g.programCounter++
		case compiler.NEQ:
			g.stackpointer--
			left := g.stack[g.stackpointer]
			g.stackpointer--
			right := g.stack[g.stackpointer]
			ans = toInt(right != left)
			g.stack[g.stackpointer] = ans
			g.stackpointer++
			g.programCounter++
		case compiler.JMP:
			g.programCounter++
			address := counterTable[int(bytearray[g.programCounter])]
			g.programCounter = address
		case compiler.JIF:
			g.programCounter++
			g.stackpointer--
			if !toBool(g.stack[g.stackpointer]) {
				address := counterTable[int(bytearray[g.programCounter])]
				g.programCounter = address
			} else {
				g.programCounter++
			}
		case compiler.ADD, compiler.SUB, compiler.MUL, compiler.DIV:
			ans=g.MathOps(opcode)
		case compiler.GT, compiler.LT, compiler.GTE, compiler.LTE, compiler.EQ:
			ans=g.Comparisons(opcode)
		case compiler.AND, compiler.OR, compiler.TRUE, compiler.FALSE:
			ans=g.BoolOps(opcode)
		}
	}
	fmt.Println(g.stack)
	return ans
}


func (g *GVM) MathOps(opcode int) int{
	var ans int
	switch compiler.Opcode(opcode) {
	case compiler.ADD:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = left + right
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.SUB:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right - left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.MUL:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right * left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.DIV:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = right / left
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	}
	return ans
}

func (g *GVM) Comparisons(opcode int) int{
	var ans int
	switch compiler.Opcode(opcode) {
	case compiler.GT:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right > left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.LT:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right < left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.GTE:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right >= left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.LTE:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right <= left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.EQ:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(right == left)
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	}
	return ans
}

func (g *GVM) BoolOps(opcode int) int{
	var ans int
	switch compiler.Opcode(opcode) {
	case compiler.AND:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(toBool(right) && toBool(left))
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.OR:
		g.stackpointer--
		left := g.stack[g.stackpointer]
		g.stackpointer--
		right := g.stack[g.stackpointer]
		ans = toInt(toBool(right) || toBool(left))
		g.stack[g.stackpointer] = ans
		g.stackpointer++
		g.programCounter++
	case compiler.TRUE:
		g.stack[g.stackpointer] = toInt(true)
		g.stackpointer++
		g.programCounter++
	case compiler.FALSE:
		g.stack[g.stackpointer] = toInt(false)
		g.stackpointer++
		g.programCounter++
	}
	return  ans
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


