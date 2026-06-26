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
func TestParseFunctionWithParams(t *testing.T) {
	l := &lexer.Lexer{}
	l.ReadLine("fn add(int x, bool y)\nend")
	p := &Parser{}

	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)
	assert.Len(t, prog.Statements, 1)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	fn, ok := exprStmt.Expr.(Function)
	assert.True(t, ok)
	assert.Equal(t, "add", fn.Name)
	assert.Len(t, fn.Params, 2)

	assert.Equal(t, "int", fn.Params[0].Typedeff)
	assert.Equal(t, "x", fn.Params[0].Name.Value)

	assert.Equal(t, "bool", fn.Params[1].Typedeff)
	assert.Equal(t, "y", fn.Params[1].Name.Value)
}

func TestParseFunctionNoParams(t *testing.T) {
	l := &lexer.Lexer{}
	l.ReadLine("fn foo\nend")
	p := &Parser{}

	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	fn, ok := exprStmt.Expr.(Function)
	assert.True(t, ok)
	assert.Equal(t, "foo", fn.Name)
	assert.Len(t, fn.Params, 0)
}
func TestParseFunctionParamsUsedInBody(t *testing.T) {
	// Params x and y must be visible inside the body, otherwise the parser panics.
	l := &lexer.Lexer{}
	l.ReadLine("fn add(int x, int y)\nint z = x + y\nend")
	p := &Parser{}

	// Should not panic
	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	fn, ok := exprStmt.Expr.(Function)
	assert.True(t, ok)
	assert.Equal(t, "add", fn.Name)
	assert.Len(t, fn.Params, 2)
	assert.Len(t, fn.Content, 1) // one statement: int z = x + y
}

func TestParseCallWithArgs(t *testing.T) {
	// Declare a function first so the symbol table knows about it,
	// then call it with two numeric arguments.
	l := &lexer.Lexer{}
	l.ReadLine("fn add(int x, int y)\nend\nadd(2, 3)")
	p := &Parser{}

	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)
	assert.Len(t, prog.Statements, 2) // fn decl + call expression

	callStmt, ok := prog.Statements[1].(ExprStatement)
	assert.True(t, ok)

	call, ok := callStmt.Expr.(Call)
	assert.True(t, ok)
	assert.Equal(t, "add", call.Function)
	assert.Len(t, call.Args, 2)

	arg0, ok := call.Args[0].Value.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "2", arg0.Value.Value)

	arg1, ok := call.Args[1].Value.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "3", arg1.Value.Value)
}

func TestParseFunctionWithReturnType(t *testing.T) {
	// fn foo(int a, int b)int — return type on same line as declaration
	l := &lexer.Lexer{}
	l.ReadLine("fn foo(int a, int b)int\nreturn 2\nend")
	p := &Parser{}

	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	fn, ok := exprStmt.Expr.(Function)
	assert.True(t, ok)
	assert.Equal(t, "foo", fn.Name)
	assert.Equal(t, "int", fn.ReturnType)
	assert.False(t, fn.IsVoid)
	assert.Len(t, fn.Params, 2)
	assert.Len(t, fn.Content, 1)

	ret, ok := fn.Content[0].(ReturnStmt)
	assert.True(t, ok)

	lit, ok := ret.Value.(Literal)
	assert.True(t, ok)
	assert.Equal(t, "2", lit.Value.Value)
}

func TestParseFunctionVoidReturn(t *testing.T) {
	// fn with no return type — IsVoid should be true, bare return produces ReturnStmt{Value:nil}
	l := &lexer.Lexer{}
	l.ReadLine("fn greet\nreturn\nend")
	p := &Parser{}

	ast := p.Run(l)
	prog, ok := ast.(Program)
	assert.True(t, ok)

	exprStmt, ok := prog.Statements[0].(ExprStatement)
	assert.True(t, ok)

	fn, ok := exprStmt.Expr.(Function)
	assert.True(t, ok)
	assert.Equal(t, "", fn.ReturnType)
	assert.True(t, fn.IsVoid)

	ret, ok := fn.Content[0].(ReturnStmt)
	assert.True(t, ok)
	assert.Nil(t, ret.Value)
}
