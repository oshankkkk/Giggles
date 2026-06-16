package parser

import (
	"testing"
	"lang/internal/lexer"
	"github.com/stretchr/testify/assert"
)

func TestParseLiteral(t *testing.T) {
	l := &lexer.Lexer{}
	l.ReadLine("42")
	p := &Parser{}
	
	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)
	assert.Len(t, prog.Statements, 1)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	lit, ok := exprStmt.Expr.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "42", lit.Value.Value)
}

func TestParseVarDecl(t *testing.T) {
	l := &lexer.Lexer{}
	l.ReadLine("int x = 10")
	p := &Parser{}
	
	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)
	assert.Len(t, prog.Statements, 1)

	varDecl, ok := prog.Statements[0].(VarDecl)
	assert.True(t, ok)
	assert.Equal(t, "int", varDecl.Typedeff)
	assert.Equal(t, "x", varDecl.Name.Value)
	
	lit, ok := varDecl.Value.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "10", lit.Value.Value)
}

func TestParseBinaryExpression(t *testing.T) {
	l := &lexer.Lexer{}
	l.ReadLine("1 + 2 * 3")
	p := &Parser{}
	
	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)
	assert.Len(t, prog.Statements, 1)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	bin1, ok := exprStmt.Expr.(Binary)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS, bin1.Operator)

	lit1, ok := bin1.Left.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "1", lit1.Value.Value)

	bin2, ok := bin1.Right.(Binary)
	assert.True(t, ok)
	assert.Equal(t, lexer.STAR, bin2.Operator)

	lit2, ok := bin2.Left.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "2", lit2.Value.Value)

	lit3, ok := bin2.Right.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "3", lit3.Value.Value)
}
