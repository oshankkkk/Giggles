package vm

import (
	"fmt"
	"lang/internal/backend/compiler"
)

func Machine(bytearray []byte, counterTable []int, varConstTable []string,stack *[]int,stackpointer *int,heap *map[string]int) int {

    var programCounter int
	var ans int
    for programCounter < len(bytearray) {

		opcode := int(bytearray[programCounter])

		fmt.Println(
			"ProCount:", programCounter,
			"OPname:", compiler.OpName[opcode],
			"SPointer:", *stackpointer,
			"STACK:", (*stack)[:*stackpointer],
		)

		switch compiler.OpName[opcode] {

		case "PUSH":
            programCounter++
            number := counterTable[int(bytearray[programCounter])]
			fmt.Println(number)
            (*stack)[*stackpointer] = number
            *stackpointer++
            programCounter++
		 case "TRUE":
            (*stack)[*stackpointer] = toInt(true)
            *stackpointer++
            programCounter++
		 case "FALSE":
            (*stack)[*stackpointer] = toInt(false)
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
//			*stackpointer++
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
			ans = right - left
            (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
	  	case "MUL":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right := (*stack)[*stackpointer]
            ans = right*left
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
		case "AND":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right :=(*stack)[*stackpointer]
            ans = toInt(toBool(right) && toBool(left))
           (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
		case "OR":
            *stackpointer--
            left := (*stack)[*stackpointer]
            *stackpointer--
            right :=(*stack)[*stackpointer]
            ans = toInt(toBool(right) || toBool(left))
           (*stack)[*stackpointer]= ans
            *stackpointer++
            programCounter++
		case "GT":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right > left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
            programCounter++
		case "LT":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right < left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
            programCounter++
		case "GTE":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right >= left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
            programCounter++
		case "LTE":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right <= left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
            programCounter++
		case "EQ":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right == left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
            programCounter++
		case "NEQ":
			*stackpointer--
			left := (*stack)[*stackpointer]
			*stackpointer--
			right := (*stack)[*stackpointer]
			ans = toInt(right != left)
			(*stack)[*stackpointer] = ans
			*stackpointer++
			programCounter++
		case "JMP":
			programCounter++
			address:=counterTable[int(bytearray[programCounter])]
			programCounter=address
		case "JIF":
			programCounter++
			*stackpointer--
			if !toBool((*stack)[*stackpointer]){
			address:=counterTable[int(bytearray[programCounter])]
			programCounter=address
			}else{
				programCounter++
			}
			*stackpointer++
			
		}
    }

	fmt.Println(heap)
	fmt.Println(stack)
    return ans
}
func toBool(val int)bool{
	if val!=0{
		return true
	}
return false
}
func toInt(val bool)int{
	if val{
		return 1
	}
	return 0
}
