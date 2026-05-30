package vm

import (
	"fmt"
	"strconv"
)
func check(err error){
	if err!=nil{
		fmt.Println(err)
	}
}
var Opcode = map[string]int{
	"PUSH": 1,
	"ADD":  2,
	"SUB":  3,
}
var OpName = map[int]string{
    1: "PUSH",
    2: "ADD",
    3: "SUB",
}
func ToBytecode(program []string)([]byte,[]int){
	var constantTable []int

	var bytearray []byte
	for _,val:=range program{
		if opcode,ok:=Opcode[val];ok{
			bytearray = append(bytearray, byte(opcode))
		}else {
			val,err:=strconv.Atoi(val)
			check(err)
			constantTable= append(constantTable, val)
			bytearray = append(bytearray, byte(len(constantTable)-1))
		}
	}
	fmt.Println("this worked")
	return bytearray,constantTable
}
func Machine(bytearray []byte, counterTable []int) int {
    var stack = make([]int, 256)
    var stackpointer int
    var programCounter int

    for programCounter < len(bytearray) {
        opcode := int(bytearray[programCounter])

        switch OpName[opcode] {
        case "PUSH":
            programCounter++
            number := counterTable[int(bytearray[programCounter])]
            stack[stackpointer] = number
            stackpointer++
            programCounter++

        case "ADD":
            stackpointer--
            left := stack[stackpointer]
            stackpointer--
            right := stack[stackpointer]
            ans := left + right
            stack[stackpointer] = ans
            stackpointer++
            programCounter++
            return ans
        }
    }
    return 0
}
