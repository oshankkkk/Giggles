package parser

import (
	"fmt"
	"lang/internal/lexer"
	"strings"
)

func Parser(tokenlist []lexer.Token) {
	pointer := 0
	prog := programparser(tokenlist, &pointer)
	Pp(prog, 0)
}

func programparser(tokenlist []lexer.Token, pointer *int) ASTNode {
	var statements []ASTNode
	for *pointer < len(tokenlist) {
		stmt := statementparser(tokenlist, pointer)
		statements = append(statements, stmt)
	}
	return Program{statements: statements}
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
		panic(fmt.Sprintf("expected identifier after 'let' at line %d",typedeff.Line))
	}
	varname:= tokenlist[*pointer]
	*pointer++
	// consume identifier

	if *pointer >= len(tokenlist) || tokenlist[*pointer].Type != lexer.EQUAL {
		panic(fmt.Sprintf("expected '=' after identifier '%s' at line %d", varname.Value, varname.Line))
	}
	*pointer++ // consume '='

	val := addsubparser(tokenlist, pointer)

	return VarDecl{
		typedeff:typedeff.Value,
		nodeName: "let-decl",
		name: varname,
		value:    val,
		line: varname.Line,
		column:   varname.Column,
	}
}

func expstatement(tokenlist []lexer.Token, pointer *int) ASTNode {
	tok := tokenlist[*pointer]
	expr := addsubparser(tokenlist, pointer)
	return ExprStatement{expr: expr, line: tok.Line, column: tok.Column}
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
			nodeName: "binary-addsub",
			left:     left,
			operator: char.Type,
			right:    right,
			line:     char.Line,
			column:   char.Column,
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
			nodeName: "binary-muldiv",
			left:     left,
			operator: char.Type,
			right:    right,
			line:     char.Line,
			column:   char.Column,
		}
	}
	return left
}

func numgroupparser(tokenList []lexer.Token, pointer *int) ASTNode {
	char := tokenList[*pointer]

	if char.Type == lexer.NUMBER {
		*pointer++
		return Literal{nodeName: "lit", value: char, line: char.Line, column: char.Column}
	}

	if char.Type == lexer.IDENTIFIER {
		*pointer++
		return Identifier{nodeName: "ident", name: char, line: char.Line, column: char.Column}
	}

	if char.Type == lexer.NOT {
		*pointer++
		exp := addsubparser(tokenList, pointer)
		return Unary{nodeName: "not", value: exp, line: char.Line, column: char.Column}
	}

	*pointer++
	exp := addsubparser(tokenList, pointer)
	*pointer++// consume RIGHT_PAREN
	return Groups{nodeName: "bracket", value: exp, line: char.Line, column: char.Column}
}

func Pp(ex ASTNode, indent int) {
	pad := strings.Repeat("  ", indent)
	switch n := ex.(type) {
	case Program:
		fmt.Printf("%sProgram\n", pad)
		for _, s := range n.statements {
			Pp(s, indent+1)
		}
	case VarDecl:
		fmt.Printf("%sLetDecl(%s)\n", pad, n.name.Value)
		Pp(n.value, indent+1)
	case ExprStatement:
		fmt.Printf("%sExprStatement\n", pad)
		Pp(n.expr, indent+1)
	case Literal:
		fmt.Printf("%sLiteral(%s)\n", pad, n.value.Value)
	case Identifier:
		fmt.Printf("%sIdentifier(%s)\n", pad, n.name.Value)
	case Binary:
		fmt.Printf("%sBinary(%s)\n", pad, n.operator)
		Pp(n.left, indent+1)
		Pp(n.right, indent+1)
	case Groups:
		fmt.Printf("%sGroup\n", pad)
		Pp(n.value, indent+1)
	case Unary:
		fmt.Printf("%sUnary\n", pad)
		Pp(n.value, indent+1)
	default:
		fmt.Printf("%s???\n", pad)
	}
}
