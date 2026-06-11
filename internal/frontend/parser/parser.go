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
var bpmap = map[lexer.TokenType]int{

	lexer.STAR: 50,
	lexer.SLASH: 50,

	lexer.PLUS: 40,
	lexer.MINUS: 40,

	lexer.GREATER: 30,
	lexer.LESS: 30,
	lexer.GREATER_EQUAL: 30,
	lexer.LESS_EQUAL: 30,

	lexer.EQUAL_EQUAL: 20,
	lexer.NOT_EQUAL: 20,

	lexer.AND: 10,
	lexer.OR:  5,
}

func (p *Parser) Run(lexer *lexer.Lexer)ASTNode{

	p.lexer=lexer
	p.nextToken()
	prog:=p.programparser()
	return prog
}

func (p *Parser)nextToken()lexer.Token{
	old:=p.current
	new:=p.lexer.NextToken()
	p.current=new
	return old
}

func (p *Parser)programparser() ASTNode {
	var statements []ASTNode
	for p.current.Type != lexer.EOF { 
		stmt := p.statementparser()
		statements = append(statements, stmt)
	}
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

	val := p.parser(0)
	fmt.Println("var is made")
	return VarDecl{
		Typedeff:typedeff.Value,
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
	expr := p.parser(0)

	return ExprStatement{Expr: expr, Line: p.current.Line, Column: p.current.Column}
}


func (p *Parser) parseStart() ASTNode {

	if p.current.Type == lexer.NUMBER ||  p.current.Type == lexer.TRUE || p.current.Type== lexer.FALSE{
		node := Literal{
			Value: p.current,
			Line: p.current.Line,
			Column: p.current.Column,
		}
		p.nextToken()
		return node
	}

	if p.current.Type == lexer.IDENTIFIER {
		node := Identifier{
			Name: p.current,
			Line: p.current.Line,
			Column: p.current.Column,
		}
		p.nextToken()
		return node
	}
	
	if p.current.Type == lexer.NOT {
		token := p.nextToken()

		return Unary{
			Value: p.parser(30),
			Line: token.Line,
			Column: token.Column,
		}
	}

	if p.current.Type == lexer.LEFT_PAREN{
		token := p.nextToken()
		inner:=p.parser(0)
		if p.current.Type==lexer.RIGHT_PAREN{
			p.nextToken()
		}	

		return Groups{
			Value: inner,
			Line: token.Line,
			Column: token.Column,
		}
	}

	//if p.current.Type==lexer.IF{
	//	token:=p.nextToken()
	//	return Condition{
	//		Value: p.parser(0),
	//		Line: token.Line,
	//		Column: token.Column,
	//	}
	//}

	//if p.current.Type==lexer.FOR{
	//	token:=p.nextToken()
	//	return Loop{
	//		Value: p.parser(0),
	//		Line: token.Line,
	//		Column: token.Column,
	//	}
	//}

	//this idk
	return Binary{}
}

func (p *Parser) parser(minBp int) ASTNode {
	node := p.parseStart()

	for {
		value, ok := bpmap[p.current.Type]
		if ok && value >minBp {

			op := p.current
			p.nextToken()

			right := p.parser(value+1)
			node = Binary{
				Left: node,
				Operator: op.Type,
				Right: right,
				Line: op.Line,
				Column: op.Column,
			}
		}else{
			break
		}

	}

	return node
}
