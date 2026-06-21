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
	basepointer int 
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
func (g *GVM) debugStack(vmop string) {
	fmt.Println(vmop)
    fmt.Printf(" SP=%-4d  BP=%-4d  stack[BP:SP]=%v\n",
         g.stackpointer, g.basepointer, g.stack[g.basepointer:g.stackpointer])
}
func (g *GVM) Machine(bytearray []byte, counterTable []int) int {
	g.stack = make([]int, 1024)
	//g.globalscope=make(map[string]int)	

	//g.debugPrint(opcode)
	var ans int
	for g.programCounter < len(bytearray) {
		
		opcode := int(bytearray[g.programCounter])
		//	fmt.Println(len(bytearray),"barray",g.programCounter,"pcpunter")
		//		fmt.Println(g.basepointer,"bp",g.stackpointer,"sp")
		g.debugPrint(opcode)
		switch compiler.Opcode(opcode) {
		case compiler.PUSH:
			g.programCounter++
			number := counterTable[int(bytearray[g.programCounter])]
			g.stack[g.stackpointer] = number
			g.stackpointer++
			g.programCounter++
		case compiler.NWFRM:
			// Layout: [saved_bp][return_addr] then locals above BP
			g.stack[g.stackpointer] = g.basepointer
			g.stackpointer++
			g.basepointer = g.stackpointer
			g.programCounter++

		case compiler.RMFRM:
			// Grab return address pushed just after NWFRM
			retAddr := g.stack[g.basepointer] 
			g.stackpointer = g.basepointer - 1  // unwind locals + saved bp slot
			g.basepointer = g.stack[g.basepointer-1] // restore caller's BP
			g.programCounter = retAddr           
//		case compiler.NWFRM:
//			fmt.Println("newframe happends",g.stackpointer)
//			g.stack[g.stackpointer]=g.basepointer
//			g.stackpointer++
//			g.basepointer=g.stackpointer
//			g.programCounter++
//			g.debugStack("new stack op")
//		case compiler.RMFRM:
//			fmt.Println("rmstack happend befor rm",g.stackpointer,g.basepointer)
//			g.stackpointer=g.basepointer
//
//			fmt.Println("after same as basepointer",g.stackpointer,g.basepointer)
//			g.stackpointer--
//
//			fmt.Println("after -1 deduct from stackpointer",g.stackpointer,g.basepointer)
//			//fmt.Println(g.stackpointer)
//			g.basepointer=g.stack[g.stackpointer]
//			fmt.Println("after stack base value interchange",g.stackpointer,g.basepointer)
//			g.programCounter++
			//g.stackpointer=g.stack[g.basepointer]
			//g.basepointer=g.stackpointer
			//g.debugStack("remove stack op")
		case compiler.SETGLOBAL:
			g.programCounter++  // move past opcode to read the index
			idx := counterTable[int(bytearray[g.programCounter])]
			g.stackpointer--
			// grow slice if needed (first time declaring)
			for len(g.globalscope) <= idx {
				g.globalscope = append(g.globalscope, 0)
			}
			g.globalscope[idx] = g.stack[g.stackpointer]
			g.programCounter++
		case compiler.GETGLOBAL:
			// the values in the counter table should be the index of the global var
			//pcounter is its index
			g.programCounter++
			g.stack[g.stackpointer]=g.globalscope[counterTable[int(bytearray[g.programCounter])]]
			g.stackpointer++
			g.programCounter++
		case compiler.ASS:
			g.programCounter++
			idx := counterTable[int(bytearray[g.programCounter])]
			g.stackpointer--                        
			right := g.stack[g.stackpointer]        
			g.globalscope[idx] = right
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
			by:=bytearray[g.programCounter]
			address := counterTable[int(by)]
			//fmt.Println()
			g.programCounter = address
			//fmt.Println(int(bytearray[ g.programCounter]),"jmp statement")
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
		case compiler.STOP:
			return ans
		default:
			fmt.Printf("no assignemnt for %s opecode you idoit\n",compiler.Opcode(opcode))
			//		case compiler.RETURN:
			//			g.stackpointer=g.framepointer
			//
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
		fmt.Println("add happends ")
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

		//case compiler.SETLOCAL:
		//case compiler.GETLOCAL:


