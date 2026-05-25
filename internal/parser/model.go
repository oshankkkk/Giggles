package parser

import "lang/internal/lexer"

type ASTNode interface {
	expression()
}

type Program struct {
	statements []ASTNode
}

type Literal struct {
	nodeName string
	value    lexer.Token
	line     int
	column   int
}

type Identifier struct {
	nodeName string
	name     lexer.Token
	line     int
	column   int
}

type Binary struct {
	nodeName string
	left     ASTNode
	right    ASTNode
	operator lexer.TokenType
	line     int
	column   int
}

type Unary struct {
	nodeName string
	value    ASTNode
	line     int
	column   int
}

type Groups struct {
	nodeName string
	value    ASTNode
	line     int
	column   int
}

type VarDecl struct {
	nodeName string
	typedeff string
	name     lexer.Token
	value    ASTNode
	line     int
	column   int
}

type ExprStatement struct {
	expr   ASTNode
	line   int
	column int
}

func (n Program) expression()       {}
func (n Literal) expression()       {}
func (n Identifier) expression()    {}
func (n Binary) expression()        {}
func (n Unary) expression()         {}
func (n Groups) expression()        {}
func (n VarDecl) expression()       {}
func (n ExprStatement) expression() {}
