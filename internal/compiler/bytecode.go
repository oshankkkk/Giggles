package compiler
import (
	"fmt"
	"strconv"

)
var Opcode = map[string]int{
	"PUSH": 1,
	"ADD":  2,
	"SUB":  3,
	"MUL":4,
	"DIV":5,
	"VAR_DEC":6,
}
var OpName = map[int]string{
    1: "PUSH",
    2: "ADD",
    3: "SUB",
	4: "MUL",
	5: "DIV",
	6:"VAR_DEC",
}
func ToBytecode(program []string)([]byte,[]int,[]string){
	var constantTable []int
	var varConstTable []string
	var bytearray []byte
	for _,val:=range program{

			fmt.Println(val,"og val")	
		if opcode,ok:=Opcode[val];ok{
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
