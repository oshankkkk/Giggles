package parser

import (
	"fmt"
	"lang/internal/lexer"
)

type Parser struct {
	pointer int
	current lexer.Token
	lexer *lexer.Lexer
	symboltb map[string]bool
	idcounter int
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

func (p *Parser) statementparser() ASTNode {
	if p.current.Type == lexer.TYPEDEFF || p.current.Type == lexer.LOCAL {
		return p.vardecparser()
	}
	if p.current.Type == lexer.RETURN {
		return p.returnparser()
	}
	return p.expstatement()
}

func (p *Parser) returnparser() ASTNode {
	tok := p.nextToken() // consume 'return', tok = RETURN token
	// If the next token can start an expression, parse it.
	// Otherwise treat it as a bare `return` (void).
	var val ASTNode
	if p.current.Type != lexer.END &&
		p.current.Type != lexer.ELSE &&
		p.current.Type != lexer.EOF {
		val = p.parser(0)
	}
	return ReturnStmt{
		Value:  val,
		Line:   tok.Line,
		Column: tok.Column,
	}
}


func (p *Parser)expstatement() ASTNode {
	expr := p.parser(0)
	return ExprStatement{Expr: expr, Line: p.current.Line, Column: p.current.Column}
}


func (p *Parser)vardecparser() ASTNode{
	//we have to call the next token now
	
	var typedeff lexer.Token
	isConst := true
	isLocal := false 
	scope:=p.nextToken()
	if scope.Type==lexer.LOCAL{
		isLocal = true 
		typedeff=p.nextToken()
		fmt.Println("skipipipi")
	}else{
		typedeff=scope
	}
	var varName lexer.Token
	//typedeff:=p.nextToken()
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
	}else{
		fmt.Println("euqla innahwa")
	}

	p.nextToken()
	// consume '='

	val := p.parser(0)
	fmt.Println("var is made")
	p.symboltb[varName.Value]=isConst
	p.idcounter++
	return 	VarDecl{
		Typedeff:typedeff.Value,
		Name: varName,
		Value:    val,
		Line: varName.Line,
		Column:   varName.Column,
		IsConst: isConst,
		IsLocal: isLocal,
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

	if p.current.Type == lexer.FN {

		p.idcounter++
		token := p.nextToken() // consume 'fn', token = FN

		line := token.Line
		col := token.Column
		funcname := p.nextToken() // consume function name

		// --- parse optional parameter list: fn foo(int x, bool y) ---
		var params []Param
		if p.current.Type == lexer.LEFT_PAREN {
			p.nextToken() // consume '('
			for p.current.Type != lexer.RIGHT_PAREN {
				if p.current.Type != lexer.TYPEDEFF {
					panic(fmt.Sprintf("expected type in parameter list but found '%s' at line %d", p.current.Value, p.current.Line))
				}
				paramType := p.nextToken() // consume type keyword (e.g. 'int')

				if p.current.Type != lexer.IDENTIFIER {
					panic(fmt.Sprintf("expected parameter name but found '%s' at line %d", p.current.Value, p.current.Line))
				}
				paramName := p.nextToken() // consume parameter name

				params = append(params, Param{
					Typedeff: paramType.Value,
					Name:     paramName,
					Line:     paramName.Line,
					Column:   paramName.Column,
				})

				if p.current.Type == lexer.COMMA {
					p.nextToken() // consume ',' between parameters
				}
			}
			p.nextToken() // consume ')'
		}
		// --- end parameter list ---

		// --- parse optional return type: fn foo(int x) int ---
		// The return type must appear on the SAME LINE as the fn declaration
		// to avoid consuming the first statement of the body (e.g. `int z = ...`).
		var returnType string
		isVoid := true
		if p.current.Type == lexer.TYPEDEFF && p.current.Line == funcname.Line {
			returnType = p.nextToken().Value // consume the return type token
			isVoid = false
		}
		// --- end return type ---


		// Register params in the symbol table so the body can reference them.
		for _, param := range params {
			p.symboltb[param.Name.Value] = false
		}

		var content []ASTNode
		for p.current.Type != lexer.END {
			content = append(content, p.statementparser())
		}
		if p.current.Type == lexer.END {
			p.nextToken()
		} else {
			panic(fmt.Sprintf("expected 'end' for fn block at Line %d", token.Line))
		}

		// Clean up params from symbol table — they are scoped to this function.
		for _, param := range params {
			delete(p.symboltb, param.Name.Value)
		}

		if funcname.Type == lexer.MAIN {
			return Function{
				Name:       funcname.Value,
				Ismain:     true,
				Params:     params,
				ReturnType: returnType,
				IsVoid:     isVoid,
				Content:    content,
				Line:       line,
				Column:     col,
			}
		}
		return Function{
			Name:       funcname.Value,
			Params:     params,
			ReturnType: returnType,
			IsVoid:     isVoid,
			Content:    content,
			Line:       line,
			Column:     col,
		}
	}


	if p.current.Type == lexer.IDENTIFIER || p.current.Type==lexer.MAIN {
		name := p.current
		p.nextToken()

		if p.current.Type == lexer.LEFT_PAREN {
			p.nextToken() 
			var args []Arg
			for p.current.Type != lexer.RIGHT_PAREN {
				argLine := p.current.Line
				argCol := p.current.Column
				expr := p.parser(0)
				args = append(args, Arg{
					Value:  expr,
					Line:   argLine,
					Column: argCol,
				})
				if p.current.Type == lexer.COMMA {
					p.nextToken() 
				}
			}
			p.nextToken() 
			return Call{
				Function: name.Value,
				Args:     args,
				Line:     name.Line,
				Column:   name.Column,
			}
		}


		// plain identifier
		value, ok := p.symboltb[name.Value]
		if !ok {
			panic("var not defined")
		}
		if value == true && p.current.Type == lexer.EQUAL {
			panic("const var cant be mutated, gotta use the mut keyword")
		}
		return Identifier{
			Name:   name,
			Line:   name.Line,
			Column: name.Column,
		}
	}
	if p.current.Type == lexer.ILLEGAL{
		fmt.Println("ILLEGAL hambuna")
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
	if p.current.Type == lexer.MINUS || p.current.Type == lexer.NOT {
		op := p.current
		p.nextToken()
		value := p.parser(100) // Assuming Unary operator has precedence 100
		return Unary{
			Value: value,
			Operator: op.Type,
			Line: op.Line,
			Column: op.Column,
		}
	}
//	fmt.Println("at the end thatma local awula")
	panic(fmt.Sprintf("unexpected token: %s at line %d column %d", p.current.Type, p.current.Line, p.current.Column))
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

func DebugTokens(l *lexer.Lexer) {
	fmt.Println("=== TOKEN STREAM ===")
	for {
		tok := l.NextToken()
		fmt.Printf("[%d:%d]\t%-20s %q\n", tok.Line, tok.Column, tok.Type, tok.Value)
		if tok.Type == lexer.EOF {
			break
		}
	}
	fmt.Println("====================")
}
