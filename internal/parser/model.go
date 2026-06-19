package parser

import "lang/internal/lexer"

type ASTNode interface {
	Expression()
}

type Program struct {
	Statements []ASTNode
}

type Call struct {
	Function string
	Args     []ASTNode
	Line     int
	Column   int
}

type Literal struct {
	Value    lexer.Token
	Line     int
	Column   int
}

type Identifier struct {
	Name     lexer.Token
	IsLocal  int
	Line     int
	Column   int
}

type Binary struct {
	Left     ASTNode
	Right    ASTNode
	Operator lexer.TokenType
	Line     int
	Column   int
}

type Unary struct {
	Value    ASTNode
	Line     int
	Column   int
}

type Groups struct {
	Value    ASTNode
	Line     int
	Column   int
}

type VarDecl struct {
	Typedeff string
	Name     lexer.Token
	Value    ASTNode
	Line     int
	Column   int
	IsConst bool
	IsLocal bool
}

type ExprStatement struct {
	Expr   ASTNode
	Line   int
	Column int
}
type Function struct{
	Name string
	Content []ASTNode
	Line   int
	Column int
	isVoid bool
}
type Condition struct{
	Condition ASTNode
	HasElse bool
	Looped bool
	Result []ASTNode
	ElseResult []ASTNode
	Line int
	Column int
	
}
func (n Program) Expression()       {}
func (n Literal) Expression()       {}
func (n Identifier) Expression()    {}
func (n Binary) Expression()        {}
func (n Function) Expression()        {}
func (n Unary) Expression()         {}
func (n Groups) Expression()        {}
func (n Condition) Expression()     {}
func (n VarDecl) Expression()       {}
func (n ExprStatement) Expression() {}
func (n Call) Expression() {}
