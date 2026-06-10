package vm

import (
	"fmt"
//	"strconv"
	"lang/internal/backend/compiler"
)


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
//            *stackpointer++
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
		case "SUB":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right :=  (*stack)[*stackpointer]
			ans = left - right
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
            ans = right / left
           (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
				
        }
    }


	fmt.Println(heap)
	fmt.Println(stack)
    return ans
}
