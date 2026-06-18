package vm

import (
	"fmt"
	"lang/internal/compiler"
)

type GVM struct {
	programCounter int
	stack          []int
	stackpointer   int
	scopepointer int
	globalscope []int
	//variable []int
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

	fmt.Println(len(g.stack),"glen mmm")
	//g.globalscope=make(map[string]int)	
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
		case compiler.SETGLOBAL:
			// the last push val is put in the global table
			g.stackpointer--
			g.globalscope = append(g.globalscope, g.stack[g.scopepointer])
			fmt.Println(len(g.globalscope),"naman")
			g.programCounter++
		case compiler.GETGLOBAL:
			// the values in the counter table should be the index of the global var
			//pcounter is its index
			g.programCounter++
			fmt.Println(g.programCounter,"cucuc")
			g.stack=append(g.stack,g.globalscope[int(bytearray[g.programCounter])])
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
	fmt.Println("alelalxxxxala")
			if !toBool(g.stack[g.stackpointer]) {
			
	fmt.Println("aee33yylelalala")

	fmt.Println(len(counterTable))
	fmt.Println(g.programCounter)
	
	fmt.Println(len(bytearray),"dhdhdhbbb")
				address := counterTable[int(bytearray[g.programCounter])]
				g.programCounter = address

		fmt.Println("alelalala")
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
	fmt.Println(len(g.stack),"glen")
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

		//case compiler.SETLOCAL:
		//case compiler.GETLOCAL:
//			case compiler.ASS:
//			*stackpointer--
//			newval:=(*stack)[*stackpointer]
//			*stackpointer--
//			index:=(*stack)[*stackpointer]
//			ident:=varConstTable[index]
//			(*heap)[ident]=newval
//			*stackpointer++
//			programCounter++
//

