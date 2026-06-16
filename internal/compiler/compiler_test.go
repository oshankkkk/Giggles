package compiler
//
//import (
//	"testing"
//	"lang/internal/lexer"
//	"lang/internalfrontend/parser"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestCompileLiteral(t *testing.T) {
//	ast := parser.Literal{
//		Value: lexer.Token{Type: lexer.NUMBER, Value: "42"},
//	}
//	compiled := Compile(ast)
//	assert.Equal(t, []string{"PUSH", "42"}, compiled)
//}
//
//func TestCompileBinary(t *testing.T) {
//	ast := parser.Binary{
//		Left: parser.Literal{
//			Value: lexer.Token{Type: lexer.NUMBER, Value: "10"},
//		},
//		Operator: lexer.PLUS,
//		Right: parser.Literal{
//			Value: lexer.Token{Type: lexer.NUMBER, Value: "20"},
//		},
//	}
//	compiled := Compile(ast)
//	assert.Equal(t, []string{"PUSH", "10", "PUSH", "20", "ADD"}, compiled)
//}
//
//func TestCompileVarDecl(t *testing.T) {
//	ast := parser.VarDecl{
//		Name: lexer.Token{Type: lexer.IDENTIFIER, Value: "myvar"},
//		Value: parser.Literal{
//			Value: lexer.Token{Type: lexer.NUMBER, Value: "99"},
//		},
//	}
//	compiled := Compile(ast)
//	assert.Equal(t, []string{"PUSH", "99", "VAR_DEC", "myvar"}, compiled)
//}
//
//func TestToBytecode(t *testing.T) {
//	program := []string{"PUSH", "10", "PUSH", "20", "ADD", "VAR_DEC", "x", "VAR", "x"}
//	
//	bytearray, constants, vars := ToBytecode(program)
//	
//	// PUSH=1, ADD=2, VAR_DEC=6, VAR=7
//	// bytearray should be:
//	// 1 (PUSH), 0 (index of 10)
//	// 1 (PUSH), 1 (index of 20)
//	// 2 (ADD)
//	// 6 (VAR_DEC), 0 (index of "x")
//	// 7 (VAR), 0 (index of "x")
//	expectedBytes := []byte{1, 0, 1, 1, 2, 6, 0, 7, 1}
//	
//	assert.Equal(t, expectedBytes, bytearray)
//	assert.Equal(t, []int{10, 20}, constants)
//	assert.Equal(t, []string{"x", "x"}, vars)
//}

