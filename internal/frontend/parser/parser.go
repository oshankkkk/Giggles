package parser

import (
	"fmt"
	"lang/internal/frontend/lexer"
)

type Parser struct {
pointer int
current lexer.Token
lexer *lexer.Lexer
}
func (p *Parser) Run(lexer *lexer.Lexer)ASTNode{
	p.lexer=lexer
	p.nextToken()
	prog:=p.programparser()
	return prog
}

func (p *Parser)nextToken()lexer.Token{
	old:=p.current
p.current=p.lexer.NextToken()
return old
}

func (p *Parser)programparser() ASTNode {
	var statements []ASTNode
	stmt :=p.statementparser()
	statements = append(statements, stmt)
	return Program{Statements: statements}
}

func (p *Parser)statementparser() ASTNode {
	if p.current.Type== lexer.TYPEDEFF{
		return p.vardecparser()
	}
	return p.expstatement()
}

func (p *Parser)vardecparser() ASTNode{
	//we have to call the next token now

	typedeff:=p.nextToken()
	//ok:=lexer.next(*token) 		
		
	if p.current.Type != lexer.IDENTIFIER {
		panic(fmt.Sprintf("expected identifier after 'let' at Line %d",typedeff.Line))
	}

	varName:=p.nextToken()

	// consume identifier

	if p.current.Type != lexer.EQUAL {
		panic(fmt.Sprintf("expected '=' after identifier '%s' at Line %d", varName.Value, varName.Line))
	}


	p.nextToken()
	// consume '='

	val := p.addsubparser()
	fmt.Println("var is made")
	return VarDecl{
		Typedeff:typedeff.Value,
		NodeName: "let-decl",
		Name: varName,
		Value:    val,
		Line: varName.Line,
		Column:   varName.Column,
		// what happend if we go 
		//int x=
		// 3						
	}
}

func (p *Parser)expstatement() ASTNode {
	expr := p.addsubparser()
	return ExprStatement{Expr: expr, Line: p.current.Line, Column: p.current.Column}
}

func (p *Parser)addsubparser() ASTNode {
	left := p.muldivparser()
	if p.current.Type == lexer.PLUS || p.current.Type == lexer.MINUS {
	op:=p.current.Type
	p.nextToken()
		right := p.muldivparser()
		left = Binary{
			NodeName: "binary-addsub",
			Left:     left,
			Operator: op,
			Right:    right,
			Line:     p.current.Line,
			Column:  p.current.Column,
		}
	}
	
	return left
}

func (p *Parser)muldivparser() ASTNode {
	left := p.numgroupparser()

		if p.current.Type == lexer.STAR || p.current.Type == lexer.SLASH {
		op:=p.current.Type
		p.nextToken()
		right := p.numgroupparser()
		left = Binary{
			NodeName: "binary-muldiv",
			Left:     left,
			Operator: op,
			Right:    right,
			Line: p.current.Line,
			Column:   p.current.Column,
		}
	}
	return left
}

func (p *Parser) numgroupparser() ASTNode {
	if p.current.Type == lexer.NUMBER {
		old:=p.nextToken()
		return Literal{NodeName: "lit", Value: old, Line: old.Line, Column: old.Column}
	}

	if p.current.Type  == lexer.IDENTIFIER {

		old:=p.nextToken()
		return Identifier{NodeName: "ident", Name: old, Line: old.Line, Column: old.Column}
	}

	if p.current.Type == lexer.NOT {

		old:=p.nextToken()
		exp := p.addsubparser()
		return Unary{NodeName: "not", Value: exp, Line: old.Line, Column: old.Column}
	}


	p.nextToken()
	exp := p.addsubparser()
	old:=p.nextToken()
	return Groups{NodeName: "bracket", Value: exp, Line: old.Line, Column: old.Column}
}

