package vm
//
//import (
//	"testing"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestMachineBinaryOp(t *testing.T) {
//	// PUSH 10, PUSH 20, ADD
//	// PUSH=1, ADD=2
//	bytearray := []byte{1, 0, 1, 1, 2}
//	counterTable := []int{10, 20}
//	varConstTable := []string{}
//	
//	stack := make([]int, 10)
//	stackpointer := 0
//	heap := make(map[string]int)
//
//	ans := Machine(bytearray, counterTable, varConstTable, &stack, &stackpointer, &heap)
//	
//	assert.Equal(t, 30, ans)
//	assert.Equal(t, 1, stackpointer)
//	assert.Equal(t, 30, stack[0])
//}
//
//func TestMachineVariables(t *testing.T) {
//	// int x = 5; x * 2
//	// PUSH 5, VAR_DEC x, VAR x, PUSH 2, MUL
//	// PUSH=1, VAR_DEC=6, VAR=7, MUL=4
//	bytearray := []byte{1, 0, 6, 0, 7, 0, 1, 1, 4}
//	counterTable := []int{5, 2}
//	varConstTable := []string{"x"}
//	
//	stack := make([]int, 10)
//	stackpointer := 0
//	heap := make(map[string]int)
//
//	ans := Machine(bytearray, counterTable, varConstTable, &stack, &stackpointer, &heap)
//	
//	assert.Equal(t, 10, ans)
//	assert.Equal(t, 1, stackpointer)
//	assert.Equal(t, 10, stack[0])
//	assert.Equal(t, 5, heap["x"])
//}
//
//func TestMachineBooleanAndConditionals(t *testing.T) {
//	// TRUE, FALSE, OR
//	// TRUE=10, FALSE=11, OR=9
//	bytearray := []byte{10, 11, 9}
//	counterTable := []int{}
//	varConstTable := []string{}
//	
//	stack := make([]int, 10)
//	stackpointer := 0
//	heap := make(map[string]int)
//
//	ans := Machine(bytearray, counterTable, varConstTable, &stack, &stackpointer, &heap)
//	
//	assert.Equal(t, 1, ans) // 1 means true
//	assert.Equal(t, 1, stackpointer)
//	assert.Equal(t, 1, stack[0])
//}
//
//func TestMachineComparisons(t *testing.T) {
//	// PUSH 10, PUSH 20, LT (10 < 20)
//	// PUSH=1, LT=13
//	bytearray := []byte{1, 0, 1, 1, 13}
//	counterTable := []int{10, 20}
//	varConstTable := []string{}
//	
//	stack := make([]int, 10)
//	stackpointer := 0
//	heap := make(map[string]int)
//
//	ans := Machine(bytearray, counterTable, varConstTable, &stack, &stackpointer, &heap)
//	
//	assert.Equal(t, 1, ans) // 10 < 20 is true
//}
