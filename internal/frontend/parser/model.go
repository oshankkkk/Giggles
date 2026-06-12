package parser

import "lang/internal/frontend/lexer"

type ASTNode interface {
	Expression()
}

type Program struct {
	Statements []ASTNode
}

type Literal struct {
	Value    lexer.Token
	Line     int
	Column   int
}

//type Literal struct {
//	Value    lexer.Token
//	Line     int
//	Column   int
//}


type Identifier struct {
	Name     lexer.Token
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
}

type ExprStatement struct {
	Expr   ASTNode
	Line   int
	Column int
}

type Condition struct{
	Condition ASTNode
	HasElse bool
	Looped bool
	Result ASTNode
	ElseResult ASTNode
	Line int
	Column int
	
}
func (n Program) Expression()       {}
func (n Literal) Expression()       {}
func (n Identifier) Expression()    {}
func (n Binary) Expression()        {}
func (n Unary) Expression()         {}
func (n Groups) Expression()        {}
func (n Condition) Expression()        {}
func (n VarDecl) Expression()       {}
func (n ExprStatement) Expression() {}

