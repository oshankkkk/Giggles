package main
import "fmt"
const (
	PUSH = iota
	ADD
	MINUS
)

var program = []int{
	PUSH, 3,
	PUSH, 4,
	ADD,
	PUSH, 5,
	MINUS,
}

func vm(program []int){
	stack:=make([]int,10)
	var stackpointer int 
	var programCounter int
	
	for programCounter<len(program){
		currentinstruction:=program[programCounter]
		switch currentinstruction{
		case ADD:
			//remove the free one
			stackpointer--
			left:=stack[stackpointer]
			right:=stack[stackpointer-1]
			stackpointer--
			sum:=left+right
			stack[stackpointer]=sum
			//add another free one
			stackpointer++

		case PUSH:
			stack[stackpointer]=program[programCounter+1]
			stackpointer++
			programCounter++

		case MINUS:
			//remove the free one
			stackpointer--
			left:=stack[stackpointer]
			right:=stack[stackpointer-1]
			stackpointer--
			sum:=left-right
			stack[stackpointer]=sum
			//add another free one
			stackpointer++

		}
	programCounter++
	fmt.Println(stackpointer)
	}

}

func main(){
	fmt.Println("vm start")
vm(program)
}
