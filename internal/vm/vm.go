package vm

import (
	"fmt"
	"strconv"
	"lang/internal/compiler"
)

func ToBytecode(program []string)([]byte,[]int,[]string){
	var constantTable []int
	var varConstTable []string
	var bytearray []byte
	for _,val:=range program{

			fmt.Println(val,"og val")	
		if opcode,ok:=compiler.Opcode[val];ok{
			bytearray = append(bytearray, byte(opcode))
			fmt.Println(val,"opopk")	
		}else {
			fmt.Println(val,"valval")	
			digit,err:=strconv.Atoi(val)
			if err!=nil{
			fmt.Println(val,"string")	
			varConstTable = append(varConstTable, val)	
			bytearray = append(bytearray, byte(len(varConstTable)-1))
			continue
			}
			constantTable= append(constantTable, digit)
			bytearray = append(bytearray, byte(len(constantTable)-1))
		}
	}
	fmt.Println("this worked")
	return bytearray,constantTable,varConstTable
}

func Machine(bytearray []byte, counterTable []int, varConstTable []string) int {
    var stack = make([]int, 1024)
    var stackpointer int
    var programCounter int
	var heap=make(map[string]int)
	var ans int
    for programCounter < len(bytearray) {
        opcode := int(bytearray[programCounter])

        switch compiler.OpName[opcode] {
        case "PUSH":
            programCounter++
            number := counterTable[int(bytearray[programCounter])]
            stack[stackpointer] = number
            stackpointer++
            programCounter++
		case "VAR_DEC":
			programCounter++
			fmt.Println("vardec have")
			stackpointer--
            globalvar:= varConstTable[int(bytearray[programCounter])]
			fmt.Println(globalvar)
            heap[globalvar] = stack[stackpointer]
            stackpointer++
            programCounter++

        case "ADD":
            stackpointer--
            left := stack[stackpointer]
            stackpointer--
            right := stack[stackpointer]
            ans = left + right
            stack[stackpointer] = ans
            stackpointer++
            programCounter++
	  	case "MUL":
            stackpointer--
            left := stack[stackpointer]
            stackpointer--
            right := stack[stackpointer]
            ans = left * right
            stack[stackpointer] = ans
            stackpointer++
            programCounter++
  		case "DIV":
            stackpointer--
            left := stack[stackpointer]
            stackpointer--
            right := stack[stackpointer]
            ans = left / right
            stack[stackpointer] = ans
            stackpointer++
            programCounter++
				
        }
    }
	fmt.Println(heap)
    return ans
}
