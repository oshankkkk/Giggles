package compiler

import (
	"fmt"
	//"lang/internal/compiler"
	//	"fmt"
	"lang/internal/lexer"
	"lang/internal/parser"
	"strconv"
)
type State struct {
	CounterTable  []int
	symboltb map[string]int
	fixups []fixup
	Entrypoint int
	Buff []byte
}
type fixup struct{
	address int
	funname string
}
func (c *State) Wrapper(ast parser.ASTNode,scope *Scope){
	c.symboltb=make(map[string]int)
	c.Buff = append(c.Buff, byte(JMP),0)	
	c.ToBytes(ast,scope)
	c.Buff[0]=byte(c.Entrypoint)
	c.Buff= append(c.Buff, byte(STOP))
}
func (c *State) ToBytes(ast parser.ASTNode, scope *Scope) {
	if value, ok := ast.(parser.Program); ok {
		for _, val := range value.Statements {
			c.ToBytes(val, scope)
		}
	}
	if value, ok := ast.(parser.ExprStatement); ok {
		c.ToBytes(value.Expr, scope)
	}
	if value, ok := ast.(parser.Groups); ok {
		c.ToBytes(value.Value, scope)
	}
	if value, ok := ast.(parser.VarDecl); ok {
		c.ToBytes(value.Value, scope)
		v := scope.AddVariable(value.Name.Value)
		c.CounterTable = append(c.CounterTable, v.id)
		c.Buff = append(c.Buff, byte(SETGLOBAL), byte(len(c.CounterTable)-1))
	}
	if value, ok := ast.(parser.Identifier); ok {
		fmt.Println("ee343")
		if info, ok := scope.VarLookup(value.Name.Value); ok {
			c.CounterTable = append(c.CounterTable, info.id)
			fmt.Println(c.CounterTable)
			c.Buff = append(c.Buff, byte(GETGLOBAL), byte(len(c.CounterTable)-1))
		} else {
			panic("var undefined")
		}
	}
	if value, ok := ast.(parser.Call); ok {
		c.Buff = append(c.Buff, byte(NWFRM))
		funcjmp := len(c.Buff) + 1
		c.Buff = append(c.Buff, byte(JMP), 0)
		if fnaddress, ok := c.symboltb[value.Function]; ok {
			fmt.Println(fnaddress, "fmfmaafress")
			fmt.Println(value.Function, "name mamamamaname funcnam")
			c.CounterTable = append(c.CounterTable, fnaddress)
			c.Buff[funcjmp] = byte(len(c.CounterTable) - 1)
			fmt.Println("yayayayyyaya")
		} else {
			c.fixups = append(c.fixups, fixup{address: funcjmp, funname: value.Function})
		}
	}
	if value, ok := ast.(parser.Function); ok {
		if value.Name == string(lexer.MAIN) {
			fmt.Println("hehehdub")
			c.Entrypoint = len(c.Buff)
		}
		c.symboltb[value.Name] = len(c.Buff)
		for _, n := range value.Content {
			fmt.Println("conconconconconc", value.Name)
			c.ToBytes(n, scope)
		}
		c.Buff = append(c.Buff, byte(RMFRM))
	}
	if value, ok := ast.(parser.Condition); ok {
		condPos := len(c.Buff)
		c.ToBytes(value.Condition, scope)
		jifPos := len(c.Buff)
		c.Buff = append(c.Buff, byte(JIF), 0)
		for _, r := range value.Result {
			c.ToBytes(r, scope)
		}
		if value.Looped {
			c.Buff = append(c.Buff, byte(JMP), byte(condPos))
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
			for _, e := range value.ElseResult {
				c.ToBytes(e, scope)
			}
		} else if value.HasElse {
			jmpPos := len(c.Buff)
			c.Buff = append(c.Buff, byte(JMP), 0)
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
			for _, e := range value.ElseResult {
				c.ToBytes(e, scope)
			}
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jmpPos+1] = byte(len(c.CounterTable) - 1)
		} else {
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
		}
	}
	if value, ok := ast.(parser.Literal); ok {
		if value.Value.Type == lexer.TRUE {
			c.Buff = append(c.Buff, byte(TRUE))
			return
		} else if value.Value.Type == lexer.FALSE {
			c.Buff = append(c.Buff, byte(FALSE))
			return
		}
		index, err := strconv.Atoi(value.Value.Value)
		if err != nil {
			panic(err)
		}
		c.CounterTable = append(c.CounterTable, index)
		c.Buff = append(c.Buff, byte(PUSH), byte(len(c.CounterTable)-1))
	}
	if value, ok := ast.(parser.Binary); ok {
		if value.Operator == lexer.EQUAL {
			c.ToBytes(value.Right, scope)
			if info, ok := scope.VarLookup(value.Left.(parser.Identifier).Name.Value); ok {
				c.CounterTable = append(c.CounterTable, info.id)
				c.Buff = append(c.Buff, byte(ASS), byte(len(c.CounterTable)-1))
			} else {
				panic("var undefined")
			}
			return
		}
		c.ToBytes(value.Left, scope)
		c.ToBytes(value.Right, scope)
		var opcode Opcode
		switch value.Operator {
		case lexer.MINUS:
			opcode = SUB
		case lexer.PLUS:
			opcode = ADD
		case lexer.SLASH:
			opcode = DIV
		case lexer.STAR:
			opcode = MUL
		case lexer.AND:
			opcode = AND
		case lexer.OR:
			opcode = OR
		case lexer.GREATER:
			opcode = GT
		case lexer.LESS:
			opcode = LT
		case lexer.GREATER_EQUAL:
			opcode = GTE
		case lexer.LESS_EQUAL:
			opcode = LTE
		case lexer.EQUAL_EQUAL:
			opcode = EQ
		case lexer.NOT_EQUAL:
			opcode = NEQ
		}
		c.Buff = append(c.Buff, byte(opcode))
	}
}
func (c *State) Fixpatchs()(error){
	// After the whole program is compiled, resolve all fixups

	for _, f := range c.fixups {
		addr, ok := c.symboltb[f.funname]
		if !ok {
			return fmt.Errorf("undefined function: %s", f.funname)
		}

		fmt.Println("babababaabab")		
		// Write the real address into the 2-byte placeholder
		// Adjust encoding to match your VM's word size / endianness
		//c.Buff[f.address]   = byte(addr)   // high byte
		c.CounterTable = append(c.CounterTable, addr)
		c.Buff[f.address] =byte(len(c.CounterTable)-1)

		//c.Buff[f.patchOffset+1] = byte(addr & 0xFF) // low byte
	}
	return nil

}


//		if value.IsLocal{
//			c.Buff = append(c.Buff, byte(SETLOCAL))
//			fmt.Println("heheh")
//		}else{

//	if value, ok:=ast.(parser.Call);ok{
//		c.Buff = append(c.Buff, byte(CALLFUNC))
//	}
//	if value, ok:=ast.(parser.Function);ok{
		//local :=EnterScope(scope)
		//start:=len(c.Buff)
		//framestart:=len(c.Buff)
		//c.Buff = append(c.Buff,byte( ENTERFUNC))
	//	for _,stmnt:=range value.Content{
		//c.Buff = append(c.Buff, c.ToBytes(stmnt,scope)...)
		//	}
			//end:=len(c.Buff)
		//c.CounterTable=append(c.CounterTable,end )	
		//fmt.Println(framestart,"dddrfff")
		//fmt.Println(len(c.CounterTable)-2,"ffffffff")

		//c.Buff = append(c.Buff, byte(JMP), byte(len(c.CounterTable)-2))
	//}
