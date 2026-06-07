package vm

import (
	"fmt"
	"strconv"
	"lang/internal/backend/compiler"
)

func ToBytecode(program []string)([]byte,[]int,[]string){
	var constantTable []int
	var varConstTable []string
	var bytearray []byte
	for _,val:=range program{
		if opcode,ok:=compiler.Opcode[val];ok{
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

func Machine(bytearray []byte, counterTable []int, varConstTable []string,stack *[]int,stackpointer *int,heap *map[string]int) int {
    var programCounter int
	var ans int
    for programCounter < len(bytearray) {
        opcode := int(bytearray[programCounter])
        switch compiler.OpName[opcode] {
        case "PUSH":
            programCounter++
            number := counterTable[int(bytearray[programCounter])]
            (*stack)[*stackpointer] = number
            *stackpointer++
            programCounter++
		case "VAR_DEC":
			programCounter++
			*stackpointer--
            globalvar:= varConstTable[int(bytearray[programCounter])]
            (*heap)[globalvar] = (*stack)[*stackpointer]
            *stackpointer++
            programCounter++
		case "VAR":
			programCounter++
			ident:=varConstTable[int(bytearray[programCounter])]
			value:=(*heap)[ident]
			(*stack)[*stackpointer]=value
			*stackpointer++
            programCounter++
        case "ADD":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right :=  (*stack)[*stackpointer]
			ans = left + right
            (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
	  	case "MUL":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right := (*stack)[*stackpointer]
            ans = left * right
            (*stack)[*stackpointer] = ans
            *stackpointer++
            programCounter++
  		case "DIV":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right :=(*stack)[*stackpointer]
            ans = left / right
           (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
				
        }
    }
	fmt.Println(heap)
    return ans
}
