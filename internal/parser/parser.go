package parser

import (
	"fmt"
	"lang/internal/lexer"
)

func Parser(tokenlist []lexer.Token)ASTNode{
	pointer := 0
	prog := programparser(tokenlist, &pointer)
	return prog
	//Pp(prog, 0)
}

func programparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	var statements []ASTNode
	for *pointer < len(tokenlist) {
		stmt := statementparser(tokenlist, pointer)
		statements = append(statements, stmt)
	}
	return Program{Statements: statements}
}

func statementparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	char:=tokenlist[*pointer]	
	if char.Type== lexer.TYPEDEFF{
		return vardecparser(tokenlist, pointer)
	}
	return expstatement(tokenlist, pointer)
}

func vardecparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	typedeff:= tokenlist[*pointer]

	*pointer++ 		
		
	if *pointer >= len(tokenlist) || tokenlist[*pointer].Type != lexer.IDENTIFIER {
		panic(fmt.Sprintf("expected identifier after 'let' at Line %d",typedeff.Line))
	}
	varName:= tokenlist[*pointer]
	*pointer++
	// consume identifier

	if *pointer >= len(tokenlist) || tokenlist[*pointer].Type != lexer.EQUAL {
		panic(fmt.Sprintf("expected '=' after identifier '%s' at Line %d", varName.Value, varName.Line))
	}
	*pointer++ // consume '='

	val := addsubparser(tokenlist, pointer)
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

func expstatement(tokenlist []lexer.Token, pointer *int) ASTNode {
	tok := tokenlist[*pointer]
	expr := addsubparser(tokenlist, pointer)
	return ExprStatement{Expr: expr, Line: tok.Line, Column: tok.Column}
}

func addsubparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	left := muldivparser(tokenlist, pointer)
	for *pointer < len(tokenlist) {
		char := tokenlist[*pointer]
		if char.Type != lexer.PLUS && char.Type != lexer.MINUS {
			break
		}
		*pointer++
		right := muldivparser(tokenlist, pointer)
		left = Binary{
			NodeName: "binary-addsub",
			Left:     left,
			Operator: char.Type,
			Right:    right,
			Line:     char.Line,
			Column:   char.Column,
		}
	}
	return left
}

func muldivparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	left := numgroupparser(tokenlist, pointer)
	for *pointer < len(tokenlist) {
		char := tokenlist[*pointer]
		if char.Type != lexer.STAR && char.Type != lexer.SLASH {
			break
		}
		*pointer++
		right := numgroupparser(tokenlist, pointer)
		left = Binary{
			NodeName: "binary-muldiv",
			Left:     left,
			Operator: char.Type,
			Right:    right,
			Line:     char.Line,
			Column:   char.Column,
		}
	}
	return left
}

func numgroupparser(tokenList []lexer.Token, pointer *int) ASTNode {
	char := tokenList[*pointer]

	if char.Type == lexer.NUMBER {
		*pointer++
		return Literal{NodeName: "lit", Value: char, Line: char.Line, Column: char.Column}
	}

	if char.Type == lexer.IDENTIFIER {
		*pointer++
		return Identifier{NodeName: "ident", Name: char, Line: char.Line, Column: char.Column}
	}

	if char.Type == lexer.NOT {
		*pointer++
		exp := addsubparser(tokenList, pointer)
		return Unary{NodeName: "not", Value: exp, Line: char.Line, Column: char.Column}
	}

	*pointer++
	exp := addsubparser(tokenList, pointer)
	*pointer++// consume Right_PAREN
	return Groups{NodeName: "bracket", Value: exp, Line: char.Line, Column: char.Column}
}

