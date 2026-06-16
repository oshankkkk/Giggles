package parser

import (
	"fmt"
	//"go/constant"
	"lang/internal/lexer"
)

type Parser struct {
	pointer int
	current lexer.Token
	lexer *lexer.Lexer
	symboltb map[string]bool
}

var bpmap = map[lexer.TokenType]int{

	lexer.STAR: 50,
	lexer.SLASH: 50,

	lexer.EQUAL: 20,

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
	p.symboltb = make(map[string]bool)
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

func (p *Parser)expstatement() ASTNode {
	expr := p.parser(0)

	return ExprStatement{Expr: expr, Line: p.current.Line, Column: p.current.Column}
}


func (p *Parser)vardecparser() ASTNode{
	//we have to call the next token now
	isConst := true
	var varName lexer.Token
	typedeff:=p.nextToken()
	//ok:=lexer.next(*token) 		

	if !(p.current.Type == lexer.IDENTIFIER || p.current.Type ==lexer.MUT) {
		panic(fmt.Sprintf("expected identifier but found %s at Line %d",p.current.Value,typedeff.Line))
	}

	mut:=p.nextToken()
	if mut.Type==lexer.MUT{
		isConst=false
		varName=p.nextToken()
	}else{
		varName=mut
	}

	// consume identifier

	if p.current.Type != lexer.EQUAL {
		panic(fmt.Sprintf("expected '=' but found '%s' at Line %d", varName.Value, varName.Line))
	}

	p.nextToken()
	// consume '='

	val := p.parser(0)
	fmt.Println("var is made")
	p.symboltb[varName.Value]=isConst

	return 	VarDecl{
		Typedeff:typedeff.Value,
		Name: varName,
		Value:    val,
		Line: varName.Line,
		Column:   varName.Column,
		isConst: isConst,
		}


}
func (p *Parser) parseStart() ASTNode {
	fmt.Println(p.current.Value)

	if p.current.Type == lexer.NUMBER ||  p.current.Type == lexer.TRUE || p.current.Type== lexer.FALSE {
		node := Literal{
			Value: p.current,
			Line: p.current.Line,
			Column: p.current.Column,
		}
		p.nextToken()
		return node
	}

	if p.current.Type == lexer.IDENTIFIER {
		old:=p.nextToken()
	if value,ok:=p.symboltb[old.Value];ok{
		if value==true && p.current.Type==lexer.EQUAL{
				panic("const var cant be mutated, gotta use the mut keyword")
			}else{
	
		node := Identifier{
			Name: old,
			Line: old.Line,
			Column: old.Column,
		}
		return node
			}
		}else{
			panic("var not defined")
		}
	}

	if p.current.Type == lexer.ILLEGAL{
		panic(fmt.Sprintf("expected '%s' at Line %d", p.current.Value,p.current.Line))
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

	if p.current.Type == lexer.IF || p.current.Type==lexer.FOR{
		var isLooped bool
		if p.current.Type==lexer.FOR{
			isLooped=true
		}
		token := p.current
		line := token.Line
		col := token.Column

		p.nextToken() // move into condition

		condition := p.parser(0)
		
		isBlock := false
		if p.current.Type == lexer.THEN {
			isBlock = true
			p.nextToken()
		}

		var thenResult []ASTNode
		if isBlock {
			for p.current.Type != lexer.ELSE && p.current.Type != lexer.END && p.current.Type != lexer.EOF {
				thenResult = append(thenResult, p.statementparser())
			}
		} else {
			thenResult = append(thenResult, p.statementparser())
		}

		var elseResult []ASTNode
		hasElse := false

		if p.current.Type == lexer.ELSE {
			hasElse = true
			p.nextToken() // consume ELSE
			if isBlock {
				for p.current.Type != lexer.END && p.current.Type != lexer.EOF {
					elseResult = append(elseResult, p.statementparser())
				}
			} else {
				elseResult = append(elseResult, p.statementparser())
			}
		}

		if isBlock {
			if p.current.Type == lexer.END {
				p.nextToken()
			} else {
				panic(fmt.Sprintf("expected 'end' for if-then block at Line %d", token.Line))
			}
		}

		return Condition{
			Condition:  condition,
			Result:     thenResult,
			ElseResult:  elseResult,
			HasElse:     hasElse,
			Looped:isLooped,
			Line:        line,
			Column:      col,
		}
	}
panic("eeeeehdhdhd")
}

func (p *Parser) parser(minBp int) ASTNode {
	node := p.parseStart()

	for {
		value, ok := bpmap[p.current.Type]
		fmt.Println(value)
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
