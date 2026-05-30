package parser

import "lang/internal/lexer"

type ASTNode interface {
	Expression()
}

type Program struct {
	Statements []ASTNode
}

type Literal struct {
	NodeName string
	Value    lexer.Token
	Line     int
	Column   int
}

type Identifier struct {
	NodeName string
	Name     lexer.Token
	Line     int
	Column   int
}

type Binary struct {
	NodeName string
	Left     ASTNode
	Right    ASTNode
	Operator lexer.TokenType
	Line     int
	Column   int
}

type Unary struct {
	NodeName string
	Value    ASTNode
	Line     int
	Column   int
}

type Groups struct {
	NodeName string
	Value    ASTNode
	Line     int
	Column   int
}

type VarDecl struct {
	NodeName string
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

func (n Program) Expression()       {}
func (n Literal) Expression()       {}
func (n Identifier) Expression()    {}
func (n Binary) Expression()        {}
func (n Unary) Expression()         {}
func (n Groups) Expression()        {}
func (n VarDecl) Expression()       {}
func (n ExprStatement) Expression() {}

